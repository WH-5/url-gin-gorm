# 🌍 短链接生成器

## 🚀 项目简介

本项目是对之前短链接生成器的重写，基于 **Gin + GORM** 框架，并采用 **Kratos 推荐的架构** 进行组织。

- **更新预告**：将会使用 **KeyDB** 代替 Redis，提升缓存性能。

## 📌 实现功能

- **创建短链接**：`POST /api/url` 生成短链接。
- **短链接跳转**：访问 `/:code` 自动重定向到原始 URL。
- **自动删除过期链接**：定期清理数据库中过期的短链接。
- **缓存优化**：短时间内的访问将直接从 缓存 读取，提升响应速度。

## 🛠️ 技术栈

- **Gin** - 轻量级 Web 框架
- **GORM** - ORM 数据库操作
- **Reids** - 高性能 NoSQL 存储（或KeyDB）

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

