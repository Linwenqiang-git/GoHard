package orm

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-xorm/xorm"
	Repo "github.linwenqiang.com/GinStudy/Reponsetory"
	tool "github.linwenqiang.com/GinStudy/Tool"
	"gopkg.in/ini.v1"
)

type XormDbContext struct {
}

func initMySQLConfig(cfg *ini.File) (*Repo.MySqlDbConfig, error) {
	port, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		return nil, err
	}
	show_sql, err := cfg.Section("mysql").Key("show_sql").Bool()
	if err != nil {
		return nil, err
	}
	return &Repo.MySqlDbConfig{
		Ip:       cfg.Section("mysql").Key("ip").String(),
		Port:     port,
		User:     cfg.Section("mysql").Key("user").String(),
		Password: cfg.Section("mysql").Key("password").String(),
		Database: cfg.Section("mysql").Key("database").String(),
		Charset:  cfg.Section("mysql").Key("charset").String(),
		Show_sql: show_sql,
	}, nil
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
	cfg, err := ini.Load("./AppSetting/DevConfig.ini")
	if err != nil {
		fmt.Println("读取配置文件失败")
	}
	config, err := initMySQLConfig(cfg)
	if err != nil {
		fmt.Println("装载配置信息失败")
	}
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
