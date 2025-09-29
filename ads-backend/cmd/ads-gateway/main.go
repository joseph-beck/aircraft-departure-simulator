package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	var resp string
	err = client.Call("ManageService.Ping", "pong", &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
