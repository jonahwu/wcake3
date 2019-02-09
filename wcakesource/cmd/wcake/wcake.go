package main

import (
	"flag"
	//"fmt"
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/wcake/cmd/handler"
	"github.com/wcake/go/jutils"
	"go.etcd.io/etcd/clientv3"
	"time"
	//	"github.com/labstack/echo/middleware"
)

func main() {
	flag.Parse()
	// here we setup the default stdout for glog
	flag.Lookup("logtostderr").Value.Set("true")
	datalog := jutils.GetHello()
	glog.Info(datalog)
	//fmt.Println(data)
	glog.Info("now into data")
	//glog.Error("glog in errorla")

	//db, err := mgo.Dial("mongo")
	//if err != nil {
	//		e.Logger.Fatal(err)
	//	}
	//	h := &handler.Handler{DB: db}
	//	h := &Handler
	//	var h Handler
	// replace the following and by using db connection here
	// and send the pointer to handler
	//see https://github.com/petronetto/echo-mongo-api/blob/master/handler/handler.go
	//con := 5
	//h := &handler.Handler{CON: &con}
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
	//_, err = cli.Put(ctx, "foo3", "bar3")
	//data, _ := cli.Get(ctx, "foo3")
	//fmt.Println(string(data.Kvs[0].Key), string(data.Kvs[0].Value))
	//fmt.Println("get size", data.Count)

	//defer cli.Close() // make sure to close the client

	h := &handler.Handler{ECon: cli, ECtx: ctx}

	e := echo.New()

	e.GET("/*", h.GetHostName)
	e.GET("/sleep", h.GetToSleep)
	e.GET("/sleepinf", h.GetToSleepInf)
	e.GET("/getstore", h.GetStore)
	e.GET("/uploadstore", h.UploadStore)
	e.POST("/buy", h.Buy)
	e.GET("/buyid", h.GetBuyID)
	e.POST("/getbuy", h.GetBuy)
	e.POST("/teststoreget", h.TestStoreGet)
	e.Logger.Fatal(e.Start(":8000"))

	glog.Flush()
	select {}
}
