package req

type closeClientParams struct {
	ClientId string `json:"clientId"`
}

type closeClientResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

func (r *Request) CloseClient(clientId string) (*closeClientResponse, error) {

	params := &closeClientParams{
		clientId,
	}
	result := &closeClientResponse{}
	err := r.do(closeClient, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
