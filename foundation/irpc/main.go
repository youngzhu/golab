package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// Common RPC request/reply definitions

type Message struct {
	Text string
	Done chan bool
}

type (
	GetArgs struct {
		Key string
	}
	GetReply struct {
		Message  *Message
		Message2 Message
	}
)

// Client

func connect() *rpc.Client {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func get(key string) *Message {
	client := connect()
	args := GetArgs{key}
	reply := GetReply{}
	err := client.Call("KV.Get", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()
	return reply.Message
}

// Server

type KV struct {
	mu       sync.Mutex
	messages map[string]*Message
}

func server() {
	kv := new(KV)
	kv.messages = map[string]*Message{}

	done := make(chan bool)
	kv.messages["Jan"] = &Message{Text: "hello", Done: done}

	rpcs := rpc.NewServer()
	rpcs.Register(kv)
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("error:", err)
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err == nil {
				go rpcs.ServeConn(conn)
			} else {
				break
			}
		}
		l.Close()
	}()
}

func (kv *KV) Get(args *GetArgs, reply *GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	val, ok := kv.messages[args.Key]
	if ok {
		reply.Message = val
		reply.Message2 = *val
	}

	return nil
}

func main() {
	server()

	fmt.Printf("get() -> %v\n", get("Jan"))

	fmt.Println("all done")
}
