package appsetting

import (
	cache "github.linwenqiang.com/GinStudy/Cache"
	repo "github.linwenqiang.com/GinStudy/Reponsetory"
	tool "github.linwenqiang.com/GinStudy/Tool"
	"gopkg.in/ini.v1"
)

//全局基础配置
type GlobalConfig struct {
	MysqlConfig *repo.MySqlDbConfig
	RedisConfig *cache.RedisConfig
}

//初始化配置文件
func GetConfig() *GlobalConfig {
	cfg, err := ini.Load("./AppSetting/DevConfig.ini")
	tool.PanicError(err)
	//init mysqlConfig
	port, err := cfg.Section("mysql").Key("port").Int()
	tool.PanicError(err)
	show_sql, err := cfg.Section("mysql").Key("show_sql").Bool()
	tool.PanicError(err)
	mysqlConfig := &repo.MySqlDbConfig{
		Ip:       cfg.Section("mysql").Key("ip").String(),
		Port:     port,
		User:     cfg.Section("mysql").Key("user").String(),
		Password: cfg.Section("mysql").Key("password").String(),
		Database: cfg.Section("mysql").Key("database").String(),
		Charset:  cfg.Section("mysql").Key("charset").String(),
		Show_sql: show_sql,
	}
	//init redis
	redisPort, _ := cfg.Section("redis").Key("port").Int()
	redisDb, _ := cfg.Section("redis").Key("db").Int()
	redisConfig := &cache.RedisConfig{
		Addr:     cfg.Section("redis").Key("addr").String(),
		Password: cfg.Section("redis").Key("password").String(),
		Port:     redisPort,
		Db:       redisDb,
	}
	return &GlobalConfig{RedisConfig: redisConfig, MysqlConfig: mysqlConfig}
}
