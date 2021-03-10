package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type ViperInPara struct {
	Test1 string `mapstructure:"test1"`
}

type ViperPara struct {
	Name    string            `mapstructure:"Name"`
	Age     int               `mapstructure:"Age"`
	Time    time.Time         `mapstructure:"Time"`
	Dur     time.Duration     `mapstructure:"Dur"`
	Contain ViperInPara       `mapstructure:"Contain"`
	ListStr []string          `mapstructure:"ListStr"`
	mapStr  map[string]string `mapstructure:"mapStr"`
}

func ReadCfg() {
	os.RemoveAll("./config")
	os.MkdirAll("./config", os.ModeDir)

	para := ViperPara{
		Name: "test",
	}

	data, _ := json.Marshal(&para)
	ioutil.WriteFile("./config/config.json", data, os.ModePerm)

	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	// viper.SetConfigType("yaml")           // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./config") // 还可以在工作目录中查找配置

	// viper.SetDefault("ContentDir", "content")
	// viper.SetDefault("LayoutDir", "layouts")
	// viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("file is not exist")
		} else {
			log.Fatalln("other err", err)
		}
	} else {
		out := ViperPara{}
		err := viper.Unmarshal(&out)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
		fmt.Printf("%+v\n", out)
	}
}

func main() {
	fmt.Println("viper test in")
	ReadCfg()
}
