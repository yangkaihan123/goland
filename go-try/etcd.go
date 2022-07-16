package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"http://172.16.163.11:2379"},
		DialTimeout: time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	kv := clientv3.NewKV(client)
	putResp, err := kv.Put(context.TODO(), "/demo/A/B", "hello", clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("putResp is :", putResp)
	fmt.Println("Revision :", putResp.Header.Revision)
	if putResp.PrevKv != nil {
		fmt.Println("Prevalue: ", string(putResp.PrevKv.Value))
	}
}
