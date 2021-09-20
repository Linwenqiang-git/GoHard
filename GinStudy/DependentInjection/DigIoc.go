package dependentinjection

/*
=========================dig容器使用总结=========================
1.注入形式简单，容易理解，轻量级；
2.构造函数注入通过dig.In的方式，需要注意的是，不能传指针类型，指针类型的必须是在容器里的对象才可；
3.只能通过Incoke的形式手动执行方法，当方法需要传参或者获取返回值都无法做到，不够灵活；

*/

import (
	"database/sql"
	"fmt"
	"strconv"

	Dto "github.linwenqiang.com/GinStudy/Model/Dto"
	. "github.linwenqiang.com/GinStudy/Reponsetory"
	"go.uber.org/dig"
	"gopkg.in/ini.v1"
)

type Config struct {
	/*内嵌了dig.In之后，dig会将该类型中的其它字段看成对象的依赖，
	  创建Object类型的对象时，会先将依赖的Arg1/Arg2/Arg3/Arg4创建好。
	  需要注意的是，方法的参数不能传指针类型，因为在容器里是没有这个对象的*/
	dig.In
	MySQL *MySqlDbConfig
}

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
