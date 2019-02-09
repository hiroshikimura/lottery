package main

import (
	"bytes"
	"fmt"
	_ "github.com/fsnotify/fsnotify"
	"github.com/mitchellh/cli"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"lottery/infrastructure"
	"os"
)

func readConfigContent(filename string) []byte {
	var buff bytes.Buffer
	dir, _ := os.Getwd()
	buff.WriteString(dir)
	buff.WriteString(filename)
	raw, err := ioutil.ReadFile(buff.String())
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("quit process")
		os.Exit(1)
	}
	fmt.Printf("%s", os.ExpandEnv(string(raw)))
	return []byte(os.ExpandEnv(string(raw)))
}

func viperInit() {
	viper.SetConfigType("json")
	err := viper.ReadConfig(bytes.NewBuffer(readConfigContent("/settings/config.json"))) // Find and read the config file
	if err != nil {                                                                      // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.MergeConfig(bytes.NewBuffer(readConfigContent("/settings/local_config.json")))
	fmt.Println("config filename=" + viper.ConfigFileUsed())
	/*
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			viper.ReadInConfig()
		})
	*/
	fmt.Println("PATH=", viper.Get("Path"))
	fmt.Println("Value.V2=", viper.Get("Value.V2"))
}

// https://qiita.com/uchiko/items/5e5cda98ecb510671e56
func main() {
	viperInit()
	// コマンドの名前とバージョンを指定
	c := cli.NewCLI("app", "1.0.0")
	// サブコマンドの引数を指定
	c.Args = os.Args[1:]

	// サブコマンド文字列 と コマンド実装の対応付け
	c.Commands = map[string]cli.CommandFactory{
		"launch": func() (cli.Command, error) {
			return &infrastructure.Launch{}, nil
		},
		"migrate": func() (cli.Command, error) {
			return &infrastructure.Migarte{}, nil
		},
		"test": func() (cli.Command, error) {
			return &infrastructure.Test{}, nil
		},
	}

	// コマンド実行
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
