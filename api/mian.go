package main

import (
	"dst-search/handlers"
	"embed"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
)

//go:embed dist
var EmbedFS embed.FS
var BindPort int

func main() {
	flag.IntVar(&BindPort, "l", 7777, "监听端口，如： -l 8080 (Listening Port, e.g. -l 8080)")
	flag.Parse()

	r := gin.Default()
	r.POST("/search", handlers.Search)
	r.POST("/detail", handlers.Detail)
	r.Use(static.ServeEmbed("dist", EmbedFS))

	if err := r.Run(fmt.Sprintf(":%d", BindPort)); err != nil {
		panic(fmt.Sprintf("服务器启动失败：%v", err))
	}

}
