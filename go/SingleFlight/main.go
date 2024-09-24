package main

import (
	"golang.org/x/sync/singleflight"
	"log"
	"time"
)

type User struct {
	Name string
	Age  int
}

var Cache map[string]User
var g *singleflight.Group

const clients = 5

func init() {
	Cache = make(map[string]User)
	g = new(singleflight.Group)
}
func GetUser(name string) User {
	if user, ok := Cache[name]; ok {
		return user
	}
	// get user from db
	return GetUserFromDBSingleFlight(name)
}
func GetUserFromDBSingleFlight(name string) User {
	// use singleflight to avoid duplicate request
	v1, _, _ := g.Do("user", func() (interface{}, error) {
		return GetUserFromDB(name), nil
	})
	return v1.(User)
}
func GetUserFromDB(name string) User {
	log.Println("get user from db")
	// time-consuming operation
	time.Sleep(100 * time.Millisecond)
	return User{
		Name: "huizhou92",
		Age:  32,
	}
}

func main() {
	for i := 0; i < clients; i++ {
		go func() {
			log.Println(GetUser("huizhou92"))
		}()
	}
	time.Sleep(1 * time.Second)
}

// list2: user doChan
func doChan(name string) User {

	result := g.DoChan("user", func() (interface{}, error) {
		data := GetUserFromDB(name)
		return data, nil
	})
	select {
	case v := <-result:
		log.Println("get user from db")
		return v.Val.(User)
	case <-time.After(1 * time.Second):
		log.Println("timeout")
		return User{}
	}
}
