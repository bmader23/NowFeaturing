package main

import (
	"github.com/bmader23/nowfeaturing/handler"
	"github.com/bmader23/nowfeaturing/repository"
	"github.com/bmader23/nowfeaturing/service"
	"github.com/gin-gonic/gin"
)

func main() {
	gingine := gin.Default()

	handler.NewFlagHandler(gingine, service.FlagService{
		Rep: repository.FileFlagRepository{
			PathRoot: "./testdata/",
		},
	})

	gingine.Run("0.0.0.0:8090")
}
