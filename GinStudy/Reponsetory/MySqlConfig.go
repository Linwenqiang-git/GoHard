package reponsetory

//数据库连接配置
type MySqlDbConfig struct {
	Ip       string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
	Show_sql bool
}
