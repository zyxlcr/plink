package main

import (
	"chatcser/pkg/plink/etcd"
	"context"
	"fmt"
)

func main() {
	arr := []string{"localhost:2379"}
	cli := etcd.NewCluster(arr)
	rch := cli.Watch(context.Background(), "endpoints6")
	go func() {
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()

	res, err := cli.Put(context.Background(), "endpoints6", "12345")
	if err != nil {
		return
	}
	println(string(res.Header.ClusterId))
	res2, err := cli.Get(context.Background(), "endpoints6")
	if err != nil {
		return
	}
	println(string(res2.Kvs[len(res2.Kvs)-1].Value))

	cli.Put(context.Background(), "endpoints6", "999")
	if err != nil {
		return
	}
	//etcd.Watch("endpoints")
	//etcd.Put("endpoints2", "123")
	//etcd.Put("endpoints", "456")
	//etcd.Get("endpoints")
	for {
	}

}
