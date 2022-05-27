package req

import (
	"encoding/json"
	"github.com/Hhhha/go-websocket-sdk-go/config"
	"github.com/valyala/fasthttp"
)

type Request struct {
	req    *fasthttp.Request
	client *fasthttp.Client
	config *config.Config
}

const (
	register        = "/api/register"
	sendToClientIds = "/api/send_to_clients"
	sendToClient    = "/api/send_to_client"
	bindToGroup     = "/api/bind_to_group"
	sendToGroup     = "/api/send_to_group"
	getOnlineList   = "/api/get_online_list"
	closeClient     = "/api/close_client"
)

func New(hosts string, systemId string) *Request {
	req := &fasthttp.Request{}
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.Header.Add("SystemId", systemId)
	return &Request{
		req:    req,
		client: &fasthttp.Client{},
		config: &config.Config{
			Hosts:    hosts,
			SystemId: systemId,
		},
	}
}
func (r *Request) do(url string, params interface{}, result interface{}) error {
	r.req.SetRequestURI(r.config.Hosts + url)

	byteData, err := json.Marshal(params)
	if err != nil {
		return err
	}
	r.req.SetBody(byteData)
	resp := &fasthttp.Response{}
	err = r.client.Do(r.req, resp)
	if err != nil {
		return err
	}
	body := resp.Body()
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}
