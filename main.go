package main


import (
	"glog/config"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	mode := config.Conf().Env
	if strings.Compare(mode, "PROD") == 0 {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	log.Printf("服务启动失败: %s", Router().Run("0.0.0.0:9600"))
}
