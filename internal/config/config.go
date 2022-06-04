package config

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
)

type Addr struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Config struct {
	AdminRootPath string
	Name          string `json:"name"`
	Server        Addr   `json:"server"`
	HtmlTemplate  *template.Template
}

var GlobalConfig *Config

func init() {
	GlobalConfig = &Config{}
	GlobalConfig.AdminRootPath = getRootPath()
}

func LoadConfig(fileName *string) *Config {
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("loadConfig: file open error, fileName=%s, err=%v.", fileName, err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("loadConfig: ioutil.ReadAll error, fileName=%s, err=%v.", fileName, err)
	}

	err = json.Unmarshal(data, GlobalConfig)
	if err != nil {
		log.Fatalf("loadConfig: json.Unmarshal error, fileName=%s, err=%v.", fileName, err)
	}
	return GlobalConfig
}

func SetHtmlTemplate(templ *template.Template) {
	GlobalConfig.HtmlTemplate = templ
}

func GetTemplateData(tplName string, data gin.H) string {
	writer := bytes.Buffer{}
	GlobalConfig.HtmlTemplate.ExecuteTemplate(&writer, tplName, data)
	return writer.String()
}

func getRootPath() string {
	_, fileName, line, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("GetRootPath: runtime.Caller is error.")
	}
	rootPath := path.Dir(path.Dir(path.Dir(fileName)))
	log.Printf("GetRootPath: rootPath=%s, fileName=%s, line=%d.", rootPath, fileName, line)
	return rootPath
}
