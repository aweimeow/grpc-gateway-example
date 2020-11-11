package main

import (
	"fmt"
	"sync"
)

var (
	wg	sync.WaitGroup
)

func main() {
	wg.Add(1)

	fmt.Println("Starting gRPC server")
	go StartGRPCServer()
	fmt.Println("Starting Http reverse proxy server")
	go StartHttpReverseProxyServer()

	wg.Wait()
}