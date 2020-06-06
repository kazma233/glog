package config

import (
	"glog/models"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var globalCong *models.Config

// Conf 获取全局的配置文件
func Conf() *models.Config {
	return globalCong
}

func init() {
	conf := models.Config{}

	file, err := os.OpenFile("./config/config.yml", os.O_RDONLY, 0755)

	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bs, &conf)

	if err != nil {
		panic(err)
	}

	globalCong = &conf
}
