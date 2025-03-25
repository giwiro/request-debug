package web

type GetRequestGroupRequestWebRequest struct {
	RequestGroupId string `params:"request_group_id" validate:"required"`
}
