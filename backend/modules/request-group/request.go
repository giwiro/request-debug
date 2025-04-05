package request_group

import "request-debug/modules/request-group/model"

type GetRequestGroupRequest struct {
	RequestGroupId string
}

type CreateRequestRequest struct {
	RequestGroupId string
	Request        *model.Request
}
