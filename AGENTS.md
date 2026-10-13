# backend-go/AGENTS.md

> 本文件是 todo-backend-**go** 的 agent 工作指引。通用开发流程原则见主仓库根 `AGENTS.md`；
> 与 Python 版 `backend/` 为**栈适配镜像**关系——API 面与分层角色对齐，实现按 Go 惯例。

## 项目概述

todo app 后端的 Go 镜像（骨架阶段）：health 端点 + 静态 OpenAPI schema，内存存储占位。

## 技术栈

- Go（标准库 `net/http`，**零三方依赖**——适配项目的刻意选择，降低首例接入变量）
- lint：**golangci-lint v2**（`.golangci.yml`：govet + gofmt + **depguard 边界** + **gocyclo 复杂度 ≥15**；注：maintidx 在 v2.14 实证 threshold 不生效，故弃用）
- 测试：标准库表驱动测试（`go test ./... -cover`；覆盖率棘轮地板 61，测量口径 = coverprofile 合并总%）

## 环境与命令

```bash
# 同步依赖（零三方依赖，仅缓存标准库）
go mod download

# 运行测试（表驱动 + 覆盖率）
go test ./... -cover

# 启动开发服务器（默认 :8001；ADDR 可覆盖）
go run ./cmd/server

# lint（golangci-lint：govet+gofmt+depguard+gocyclo 一体）
golangci-lint run
```

## 目录结构

```
backend-go/
├── cmd/server/        # 入口层（对齐 backend 的 config/ 角色：只装配，不含业务）
├── internal/app/      # 业务层（对齐 apps/ 角色：handler + 嵌入式 openapi.json + 测试）
├── go.mod
```

## 约定（镜像 backend 的分层边界）

- **依赖方向**：`cmd/server → internal/app` 单向；`internal/*` 禁止反向依赖 `cmd/*`，
  `internal` 内各包相互独立——**depguard 机检已落地**（`.golangci.yml`；反向依赖另被 Go import cycle 语言级拦截）。
- API 形状（路径、JSON 字段、尾斜杠行为）与 `backend/` 保持一致，漂移即 bug。
- `internal/app/openapi.json` 为静态嵌入（`go:embed`）——go:embed 只能引用包内文件，故不入 `api/` 目录。
