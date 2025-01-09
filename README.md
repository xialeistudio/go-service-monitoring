# Go 服务监控示例

这是一个展示如何使用 Prometheus、Alertmanager 和 Grafana 构建完整监控告警体系的 Go 微服务示例项目。

## 功能特性

- 基于 Gin 框架的 RESTful API 服务
- Prometheus 指标采集和监控
- Alertmanager 告警管理
- Grafana 数据可视化
- 完整的 Docker Compose 部署配置

## 快速开始

### 前置要求

- Docker 和 Docker Compose
- Go 1.21 或更高版本（本地开发时需要）

### 启动服务

```bash
# 克隆项目
git clone https://github.com/yourusername/go-service-monitoring.git
cd go-service-monitoring

# 启动所有服务
docker-compose up -d
```

### 访问服务

- Go 应用: http://localhost:8080
  - `/hello` - 示例 API 端点
  - `/metrics` - Prometheus 指标端点
  - `/webhook` - 告警 Webhook 端点
- Prometheus: http://localhost:9090
- Alertmanager: http://localhost:9093
- Grafana: http://localhost:3000
  - 用户名: admin
  - 密码: admin

## 项目结构

```
.
├── cmd/
│   └── server/          # 主服务入口
├── internal/
│   └── metrics/         # 指标收集器
├── prometheus/
│   ├── rules/           # 告警规则
│   ├── alertmanager.yml # Alertmanager 配置
│   └── prometheus.yml   # Prometheus 配置
├── docker-compose.yml   # 容器编排配置
└── Dockerfile          # Go 服务容器构建文件
```

## 监控指标

- HTTP 请求延迟
- 请求计数
- 错误率
- Goroutine 数量

## 告警规则

- Goroutine 数量超过阈值
- HTTP 请求错误率过高
- 服务响应时间异常

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！ 