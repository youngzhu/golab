package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"
)

// 测试RPC之间能不能传递指针对象

// 失败的例子，想岔了。。

/*
通过RPC传递，地址肯定变了，即对象已不是原来的对象了。
*/

type Message struct {
	Text       string
	Status     string
	CreateTime time.Time
	UpdateTime time.Time
}

func (m Message) String() string {
	return fmt.Sprintf("Text:%s, Status:%s, %p", m.Text, m.Status, &m)
}

type (
	SendArgs struct {
		Message *Message
	}
	SendReply struct {
	}

	SendArgs2 struct {
		Message Message
	}
	SendReply2 struct {
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

func send(msg *Message) {
	msg.Status = "sent-first"

	client := connect()
	args := SendArgs{msg}
	reply := SendReply{}
	err := client.Call("KV.Send", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()

	msg.Status = "sent-last"
}
func send2(msg Message) {
	msg.Status = "sent2-first"

	client := connect()
	args := SendArgs2{msg}
	reply := SendReply2{}
	err := client.Call("KV.Send2", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()

	msg.Status = "sent2-last"
}

// Server

type KV struct {
	mu        sync.Mutex
	messages  map[string]*Message
	messages2 map[string]Message
}

func server(kv *KV) {

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

func (kv *KV) Send(args *SendArgs, reply *SendReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.messages["test"] = args.Message

	return nil
}

func (kv *KV) Send2(args *SendArgs2, reply *SendReply2) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.messages2["test2"] = args.Message

	return nil
}

func main() {
	kv := new(KV)
	kv.messages = map[string]*Message{}
	kv.messages2 = map[string]Message{}

	m := &Message{}
	m2 := Message{}
	kv.messages[""] = m
	kv.messages2["2"] = m2

	fmt.Printf("m: %p\n", m)
	mm := kv.messages[""]
	fmt.Printf("mm: %p\n", mm)

	fmt.Printf("m2: %p\n", &m2)
	mm2 := kv.messages2[""]
	fmt.Printf("mm2: %p\n", &mm2)

	server(kv)

	println("指针传递")
	msg := &Message{Text: "hello", Status: "create"}
	fmt.Printf("%p\n", msg)
	fmt.Printf("%p\n", msg)
	fmt.Println(msg)
	fmt.Println(msg)
	send(msg)
	fmt.Println(msg)
	fmt.Printf("%p\n", msg)
	svrMsg := kv.messages["test"]
	fmt.Println(svrMsg)
	fmt.Printf("%p\n", svrMsg)

	println("值传递")
	msg2 := Message{Text: "hi", Status: "create"}
	fmt.Printf("%p\n", &msg2)
	fmt.Printf("%p\n", &msg2)
	fmt.Println(msg2)
	fmt.Println(msg2)
	send2(msg2)
	fmt.Println(msg2)
	fmt.Printf("%p\n", &msg2)
	svrMsg2 := kv.messages2["test2"]
	fmt.Println(svrMsg2)
	fmt.Printf("%p\n", &svrMsg2)
}
