package req

type getOnlineListParams struct {
	GroupName string `json:"groupName"`
}

type getOnlineListDataResponse struct {
	Count int      `json:"count"`
	List  []string `json:"list"`
}

type getOnlineListResponse struct {
	Code int                        `json:"code"`
	Msg  string                     `json:"msg"`
	Data *getOnlineListDataResponse `json:"data"`
}

func (r *Request) GetOnlineList(groupName string) (*getOnlineListResponse, error) {

	params := &getOnlineListParams{
		groupName,
	}
	result := &getOnlineListResponse{}
	err := r.do(getOnlineList, params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
