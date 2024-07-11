package main

import (
	"reflect"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	var config *Config
	for i := 0; i < b.N; i++ {
		config = new(Config)
	}
	_ = config
}

func BenchmarkReflectNew(b *testing.B) {
	var config *Config
	typ := reflect.TypeOf(Config{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config, _ = reflect.New(typ).Interface().(*Config)
	}
	_ = config
}

func BenchmarkFieldSet(b *testing.B) {
	config := new(Config)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config.Net = "tcp4"
		config.Addr = "127.0.0.1:3306"
		config.Passwd = "123456"
		config.User = "admin"
	}
}

func BenchmarkFieldSetFieldByName(b *testing.B) {
	config := new(Config)
	value := reflect.ValueOf(config).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value.FieldByName("Net").SetString("tcp4")
		value.FieldByName("Addr").SetString("127.0.0.1:3306")
		value.FieldByName("Passwd").SetString("123456")
		value.FieldByName("User").SetString("admin")
	}
}
func BenchmarkFieldSetFieldByNameCache(b *testing.B) {
	config := new(Config)
	typ := reflect.TypeOf(Config{})
	value := reflect.ValueOf(config).Elem()
	cache := make(map[string]int)
	for i := 0; i < typ.NumField(); i++ {
		cache[typ.Field(i).Name] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value.Field(cache["Net"]).SetString("tcp4")
		value.Field(cache["Addr"]).SetString("127.0.0.1:3306")
		value.Field(cache["Passwd"]).SetString("123456")
		value.Field(cache["User"]).SetString("admin")
	}
}

func BenchmarkFieldSetField(b *testing.B) {
	config := new(Config)
	value := reflect.Indirect(reflect.ValueOf(config))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value.Field(0).SetString("tcp4")
		value.Field(1).SetString("127.0.0.1:3306")
		value.Field(2).SetString("123456")
		value.Field(3).SetString("admin")
	}
}
