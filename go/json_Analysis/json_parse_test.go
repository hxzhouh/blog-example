package json_parse

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"testing"
)

const jsonStr = `{"squadName":"Super hero squad","homeTown":"Metro City","formed":2016,"secretBase":"Super tower","active":true,"members":[{"name":"Molecule Man","age":29,"secretIdentity":"Dan Jukes","powers":["Radiation resistance","Turning tiny","Radiation blast"]},{"name":"Madame Uppercut","age":39,"secretIdentity":"Jane Wilson","powers":["Million tonne punch","Damage resistance","Superhuman reflexes"]},{"name":"Eternal Flame","age":1000000,"secretIdentity":"Unknown","powers":["Immortality","Heat Immunity","Inferno","Teleportation","Interdimensional travel"]}]}`

type JsonParse struct {
	SquadName  string   `json:"squadName"`
	HomeTown   string   `json:"homeTown"`
	Formed     int      `json:"formed"`
	SecretBase string   `json:"secretBase"`
	Active     bool     `json:"active"`
	Members    []Member `json:"members"`
}
type Member struct {
	Name           string   `json:"name"`
	Age            int      `json:"age"`
	SecretIdentity string   `json:"secret"`
	Powers         []string `json:"powers"`
}

func ParseJson() string {
	var j JsonParse
	err := json.Unmarshal([]byte(jsonStr), &j)
	if err != nil {
		panic(err)
	}
	return j.SquadName
}

// Parse By map[string]interface
func ParseByMap() string {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		panic(err)
	}
	if _, ok := result["squadName"]; ok {
		return result["squadName"].(string)
	}
	return ""
}

func ParseByGjson() string {
	if gjson.Get(jsonStr, "squadName").Exists() {
		return gjson.Get(jsonStr, "squadName").Str
	}
	return ""
}

func BenchmarkParseJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseJson()
	}
}

func BenchmarkParseByMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseByMap()
	}
}

func BenchmarkParseByGjson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseByGjson()
	}
}
