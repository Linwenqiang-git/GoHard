package cache

type RedisConfig struct {
	Addr     string
	Password string
	Port     int
	Db       int
}
