# todo-backend-go

todo backend 的 **Go 栈适配镜像**（dev-env-settings 仓库 `backend-go/` 子模块）——与 `devcontainer-backend`
保持相同 API 面与角色分工（骨架阶段：health + OpenAPI schema），作为 go 栈工具链接入的 dogfooding 载体
（stack.json 声明 / CI golang 分支 / 守卫注册表的真实首例）。

- 镜像关系与边界约定见模块 `AGENTS.md`
- 存储为内存占位（db 字段恒 ok）；todo CRUD 随主 backend 同步演进
