// 业务层（对齐 backend 的 apps/ 角色）：health 与 OpenAPI schema 端点。
// 边界约定：internal/* 不反向依赖 cmd/*（将来经 golangci-lint depguard 机检）。
package app

import (
	_ "embed"
	"encoding/json"
	"net/http"
)

//go:embed openapi.json
var openapiSchema []byte

// healthResponse 对齐 backend HealthResponseSerializer 的字段形状（db 为内存存储连通性占位，恒 ok）。
type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	DB      string `json:"db"`
}

// Handler 返回业务路由（ServeMux 1.22+ 方法+路径模式）。
func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/health", health)
	mux.HandleFunc("GET /api/v1/health/", health)
	mux.HandleFunc("GET /api/v1/schema", schema)
	mux.HandleFunc("GET /api/v1/schema/", schema)
	return mux
}

func health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{Status: "ok", Service: "todo-backend-go", DB: "ok"})
}

func schema(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(openapiSchema)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, `{"error":"encode failed"}`, http.StatusInternalServerError)
	}
}
