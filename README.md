# get-service-start

一个可复用的 **Go API 脚手架**：Gin + Viper 配置 + JWT 鉴权 + GORM/MySQL，内置用户**注册 / 登录 / 当前用户**，适合作为新服务的起点。

仓库地址：<https://github.com/nanami-wq/get-service-start>

## 技术栈

| 类别 | 选型 |
|------|------|
| Web | [Gin](https://github.com/gin-gonic/gin) |
| 配置 | [Viper](https://github.com/spf13/viper)（`config/app.yaml`） |
| 鉴权 | [golang-jwt](https://github.com/golang-jwt/jwt)（HS256） |
| 数据库 | [GORM](https://gorm.io/) + MySQL |
| 密码 | bcrypt |

## 目录结构（简要）

```
.
├── main.go                 # 入口、HTTP Server、优雅退出
├── config/                 # Viper 与 app.yaml 示例
├── core/gin/               # Gin 初始化、路由注册
├── core/store/mysql/       # MySQL 连接、AutoMigrate
├── core/libx/              # 统一响应辅助（Ok / Err、Uid）
├── api/
│   ├── controller/         # HTTP 控制器
│   ├── middleware/         # JWT、CORS、统一 JSON 包装
│   └── router/             # 路由
├── domain/                 # 实体与仓储接口
├── repository/             # GORM 实现
└── usecase/                # 业务用例（注册、登录等）
```

## 环境要求

- Go 1.22+
- MySQL 5.7+ / 8.x（已创建空库）

## 快速开始

1. **克隆**

   ```bash
   git clone https://github.com/nanami-wq/get-service-start.git
   cd get-service-start
   ```

2. **配置**

   ```bash
   cp config/app.yaml.example config/app.yaml
   ```

   编辑 `config/app.yaml`：填写 `mysql`（主机、库名、账号密码）和 `jwt.secret`（请改为足够长的随机字符串，勿提交到公开仓库）。

3. **运行**

   ```bash
   go run .
   ```

   启动时会自动 **AutoMigrate** `users` 表。默认监听端口见配置中的 `app.port`（示例为 `8080`）。

## HTTP 接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/ping` | 健康检查 |
| POST | `/api/auth/register` | 注册，JSON：`username`、`email`、`password`（密码至少 8 位） |
| POST | `/api/auth/login` | 登录，JSON：`account`（用户名或邮箱）、`password`；返回 `access_token` |
| GET | `/api/auth/me` | 当前用户，需 Header：`Authorization: Bearer <token>` |

登录示例：

```bash
curl -s -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"account":"yourname","password":"yourpassword"}'
```

除统一响应包装外，部分路由也会直接返回 JSON（如未带 token 时的 401），以实际响应为准。

## 模块路径与 Fork

当前 Go Module 为：

```text
github.com/nanami-wq/get-service-start
```

若 Fork 到自己的账号或改名仓库，请修改 `go.mod` 第一行，并在全项目将 import 前缀替换为你的模块路径。

## 安全说明

- `config/app.yaml` 已列入 `.gitignore`，请勿将含真实密码、JWT 密钥的配置提交到 Git。
- 生产环境请使用强密钥、HTTPS，并按需加固 CORS、限流与审计。

## 许可

按你的需要自行补充 LICENSE；未特别声明时以仓库内文件为准。
