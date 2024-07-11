package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// https://github.com/go-sql-driver/mysql/blob/v1.8.1/dsn.go#L37
type Config struct {
	User   string `json:"user"`    // Username
	Passwd string `json:"passwd"`  // Password (requires User)
	Net    string `json:"net"`     // Network (e.g. "tcp", "tcp6", "unix". default: "tcp")
	Addr   string `json:"addr"`    // Address (default: "127.0.0.1:3306" for "tcp" and "/tmp/mysql.sock" for "unix")
	DBName string `json:"db_name"` // Database name
	// 。。。。。
}

func InitConfig() *Config {
	cfg := &Config{}
	keys := []string{"CONFIG_MYSQL_USER", "CONFIG_MYSQL_PASSWD", "CONFIG_MYSQL_NET", "CONFIG_MYSQL_ADDR", "CONFIG_MYSQL_DB_NAME"}
	for _, key := range keys {
		if env, exist := os.LookupEnv(key); exist {
			switch key {
			case "CONFIG_MYSQL_USER":
				cfg.User = env
			case "CONFIG_MYSQL_PASSWORD":
				cfg.Passwd = env
			case "CONFIG_MYSQL_NET":
				cfg.Net = env
			case "CONFIG_MYSQL_ADDR":
				cfg.Addr = env
			case "CONFIG_MYSQL_DB_NAME":
				cfg.DBName = env
			}
		}
	}
	return cfg
}
func InitConfig2() *Config {
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_MYSQL_%s", strings.ToUpper(v))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &config
}

func main() {
	os.Setenv("CONFIG_MYSQL_USER", "root")
	os.Setenv("CONFIG_MYSQL_PASSWD", "123456")
	os.Setenv("CONFIG_MYSQL_NET", "tcp")
	os.Setenv("CONFIG_MYSQL_ADDR", "localhost:3306")
	os.Setenv("CONFIG_MYSQL_DB_NAME", "test")
	cfg := InitConfig2()
	fmt.Println(cfg)
}
