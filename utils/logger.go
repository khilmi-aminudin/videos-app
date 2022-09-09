package utils

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func SetLogOutputFile(pathfile ...string) {
	var file string
	if pathfile == nil {
		file = "log.log"
	} else {
		file = pathfile[0]
	}
	f, _ := os.Create(file)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
