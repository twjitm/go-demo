package server

import (
	"fmt"
	"net/http"
)
import "github.com/gin-gonic/gin"

func Server(host string) {
	r := gin.Default()

	r.GET("/getData", GetData)

	err := r.Run(host)
	if err != nil {
		fmt.Println("start server error")
		return
	}
}

const FilePath = "/data/docker/data.json"

type ContainerData struct {
	ContainerName string `json:"container_name"`
	ContainerCode int32  `json:"container_code"`
	IP            string `json:"ip"`
	BranchName    string `json:"branch_name"`
	GamePort      int32  `json:"game_port"`
	VersionPort   int32  `json:"version_port"`
	GGatePort     int32  `json:"g_gate_port"`
	CCatePort     int32  `json:"c_cate_port"`
	GmPort        int32  `json:"gm_port"`
}

func GetData(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}
