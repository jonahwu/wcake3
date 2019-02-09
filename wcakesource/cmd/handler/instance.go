package handler

import (
	"context"
	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/tidwall/gjson"
	"github.com/wcake/go/jutils"
	"go.etcd.io/etcd/clientv3"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//import (
//	"gopkg.in/mgo.v2"
//)

//type (
// Handler - Handle with DB connection
//	Handler struct {
//			DB *mgo.Session
//			}
//		)

type Handler struct {
	//	DB *Mongo
	//CON *int
	ECon *clientv3.Client
	ECtx context.Context
}

func (h *Handler) GetHostName(c echo.Context) error {
	hn := jutils.GetHostName()
	glog.Info(hn)
	return c.JSON(http.StatusCreated, hn)
}

//curl -H "SLEEPTIME: 3"  http://localhost:8080/sleep  -v
func (h *Handler) GetToSleep(c echo.Context) error {
	//sleeptime := c.Request().Header.Get("SLEEPTIME")
	sleeptime := c.Request().Header.Get("SLEEPTIME")
	ist, _ := strconv.Atoi(sleeptime)
	glog.Info(sleeptime)
	time.Sleep(time.Duration(ist) * time.Second)

	return c.JSON(http.StatusOK, sleeptime)
}

func (h *Handler) GetToSleepInf(c echo.Context) error {
	//sleeptime := c.Request().Header.Get("SLEEPTIME")
	sleeptime := "10000000"
	ist, _ := strconv.Atoi(sleeptime)
	glog.Info(sleeptime)
	time.Sleep(time.Duration(ist) * time.Second)

	return c.JSON(http.StatusOK, sleeptime)
}

func (h *Handler) GetStore(c echo.Context) error {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	Value := c.Request().Header.Get("key")
	glog.Info("get from header", Value)
	//if data, err := h.ECon.Get(context.Background(), "foo6"); err != nil {
	if data, err := h.ECon.Get(ctx, "foo6"); err != nil {
		glog.Error(err)
	} else {
		glog.Info("here is the data we get", data)
		glog.Info(string(data.Kvs[0].Key), string(data.Kvs[0].Value))
		glog.Info("get size", data.Count)
	}
	cancel()
	return c.JSON(http.StatusOK, "")
}
func (h *Handler) UploadStore(c echo.Context) error {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	if _, err := h.ECon.Put(ctx, "foo6", "bar6"); err != nil {
		glog.Info(err)
	}
	glog.Info("upload data stored")
	cancel()
	return c.JSON(http.StatusOK, "")
}

func (h *Handler) GetBuy(c echo.Context) error {
	if body, err := ioutil.ReadAll(c.Request().Body); err != nil {
		glog.Error("something wrong")
	} else {
		buyid := gjson.Get(string(body), "buyid").String()
		buyvalue := h.GetEtcdKV(buyid)
		glog.Info("show buyvalue:", buyvalue)
		return c.JSON(http.StatusOK, buyvalue)
	}
	return c.JSON(http.StatusOK, "")
}
func (h *Handler) Buy(c echo.Context) error {
	//var body []byte
	if body, err := ioutil.ReadAll(c.Request().Body); err != nil {
		glog.Error(err)
	} else {
		glog.Info(string(body))
		userid := gjson.Get(string(body), "buy.userinfo.userid").String()
		glog.Info("show userid:", userid)
		buyid := gjson.Get(string(body), "buy.buyinfo.buyid").String()
		glog.Info("show buyid:", buyid)
		buyidK := jutils.CombineString("buyid-", buyid)
		h.StoreEtcdKV(buyidK, buyid)
		h.StoreEtcdKV(buyid, string(body))
		return c.JSON(http.StatusOK, buyid)

	}
	return c.JSON(http.StatusOK, "")
}
func (h *Handler) PreBuy(c echo.Context) error {

	//sleeptime := c.Request().Header.Get("SLEEPTIME")
	var body []byte
	body, err := ioutil.ReadAll(c.Request().Body)
	glog.Info(string(body))
	glog.Info(err)
	glog.Info(gjson.Get(string(body), "buy.userid"))
	items := gjson.Get(string(body), "buy.items")
	glog.Info(items)
	//	{"buy":{"userid":"aaaabbbbccc","items":[{"itemid":"itemaaaabbb", "item":"nb1400"}, {"itemid":"item999333", "item":"nb993"}]}}
	for _, data := range items.Array() {
		glog.Info("show items", data, "itemid:", data.Get("itemid"), "itemname:", data.Get("item"))
	}
	/*
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		if _, err := h.ECon.Put(ctx, "foo6", "bar6"); err != nil {
			glog.Info(err)
		}
		glog.Info("upload data stored")
		cancel()
	*/
	return c.JSON(http.StatusOK, "")
}

func (h *Handler) GetBuyID(c echo.Context) error {
	//return buyid and timestamp
	buyuuid := getBuyID()
	buytime := getBuyTime()
	buy := Buy{
		BuyID:   buyuuid,
		BuyTime: buytime,
	}
	return c.JSON(http.StatusOK, buy)
}

func (h *Handler) TestStoreGet(c echo.Context) error {
	//var body []byte
	if body, err := ioutil.ReadAll(c.Request().Body); err != nil {
		glog.Error(err)
	} else {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		if _, err := h.ECon.Put(ctx, "testbuy-id", string(body)); err != nil {
			glog.Error(err)
		}
		cancel()
	}
	glog.Info("upload data stored")

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	if data, err := h.ECon.Get(ctx, "testbuy-id"); err != nil {
		glog.Error(err)
	} else {
		glog.Info(data)
		glog.Info("key:", string(data.Kvs[0].Key))
		glog.Info("value:", string(data.Kvs[0].Value))
		gjsonValue := gjson.ParseBytes(data.Kvs[0].Value)
		userid := gjsonValue.Get("buy.userid")
		glog.Info("pase result and get userid from gjson:", userid)
	}
	cancel()

	return c.JSON(http.StatusOK, "")

}
