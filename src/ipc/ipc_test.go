package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"200",method + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1,_ := client1.Call("Get","From Client1")
	resp2,_ := client2.Call("Get","From Client2")

	if resp1.Code != "200" || resp1.Body != "GetFrom Client1" || 
		resp2.Code != "200" || resp2.Body != "GetFrom Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:",resp2)
	}
	client1.Close()
	client2.Close()
}
