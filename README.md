# Go Blog

基于 Go + Gin 的博客后端服务。

## 技术栈

- **框架**: Gin v1.10
- **数据库**: MySQL (GORM v1.25)
- **缓存**: Redis
- **搜索**: Elasticsearch 7.x
- **定时任务**: robfig/cron/v3
- **日志**: Logrus
- **其他**: JWT 认证、七牛云存储、QQ登录、WebSocket 聊天、AI 接口

## 项目结构

```
.
├── api/                # HTTP 接口层（按模块划分）
│   ├── article_api/    # 文章管理
│   ├── banner_api/     # 轮播图管理
│   ├── captcha_api/    # 验证码
│   ├── chat_api/       # WebSocket 聊天
│   ├── comment_api/    # 评论管理
│   ├── data_api/       # 数据统计
│   ├── focus_api/      # 用户关注
│   ├── global_notification_api/  # 全站通知
│   ├── image_api/      # 图片管理
│   ├── log_api/        # 日志管理
│   ├── search_api/     # ES 搜索
│   ├── site_api/       # 站点配置
│   ├── site_msg_api/   # 站内消息
│   └── user_api/       # 用户管理
├── common/             # 公共常量、枚举
├── conf/               # 配置结构定义
├── core/               # 核心初始化（DB/Redis/ES/日志等）
├── flags/              # 命令行参数与工具
├── global/             # 全局变量
├── middleware/         # Gin 中间件
├── models/             # 数据模型（GORM）
├── router/             # 路由注册
├── service/            # 业务逻辑层
├── testdata/           # 测试/示例代码
├── uploads/            # 静态文件上传目录
├── main.go             # 程序入口
├── dockerfile          # Docker 构建文件
└── settings.yaml       # 配置文件（需自行创建）
```

## 快速开始

### 前置依赖

- Go 1.24+
- MySQL
- Redis
- Elasticsearch 7.x（可选，用于搜索功能）

### 本地运行

```bash
# 1. 创建配置文件 settings.yaml（参考下方配置说明）

# 2. 安装依赖
go mod download

# 3. 运行
go run main.go -f settings.yaml
```

### 命令行参数

| 参数 | 说明 |
|------|------|
| `-f` | 指定配置文件路径，默认 `settings.yaml` |
| `-db` | 执行数据库迁移后退出 |
| `-es` | 创建 ES 索引后退出 |
| `-t user -s create` | 创建用户 |
| `-v` | 显示版本 |

### Docker 运行

```bash
docker build -t goblog .
docker run -d -p 8080:8080 -v $(pwd)/settings.yaml:/app/settings.yaml goblog
```

## 配置文件说明

创建 `settings.yaml`，包含以下配置项：

```yaml
system:
  host: "0.0.0.0"
  port: 8080
  gin_mode: "release"   # debug / release

db:
  - user: "root"
    password: "password"
    host: "127.0.0.1"
    port: 3306
    dbname: "goblog"

redis:
  addr: "127.0.0.1:6379"
  password: ""
  db: 0

jwt:
  secret: "your-secret-key"
  expires: 24

email:
  # 邮件服务配置

qiniu:
  # 七牛云存储配置

qq:
  # QQ 登录配置

ai:
  # AI 接口配置

es:
  # Elasticsearch 配置
```

## API 模块概览

所有接口前缀为 `/api`：

| 模块 | 路由前缀 | 功能 |
|------|---------|------|
| 站点 | `/api/site` | 站点信息、配置管理 |
| 用户 | `/api/user` | 注册、登录、JWT认证、用户管理 |
| 文章 | `/api/article` | 文章 CRUD、点赞、收藏 |
| 评论 | `/api/comment` | 评论发布、列表、删除 |
| 轮播图 | `/api/banner` | Banner 管理 |
| 图片 | `/api/image` | 图片上传、管理 |
| 搜索 | `/api/search` | ES 全文搜索 |
| 聊天 | `/api/chat` | WebSocket 实时聊天 |
| 关注 | `/api/focus` | 用户关注/粉丝 |
| 消息 | `/api/message` | 站内消息、全站通知 |
| 验证码 | `/api/captcha` | 图形验证码 |
| 日志 | `/api/log` | 系统日志管理 |
| AI | `/api/ai` | AI 相关接口 |
| 数据 | `/api/data` | 数据统计 |

## 核心功能特性

- **文章系统**: 支持 Markdown、点赞、收藏、评论
- **用户系统**: JWT 认证、QQ 第三方登录、用户关注
- **实时聊天**: WebSocket 实现的即时通讯
- **搜索**: 基于 Elasticsearch 的全文检索
- **文件存储**: 支持本地存储 + 七牛云 CDN
- **数据同步**: 基于 go-mysql-elasticsearch 的 MySQL 到 ES 同步
- **定时任务**: 内置 cron 定时任务服务
- **验证码**: base64Captcha 图形验证码
- **IP 归属地**: ip2region 离线 IP 定位
- **日志**: 结构化日志 + 文件切割
