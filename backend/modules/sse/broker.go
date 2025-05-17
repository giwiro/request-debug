package sse

import (
	"fmt"
	"request-debug/logger"
	"sync"
)

type ClientOperation struct {
	groupId string
	channel chan []byte
}

type Broker interface {
	AddNewClient(string) chan []byte
	RemoveClient(string, chan []byte)
	BroadcastGroup(string, []byte)
	GetClients() map[string][]string
	Handler()
}

type sseBroker struct {
	clientsLock sync.Mutex
	// New client connections are pushed to this channel
	newClients chan ClientOperation
	// Closed client connections are pushed to this channel
	closingClients chan ClientOperation
	// Client connections registry
	clients map[string]map[chan []byte]bool
}

func NewBroker() Broker {
	return &sseBroker{
		newClients:     make(chan ClientOperation),
		closingClients: make(chan ClientOperation),
		clients:        make(map[string]map[chan []byte]bool),
	}
}

func (s *sseBroker) AddNewClient(id string) chan []byte {
	ch := make(chan []byte)

	clientOperation := ClientOperation{
		groupId: id,
		channel: ch,
	}

	s.newClients <- clientOperation

	return ch
}

func (s *sseBroker) RemoveClient(id string, ch chan []byte) {
	clientOperation := ClientOperation{
		groupId: id,
		channel: ch,
	}

	s.closingClients <- clientOperation
}

func (s *sseBroker) BroadcastGroup(id string, message []byte) {
	if s.clients[id] == nil || len(s.clients[id]) == 0 {
		return
	}

	fmt.Println("BROADCASTING", len(s.clients[id]))
	for ch, _ := range s.clients[id] {
		fmt.Println(">", message)
		ch <- message
	}
}

func (s *sseBroker) GetClients() map[string][]string {
	clients := map[string][]string{}

	for k, v := range s.clients {
		conn := make([]string, 0)

		for i, _ := range v {
			a := fmt.Sprintf("%v", i)
			conn = append(conn, a)
		}

		clients[k] = conn
	}

	return clients
}

func (s *sseBroker) Handler() {
	logger.Logger.Info().Msgf("SSE Handler activated")

	for {
		select {
		case c := <-s.newClients:
			s.clientsLock.Lock()
			if s.clients[c.groupId] == nil {
				s.clients[c.groupId] = make(map[chan []byte]bool)
			}

			s.clients[c.groupId][c.channel] = true
			s.clientsLock.Unlock()

		case c := <-s.closingClients:
			s.clientsLock.Lock()

			if s.clients[c.groupId] == nil {
				return
			}

			delete(s.clients[c.groupId], c.channel)

			if len(s.clients[c.groupId]) == 0 {
				delete(s.clients, c.groupId)
			}

			s.clientsLock.Unlock()
		}

	}
}
