package orm

import (
	"fmt"
	"strconv"

	"github.com/go-xorm/xorm"
	. "github.linwenqiang.com/GinStudy/Reponsetory"
	"gopkg.in/ini.v1"
)

type XormDbContext struct {
}

func initMySQLConfig(cfg *ini.File) (*MySqlDbConfig, error) {
	port, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		return nil, err
	}
	show_sql, err := cfg.Section("mysql").Key("show_sql").Bool()
	if err != nil {
		return nil, err
	}
	return &MySqlDbConfig{
		Ip:       cfg.Section("mysql").Key("ip").String(),
		Port:     port,
		User:     cfg.Section("mysql").Key("user").String(),
		Password: cfg.Section("mysql").Key("password").String(),
		Database: cfg.Section("mysql").Key("database").String(),
		Charset:  cfg.Section("mysql").Key("charset").String(),
		Show_sql: show_sql,
	}, nil
}
func initMysqlStr(option *MySqlDbConfig) string {
	var constr string = option.User + ":"
	constr += option.Password + "@"
	constr += "@tcp"
	constr += "(" + option.Ip + ":" + strconv.Itoa(option.Port) + ")/"
	constr += option.Database
	return constr
}

var Engine *xorm.Engine = nil

func init() {
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
}
