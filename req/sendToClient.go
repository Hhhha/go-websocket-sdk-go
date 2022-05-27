package req

type sendToClientParams struct {
	ClientId   string      `json:"clientId"`
	SendUserId string      `json:"sendUserId"`
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

type sendToClientDataResponse struct {
	MessageId string `json:"messageId"`
}

type sendToClientResponse struct {
	Code int                       `json:"code"`
	Msg  string                    `json:"msg"`
	Data *sendToClientDataResponse `json:"data"`
}

//	发送信息给指定客户端
func (r *Request) SendToClient(clientId string, sendUserId string, code int, msg string, data interface{}) (*sendToClientResponse, error) {
	params := &sendToClientParams{
		clientId,
		sendUserId,
		code,
		msg,
		data,
	}
	result := &sendToClientResponse{}
	err := r.do(sendToClient, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}
