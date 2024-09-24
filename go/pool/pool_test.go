package pool

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Title string `json:"title"`
}

func (s *User) reset() {
	s.Name = ""
	s.Title = ""
	s.Age = 0
}

var UserPool = sync.Pool{
	New: func() interface{} {
		return new(User)
	},
}
var jsonString = `{"name":"test","age":18,"title":"title"}`

func BenchmarkSyncPool(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		u := UserPool.Get().(*User)
		_ = json.Unmarshal([]byte(jsonString), u)
		//u.reset()
		UserPool.Put(u)
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		u := &User{}
		_ = json.Unmarshal([]byte(jsonString), u)
	}
}

func TestUser(t *testing.T) {
	u := &User{
		Name:  "test",
		Age:   18,
		Title: "title",
	}
	b, err := json.Marshal(u)
	fmt.Println(string(b), err)
}
