# 元耕睿枢（MetaCult_Insight）
 The first project from 2023&2024 Geekers.

### 一、项目目录

```
MetaCult_Insight/
├──src
|   └── cmd/                 # 主程序入口
|   │   └── main.go          # 启动入口（初始化配置、路由等）
|   ├── configs/             # 配置目录
|   │   └──config.yaml       # 配置文件
|   ├── internal/            # 业务逻辑
|   │   ├── controller/      # 控制器层（处理 HTTP 请求）
|   │   │   └── user.go      # 用户相关接口
|   │   ├── service/         # 核心业务逻辑
|   │   │   └── user.go      # 用户业务逻辑
|   │   ├── model/           # 模型层（数据结构定义）
|   │   │   └── user.go      # 用户模型
|   │   ├── middleware/      # 中间件（认证、日志等）
|   │   │   └── auth.go      # 认证中间件
|   │   └── router/          # 路由定义
|   │       └── router.go    # 路由配置
|   ├── dao/                 # 数据访问实现
|   │   └── mysql/           # 使用gorm 
├── go.mod               # Go 模块文件
├── go.sum
└── README.md
```
 
