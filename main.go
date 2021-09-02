package main

import (
	"flag"
	"fmt"
	"github.com/wslio/TicketReported/pkg"
	"log"
)

func main() {
	// 读取app.toml配置文件
	flag.Parse()
	if err := pkg.Init(); err != nil {
		log.Println("conf.Init() err:{}", err)
	}
	mysqlConf := pkg.Conf.MySql
	fmt.Println(mysqlConf.DbName)

	// 使用mysql 查询所有的org， url
	// 使用协程调取所有站的票量
	// 分析汇总，上报数据
}
