package main

import (
	"encoding/json"
	"fmt"
	v2_json "github.com/go-json-experiment/json"
	"math"
	"time"
)

type Post struct {
	Id         int64           `json:"id,omitempty,omitzero"`
	CreateTime time.Time       `json:"create_time,omitempty,omitzero"`
	TagList    []Tag           `json:"tag_list,omitempty"`
	Name       string          `json:"name,omitempty"`
	Score      ScoreType       `json:"score,omitempty,omitzero"`
	Category   Category        `json:"category,omitempty,omitzero"`
	LikePost   map[string]Post `json:"like,omitempty"`
}
type ScoreType float64

func (s ScoreType) IsZero() bool {
	return s < math.MinInt64
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Category struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
}

func main() {
	v1String, _ := json.Marshal(new(Post))
	fmt.Println(string(v1String))
	v2String, _ := v2_json.Marshal(new(Post))
	fmt.Println(string(v2String))
}
