package handler

import (
	//"github.com/golang/glog"
	"context"
	"github.com/golang/glog"
	"github.com/tidwall/gjson"
	"github.com/wcake/pkg/jutils"
	"time"
)

type Buy struct {
	BuyID   string `json:"buyid"`
	BuyTime string `json:"buytime"`
}

func getBuyID() string {
	return jutils.GetUuid()
}

func getBuyTime() string {
	return time.Now().Format("20060102150405")
}

func ParseBuyID(body []byte) {
	gjson.Get(string(body), "buy.buyinfo.buyid")
}

func (h *Handler) StoreEtcdKV(key string, value string) {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	if _, err := h.ECon.Put(ctx, key, value); err != nil {
		glog.Info(err)
	}
	glog.Info("upload data stored")
	cancel()
}

func (h *Handler) GetEtcdKV(key string) string {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	if data, err := h.ECon.Get(ctx, key); err != nil {
		glog.Info(err)
	} else {
		cancel()
		glog.Info("get data:", string(data.Kvs[0].Value))
		return string(data.Kvs[0].Value)
	}
	return ""
}
