package main

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"
	 _ "github.com/astaxie/beego/session/redis"
	"quickstart/controllers"
	"github.com/astaxie/beego/logs"
	"time"
	"errors"
	"github.com/astaxie/beego/toolbox"
)

type DatabaseCheck struct {
}

func (dc *DatabaseCheck) Check() error {
	if 1==2 {
		return nil
	} else {
		return errors.New("can't connect database")
	}
}

func main() {
	formate := "rotate_day." + time.Now().Add(-24*time.Hour).Format("2006-01-02") + ".log"
	logs.SetLogger(logs.AdapterFile,`{"filename":"test.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`,formate)
	logs.EnableFuncCallDepth(true)
	beego.ErrorController(&controllers.ErrorController{})

	l := logs.GetLogger()
	l.Println("this is a message of http")
	logs.GetLogger("ORM").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
	toolbox.AddHealthCheck("database",&DatabaseCheck{})

	beego.SetStaticPath("/files","upload")

	beego.Run()


}

