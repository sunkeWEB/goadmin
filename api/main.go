package main

import (
	"fmt"
	"go-admin/cmd"
)

// @title go-admin API
// @version 1.0.1
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description 添加qq群: 74520518 进入技术交流群 请备注，谢谢！
// @license.name MIT
// @license.url https://github.com/wenjianzhang/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

//func main() {
//	configName := "settings"
//
//
//	config.InitConfig(configName)
//
//	gin.SetMode(gin.DebugMode)
//	log.Println(config.DatabaseConfig.Port)
//
//	err := gorm.AutoMigrate(orm.Eloquent)
//	if err != nil {
//		log.Fatalln("数据库初始化失败 err: %v", err)
//	}
//
//	if config.ApplicationConfig.IsInit {
//		if err := models.InitDb(); err != nil {
//			log.Fatal("数据库基础数据初始化失败！")
//		} else {
//			config.SetApplicationIsInit()
//		}
//	}
//
//	r := router.InitRouter()
//
//	defer orm.Eloquent.Close()
//
//	srv := &http.Server{
//		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
//		Handler: r,
//	}
//
//	go func() {
//		// 服务连接
//		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			log.Fatalf("listen: %s\n", err)
//		}
//	}()
//	log.Println("Server Run ", config.ApplicationConfig.Host+":"+config.ApplicationConfig.Port)
//	log.Println("Enter Control + C Shutdown Server")
//	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
//	quit := make(chan os.Signal)
//	signal.Notify(quit, os.Interrupt)
//	<-quit
//	log.Println("Shutdown Server ...")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	if err := srv.Shutdown(ctx); err != nil {
//		log.Fatal("Server Shutdown:", err)
//	}
//	log.Println("Server exiting")
//}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

type People interface {
	Eat()
	Run()
}

type Animal interface {
	Fly()
}

type ChinaPeople struct {
	Name string
}

func (receiver ChinaPeople) Run() {
	fmt.Printf("%s 正在跑步\n", receiver.Name)
}

func (receiver ChinaPeople) Eat() {
	fmt.Printf("%s 正在吃东西\n", receiver.Name)
}

func (receiver ChinaPeople) Fly() {
	fmt.Printf("%s 正在吃东西\n", receiver.Name)
}

func main() {
	cmd.Execute()
	//var p People = ChinaPeople{"孙科"}

	//b:=p.(Animal)
	//b.Fly()
	//fmt.Printf("%s\n",p)

	//b := [...]int{1, 2}
	//months := []string{1: "January", 12: "December"}
	//months=append(months,"asa")
	//c:=months[0:1]
	//data,err:=json.MarshalIndent(months,""," ")
	//if err!=nil{
	//
	//}
	//print(len(months),len(b),len(c))
	//fmt.Printf("%s\n",data)

	//const jsonStream = `{"Name": "Ed", "Text": "Knock knock."}`
	//
	//type Message struct {
	//	Name, Text string
	//}
	//
	//var m Message
	//
	//// 用json.NewDecoder
	//dec := json.NewDecoder(strings.NewReader(jsonStream))
	//dec.Decode(&m)
	//fmt.Printf("%s: %s\n", m.Name, m.Text)
	//
	//// 用 json.Unmarshal
	//json.Unmarshal([]byte(jsonStream), &m)
	//fmt.Printf("%s: %s\n", m.Name, m.Text)
	//
	//fmt.Println("done")

}
