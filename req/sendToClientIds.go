package req

type sendToClientIdsParams struct {
	ClientIds  []string    `json:"clientIds"`
	SendUserId string      `json:"sendUserId"`
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

type sendToClientIdsDataResponse struct {
	MessageId string `json:"messageId"`
}

type sendToClientIdsResponse struct {
	Code int                          `json:"code"`
	Msg  string                       `json:"msg"`
	Data *sendToClientIdsDataResponse `json:"data"`
}

//	批量发送信息给指定客户端
func (r *Request) SendToClientIds(clientIds []string, sendUserId string, code int, msg string, data interface{}) (*sendToClientIdsResponse, error) {
	params := &sendToClientIdsParams{
		clientIds,
		sendUserId,
		code,
		msg,
		data,
	}
	result := &sendToClientIdsResponse{}
	err := r.do(sendToClientIds, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
