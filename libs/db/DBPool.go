/**
 * 数据库链接池
 *
 * @param addTime 2022-05-07
 * @param author  ChengBo
 */
package db

import (
	"GoEngine/libs/helper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

//定义链接池配置Map
var EnvDBMap = make(map[string]*gorm.DB, 8)

//初始化连接池链接
func init() {
	//提取配置文件信息
	config := helper.GetConfig()

	//检查是否有DB链接配置
	if len(config.MySQLConfig) == 0 {
		return
	}

	//循环初始化链接DB
	for _, dbConfig := range config.MySQLConfig {
		//检查是否初始化过
		if _, ok := EnvDBMap[dbConfig.Key]; ok {
			continue
		}

		//链接DB
		dbLink, err := gorm.Open(dbConfig.Dirver, fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
			dbConfig.UserName,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.DataBase,
			dbConfig.CharSet,
		))
		if err != nil {
			panic(err)
		}

		//设置链接配置
		//dbLink.LogMode(true)
		dbLink.DB().SetMaxOpenConns(10)
		dbLink.DB().SetMaxIdleConns(10)
		dbLink.DB().SetConnMaxLifetime(time.Second * 300)

		//检查链接配置
		if err = dbLink.DB().Ping(); err != nil {
			panic(err)
		}

		//设置链接配置
		EnvDBMap[dbConfig.Key] = dbLink
	}

	return
}

//封装一个通过名字获取数据库hander的方法。
func GetDbLink(name string) *gorm.DB {
	//返回指定key的数据实例
	if dbLink, ok := EnvDBMap[name]; ok {
		return dbLink
	}

	//打印日志
	panic(fmt.Sprintf("[ExceptionMsg] GetDbLink Not Exists Name: %s", name))

	return nil
}
