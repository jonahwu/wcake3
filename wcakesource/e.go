package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type ServiceInfo struct {
	User    string `json:"user"`
	Address string `json:"address"`
}

func main() {
	fmt.Println("connecting......")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.2.177:2379"},
		DialTimeout: 5 * time.Second,
	})
	timeout := 5 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connected......")
	defer cli.Close() // make sure to close the client

	// put one data
	fmt.Println("put data ......")
	//_, err = cli.Put(context.TODO(), "foo1", "bar1")
	_, err = cli.Put(ctx, "foo1", "bar1")
	//data, _ := cli.Get(context.TODO(), "foo1")
	data, _ := cli.Get(ctx, "foo1")
	fmt.Println(string(data.Kvs[0].Key), string(data.Kvs[0].Value))
	fmt.Println("get size", data.Count)

	for _, ev := range data.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	// another put struct
	serviceInfo := ServiceInfo{User: "aaa", Address: "bbb"}
	value, _ := json.Marshal(serviceInfo)
	_, err = cli.Put(ctx, "s1", string(value))

	datas, _ := cli.Get(ctx, "s1")
	fmt.Println("the key:", string(datas.Kvs[0].Key), "the value:", string(datas.Kvs[0].Value))
	//datas.Kvs[0].Value
	//var dataOutput ServiceInfo
	var dataOutput map[string]interface{}
	_ = json.Unmarshal(datas.Kvs[0].Value, &dataOutput)
	fmt.Println("unmashal result", dataOutput)
	fmt.Println("unmashal result element", dataOutput["address"])

	fmt.Println("get size", datas.Count)
	for _, ev := range datas.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	//data.Kvs
	// another get nothing
	datan, _ := cli.Get(ctx, "nothing")
	fmt.Println("nothing size:", datan.Count)

	// get a sub list
	// with put aaa/aaa "aaa"; aaa/bbb "bbb"; aaa/ccc "ccc"
	//dataL, _ := cli.Get(ctx, "aaa", clientv3.WithPrefix(), clientv3.WithKeysOnly())
	// aaa = buy/userID/product(1,2,3)->value(prodcut1,2,3)   buy/product1 -> value;  buy:process:done
	dataL, _ := cli.Get(ctx, "aaa", clientv3.WithPrefix())
	fmt.Println("List size:", dataL.Count)
	for _, ev := range dataL.Kvs {
		fmt.Printf("show on List%s : %s\n", ev.Key, ev.Value)
	}

	//_, err = cli.Put(context, "foo1", "bar1")
	if err != nil {
		fmt.Println(err)
	}
	// watch function
	rch := cli.Watch(context.Background(), "/test/hello", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

}
