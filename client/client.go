package main

import (
	"fmt"
	"net/rpc"
	"os"
)

// Args represents the arguments for Put and Get operations
type Args struct {
	Key   string
	Value string
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: kvClient [put/get] [key] [value]")
		return
	}

	action := os.Args[1]
	key := os.Args[2]

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	switch action {
	case "put":
		if len(os.Args) != 4 {
			fmt.Println("Usage: kvClient put [key] [value]")
			return
		}
		value := os.Args[3]
		args := Args{Key: key, Value: value}
		var reply string
		err = client.Call("Storage.Put", &args, &reply)
		if err != nil {
			fmt.Println("Error calling Put:", err)
			return
		}
		fmt.Println("Put result:", reply)
	case "get":
		var reply string
		err = client.Call("Storage.Get", &key, &reply)
		if err != nil {
			fmt.Println("Error calling Get:", err)
			return
		}
		fmt.Println("Get result:", reply)
	default:
		fmt.Println("Unknown action:", action)
	}
}
