package web

type GetRequestGroupWebRequest struct {
	RequestGroupId string `params:"request_group_id" validate:"required"`
}

type CreateRequestWebRequest struct {
	RequestGroupId string `params:"request_group_id" validate:"required"`
}
