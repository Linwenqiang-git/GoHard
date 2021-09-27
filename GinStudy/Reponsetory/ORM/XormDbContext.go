package orm

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-xorm/xorm"
	GlocConfig "github.linwenqiang.com/GinStudy/AppSetting"
	Repo "github.linwenqiang.com/GinStudy/Reponsetory"
	tool "github.linwenqiang.com/GinStudy/Tool"
)

type XormDbContext struct {
}

func initMysqlStr(option *Repo.MySqlDbConfig) string {
	var constr string = option.User + ":"
	constr += option.Password + "@"
	constr += "tcp"
	constr += "(" + option.Ip + ":" + strconv.Itoa(option.Port) + ")/"
	constr += option.Database
	return constr
}

var Engine *xorm.Engine = nil

//创建数据库上下文
func newDbContext() *xorm.Engine {
	if Engine != nil {
		return Engine
	}
	config := GlocConfig.GetConfig().MysqlConfig
	Engine, err := xorm.NewEngine("mysql", initMysqlStr(config))
	if err != nil {
		fmt.Println("xorm连接数据库失败：", err)
	}
	Engine.ShowSQL(true)
	//将sql的信息打印到日志
	f, err := os.Create("sql.log")
	tool.PrintInfoError(err)
	Engine.SetLogger(xorm.NewSimpleLogger(f))
	return Engine
}

//===================================仓储层工厂方法===================================
func NewOrderDao() *OrderDao {
	engine := newDbContext()
	return &OrderDao{engine: engine}
}
