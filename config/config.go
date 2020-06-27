package config

//go:generate go run ../cmd/binclude

import (
	"github.com/lu4p/binclude"
	"glog/models"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var globalCong *models.Config

// Conf 获取全局的配置文件
func Conf() *models.Config {
	return globalCong
}

func init() {
	binclude.Include("../assets")

	conf := models.Config{}

	f, err := BinFS.Open("../assets/config.yml")

	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bs, &conf)

	if err != nil {
		panic(err)
	}

	globalCong = &conf
}
