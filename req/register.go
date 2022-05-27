package req

type registerParams struct {
	SystemId string `json:"systemId"`
}

type registerResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

func (r *Request) Register(systemId string) (*registerResponse, error) {

	params := &registerParams{
		systemId,
	}
	result := &registerResponse{}
	err := r.do(register, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
