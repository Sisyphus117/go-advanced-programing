package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			fmt.Println("watch:", keyChanged)
		}
	}()

	err := client.Call("KVStoreService.Set", [2]string{"abc", "abc-value"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)
}
