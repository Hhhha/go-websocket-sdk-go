package req

type sendToGroupParams struct {
	SendUserId string      `json:"sendUserId"`
	GroupName  string      `json:"groupName"`
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

type sendToGroupDataResponse struct {
	MessageId string `json:"messageId"`
}

type sendToGroupResponse struct {
	Code int                      `json:"code"`
	Msg  string                   `json:"msg"`
	Data *sendToGroupDataResponse `json:"data"`
}

//	发送信息给指定分组
func (r *Request) SendToGroup(sendUserId, groupName string, code int, msg string, data interface{}) (*sendToGroupResponse, error) {

	params := &sendToGroupParams{
		sendUserId,
		groupName,
		code,
		msg,
		data,
	}
	result := &sendToGroupResponse{}
	err := r.do(sendToGroup, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
