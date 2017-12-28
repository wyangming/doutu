package http

import (
	"doutu/module/http/controller"
	"fmt"
	"github.com/astaxie/beego"
)

func init() {
	fmt.Println("http module start")
	go run()
}
func run() {
	init_filter()
	init_Router()
	beego.Run()
}

func init_filter() {
}

func init_Router() {
	//index
	indexController := &controller.IndexController{}
	beego.Router("/", indexController, "*:Index")
	//临时生成图片使用
	beego.Router("/genimg", indexController, "*:GenImg")
}
