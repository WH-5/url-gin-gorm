# 🌍 短链接生成器

## 🚀 项目简介

本项目是对之前短链接生成器的重写，基于 **Gin + GORM** 框架，并采用 **Kratos 推荐的架构** 进行组织。

- **更新预告**：将会使用 **KeyDB** 代替 Redis，提升缓存性能。

## 📌 实现功能

- **创建短链接**：`POST /api/url` 生成短链接。
- **短链接跳转**：访问 `/:code` 自动重定向到原始 URL。
- **自动删除过期链接**：定期清理数据库中过期的短链接。
- **缓存优化**：短时间内的访问将直接从 缓存 读取，提升响应速度。
- **记录ip访问次数**：由于很多疑似黑客的ip访问网页，所以记录一下，想一个应对方法

## 🛠️ 技术栈

- **Gin** - 轻量级 Web 框架
- **GORM** - ORM 数据库操作
- **Redis** - 高性能 NoSQL 存储（或KeyDB）

## 📡 API 设计

### 1️⃣ 创建短链接

**`POST /api/url`**

#### 请求示例：

```json
{
  "original_url": "https://www.youtube.com",
  "custom_code": "2345",
  "duration": 1
}
```
| 参数名         | 类型   | 是否必填 | 说明 |
|--------------|------|------|--------------------------------|
| `original_url` | `string` | ✅ 必填 | 原始长链接，不能为空 |
| `custom_code` | `string` | ❌ 选填 | 自定义短链接代码，如果不填则自动生成 |
| `duration`    | `int`    | ❌ 选填 | 短链接的有效期，单位：小时，若不指定则使用默认值 |

# 🚀 运行教程

## **1️⃣ 安装依赖**
本项目依赖 **Go 语言环境** 和 **Docker**，请确保已安装：

- [Go 官方安装指南](https://go.dev/doc/install)
- [Docker 官方安装指南](https://docs.docker.com/get-docker/)


## **2️⃣ 下载项目**
```sh
git clone https://github.com/WH-5/url-gin-gorm.git
cd url-gin-gorm
```
### **3️⃣ 初始化数据库和缓存**
**首次运行**，执行以下命令创建数据库和缓存容器：
```sh
make make_db
```
如果如果数据库和缓存容器已存在，后续启动时只需执行：
```sh
make start_db
```
## **4️⃣ 启动项目**
```sh
go run main.go
```
