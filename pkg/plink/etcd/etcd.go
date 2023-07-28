package etcd

import (
	"context"
	"fmt"
	"time"

	client "go.etcd.io/etcd/client/v3"
)

var RECDCil *client.Client

func NewCluster(endpoints []string) *client.Client {
	c := client.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	}
	RECDCil, err := client.New(c)
	if err != nil {
		return nil
	}
	//defer cli.Close()
	return RECDCil
}

func cancelFunc() {
	print("cancelFunc")
}

func Get(key string) (*client.GetResponse, error) {
	timeout := 15 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	cancel()
	return RECDCil.Get(ctx, key)
}

func Put(key string, value string) (*client.PutResponse, error) {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	cancel()
	return RECDCil.Put(ctx, key, value)
}

func Watch(key string) client.WatchChan {
	rch := RECDCil.Watch(context.Background(), key)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
	return rch
}

func Delete(key string) (*client.DeleteResponse, error) {

	return RECDCil.Delete(context.Background(), key)
}

func Lease(key string, value string, second int64) {
	// minimum lease TTL is 5-second
	resp, err := RECDCil.Grant(context.TODO(), second)
	if err != nil {
		print(err)
	}

	// after 5 seconds, the key 'foo' will be removed
	_, err = RECDCil.Put(context.TODO(), key, value, client.WithLease(resp.ID))
	if err != nil {
		print(err)
	}
}

func Revoke(key string, value string, second int64) {
	resp, err := RECDCil.Grant(context.TODO(), second)
	if err != nil {
		print(err)
	}
	_, err = RECDCil.Put(context.TODO(), key, value, client.WithLease(resp.ID))
	if err != nil {
		print(err)
	}
	// revoking lease expires the key attached to its lease ID
	_, err = RECDCil.Revoke(context.TODO(), resp.ID)
	if err != nil {
		print(err)
	}

	gresp, err := RECDCil.Get(context.TODO(), "foo")
	if err != nil {
		print(err)
	}
	fmt.Println("number of keys:", len(gresp.Kvs))

}

func keep() {
	resp, err := RECDCil.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Println(err)
	}

	_, err = RECDCil.Put(context.TODO(), "foo", "bar", client.WithLease(resp.ID))
	if err != nil {
		fmt.Println(err)
	}

	// the key 'foo' will be kept forever
	ch, kaerr := RECDCil.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		fmt.Println(kaerr)
	}

	ka := <-ch
	if ka != nil {
		fmt.Println("ttl:", ka.TTL)
	} else {
		fmt.Println("Unexpected NULL")
	}
}
