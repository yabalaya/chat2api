package main

import (
	"chat2api/config"
	"chat2api/v1/chat"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	go ghttp.StartPProfServer(8199)
	s := g.Server()
	s.SetPort(config.PORT)
	s.BindHandler("/", Index)
	chatGroup := s.Group("/v1/chat")
	chatGroup.Middleware(MiddlewareCORS)

	chatGroup.ALL("/completions", chat.Completions)
	chatGroup.ALL("/gpt4v", chat.Gpt4v)
	chatGroup.ALL("/dalle3", chat.Dalle3)
	s.Run()
}
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func Index(r *ghttp.Request) {
	r.Response.Write("Hello Xyhelper! This is chat2api.")
}
