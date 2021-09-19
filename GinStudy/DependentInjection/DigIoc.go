package dependentinjection

import (
	"database/sql"
	"fmt"
	"strconv"

	. "github.linwenqiang.com/GinStudy/InitSql"
	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
	"go.uber.org/dig"
	"gopkg.in/ini.v1"
)

//加载配置文件
func initConfig() (*ini.File, error) {
	cfg, err := ini.Load("./AppSetting/DevConfig.ini")
	return cfg, err
}

//初始化mysql配置信息
func initMySQLConfig(cfg *ini.File) (*MySqlDbConfig, error) {
	port, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		return nil, err
	}
	return &MySqlDbConfig{
		Ip:       cfg.Section("mysql").Key("ip").String(),
		Port:     port,
		User:     cfg.Section("mysql").Key("user").String(),
		Password: cfg.Section("mysql").Key("password").String(),
		Database: cfg.Section("mysql").Key("database").String(),
	}, nil
}

//打开mysql连接
func initMysqlDb(option *MySqlDbConfig) (*sql.DB, error) {
	var constr string = option.User + ":"
	constr += option.Password + "@"
	constr += "@tcp"
	constr += "(" + option.Ip + ":" + strconv.Itoa(option.Port) + ")/"
	constr += option.Database
	return sql.Open("mysql", constr)
}

func PrintInfo(config Config) {

	fmt.Println("=========== mysql section ===========")
	fmt.Println("mysql ip:", config.MySQL.Ip)
	fmt.Println("mysql port:", config.MySQL.Port)
	fmt.Println("mysql user:", config.MySQL.User)
	fmt.Println("mysql password:", config.MySQL.Password)
	fmt.Println("mysql db:", config.MySQL.Database)
}
func AccessDb(context DbContext) {
	fmt.Println("进入到AccessDb方法")
	resutlt := context.Query("select * from userdto", &Dto.UserDto{})
	fmt.Printf("%v", resutlt)
}

func PrintDI_Error(err error) {
	if err != nil {
		fmt.Println("dig 注入出错：", err)
	}
}

//基础服务配置
func ServiceBaseConfiguration() *dig.Container {
	container := dig.New()

	container.Provide(initConfig)
	container.Provide(initMySQLConfig)
	container.Provide(initMysqlDb)

	//container.Invoke(AccessDb)
	return container
}
