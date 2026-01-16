package main

import (
	"fmt"
	"log"

	"github.com/sisyphus/go-advanced-programming/ch4-rpc/hello1/api/hello"
)

func main() {
	client, err := hello.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialog:", err)
	}

	var reply string
	err = client.Hello("client", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
