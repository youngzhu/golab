package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// 测试RPC之间能不能传递指针对象
/*
从服务端获取对象
在客户端改变某些字段
这时服务端有变化吗？===没变化
*/

/*
通过RPC传递，地址肯定变了，即对象已不是原来的对象了。
*/

// 结论：服务端和客户端是两个世界。是通过gob转码和编码的，对其状态互不影响

type Message struct {
	Text   string
	Status string
}

func (m Message) String() string {
	return fmt.Sprintf("Text:%s, Status:%s, %p", m.Text, m.Status, &m)
}

type (
	GetArgs struct {
	}
	GetReply struct {
		Message *Message
	}

	GetArgs2 struct {
	}
	GetReply2 struct {
		Message Message
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

func get() *Message {
	client := connect()
	args := GetArgs{}
	reply := GetReply{}
	err := client.Call("KV.Get", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()

	msg := reply.Message
	msg.Status = "sent-last"

	return msg
}
func get2() Message {
	client := connect()
	args := GetArgs2{}
	reply := GetReply2{}
	err := client.Call("KV.Get2", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()

	msg := reply.Message
	msg.Status = "sent2-last"

	return msg
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

func (kv *KV) Get(args *GetArgs, reply *GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	reply.Message = kv.messages["test"]

	return nil
}

func (kv *KV) Get2(args *GetArgs2, reply *GetReply2) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	reply.Message = kv.messages2["test2"]

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
	kv.messages["test"] = msg
	fmt.Printf("%p\n", msg)
	fmt.Printf("%p\n", msg)
	fmt.Println(msg)
	fmt.Println(msg)
	get()
	fmt.Println(msg)
	fmt.Printf("%p\n", msg)
	svrMsg := kv.messages["test"]
	fmt.Println(svrMsg)
	fmt.Printf("%p\n", svrMsg)

	println("==========")
	println("值传递")
	msg2 := Message{Text: "hi", Status: "create"}
	kv.messages2["test2"] = msg2
	fmt.Printf("%p\n", &msg2)
	fmt.Printf("%p\n", &msg2)
	fmt.Println(msg2)
	fmt.Println(msg2)
	get2()
	fmt.Println(msg2)
	fmt.Printf("%p\n", &msg2)
	svrMsg2 := kv.messages2["test2"]
	fmt.Println(svrMsg2)
	fmt.Printf("%p\n", &svrMsg2)
}
