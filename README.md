# Golang实现 woodylan/go-websocket的SDK


[![Go](https://img.shields.io/badge/Go-1.13-blue.svg)](https://golang.google.cn)
[![star](https://img.shields.io/github/stars/Hhhha/go-websocket-sdk-go?style=social)](https://github.com/woodylan/go-websocket/stargazers)

## install
```go
    go get github.com/Hhhha/go-websocket-sdk-go
```
## api
```go
    //  获取请求
    request := req.New("localhost:6000", "123")
    list, err := request.GetOnlineList("groupName")
    //  绑定用户到分组
    request.BindToGroup("sendUserId", "clientId", "groupName")
    //  给分组发送消息
    request.SendToGroup("sendUserId", "groupName", 200, "msg", "data")
    //  单独发送消息
    request.SendToClient("clientId", "sendUserId", 200, "msg", "data")
    clientIds := make([]string, 0)
    clientIds = append(clientIds, "clientId")
    //  批量发送消息
    request.SendToClientIds(clientIds, "sendUserId", 200, "msg", "data")
    //  关闭链接
    request.CloseClient("clientId")
    fmt.Println(list, err)
```