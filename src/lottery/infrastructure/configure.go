package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/imdario/mergo"
	"io/ioutil"
	"os"
	"reflect"
)

type Configure struct {
	Env     string
	Binding string
	Path    string
	Value   struct {
		V1 string
		V2 string
		V3 int
	}
}

func LoadConfig() Configure {
	var buff bytes.Buffer
	dir, _ := os.Getwd()
	buff.WriteString(dir)
	buff.WriteString("/settings/config.json")
	raw, err := ioutil.ReadFile(buff.String())
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("quit process")
		os.Exit(1)
	}
	expandContent := []byte(os.ExpandEnv(string(raw)))
	fmt.Println(string(os.ExpandEnv(string(raw))))
	var val interface{}
	var val2 interface{}
	json.Unmarshal(expandContent, &val)
	fmt.Printf("1:value=%v\n", val)
	fmt.Printf("2:value=%v\n", val.(map[string]interface{})["Value"].(map[string]interface{})["V3"])
	fmt.Printf("3:TypeOf(v)=%v\n", reflect.TypeOf(val))
	fmt.Printf("4:value=%v\n", reflect.TypeOf(val.(map[string]interface{})["Value"]))
	fmt.Printf("5:value=%v\n", reflect.TypeOf(val.(map[string]interface{})["Value"].(map[string]interface{})["V3"]))

	if err := mergo.Map(&val2, val); err != nil {
	}
	fmt.Printf("1:value=%v\n", val2)

	var conf Configure
	json.Unmarshal(expandContent, &conf)

	return conf
}
