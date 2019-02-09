package jutils

import (
	//"flag"
	//"fmt"
	"github.com/golang/glog"
)

func GetHello() string {
	glog.Info("inside the GetHello")
	inDir := GetDir()
	glog.Info(inDir)
	//fmt.Println("haha")
	return "Hello"
}
