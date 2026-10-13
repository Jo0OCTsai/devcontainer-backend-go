// 入口层（对齐 backend 的 config/ 角色）：只做装配与监听，业务逻辑在 internal/app。
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Jo0OCTsai/devcontainer-backend-go/internal/app"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8001"
	}
	srv := &http.Server{
		Addr:         addr,
		Handler:      app.Handler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Printf("todo-backend-go listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
