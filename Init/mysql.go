package Init

import (
	"dousheng/global"
	"fmt"

	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//利用配置信息连接mysql
func MysqlDB() {
	mysqlInfo := global.Settings.Mysqlinfo
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.DBName)
	//注册数据库及全局变量
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.DB = db
	color.Green("Database Connection Initialize")
}
