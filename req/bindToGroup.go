package req

type bindToGroupParams struct {
	SendUserId string `json:"sendUserId"`
	ClientId   string `json:"clientId"`
	GroupName  string `json:"groupName"`
}

type bindToGroupResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

//	绑定客户端到分组
func (r *Request) BindToGroup(sendUserId, clientId, groupName string) (*bindToGroupResponse, error) {

	params := &bindToGroupParams{
		sendUserId,
		clientId,
		groupName,
	}
	result := &bindToGroupResponse{}
	err := r.do(bindToGroup, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
