package web

type GetRequestGroupWebRequest struct {
	RequestGroupId string `params:"request_group_id" validate:"required"`
}

type CreateRequestWebRequest struct {
	RequestGroupId string `params:"request_group_id" validate:"required"`
}

type DeleteRequestWebRequest struct {
	RequestGroupId string `params:"request_group_id" validate:"required"`
	RequestId      string `params:"request_id" validate:"required"`
}

type ConnectSseWebRequest struct {
	RequestGroupId string `params:"request_group_id" validate:"required"`
}
