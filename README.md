# Go Service Monitoring Example

[中文文档](README-CN.md)

A Go microservice example demonstrating how to build a complete monitoring and alerting system using Prometheus, Alertmanager, and Grafana.

## Features

- RESTful API service based on Gin framework
- Prometheus metrics collection and monitoring
- Alertmanager for alert management
- Grafana for data visualization
- Complete Docker Compose deployment configuration

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.21 or higher (for local development)

### Launch Services

```bash
# Clone the repository
git clone https://github.com/yourusername/go-service-monitoring.git
cd go-service-monitoring

# Start all services
docker-compose up -d
```

### Access Services

- Go Application: http://localhost:8080
  - `/hello` - Example API endpoint
  - `/metrics` - Prometheus metrics endpoint
  - `/webhook` - Alert webhook endpoint
- Prometheus: http://localhost:9090
- Alertmanager: http://localhost:9093
- Grafana: http://localhost:3000
  - Username: admin
  - Password: admin

## Project Structure

```
.
├── cmd/
│   └── server/          # Main service entry
├── internal/
│   └── metrics/         # Metrics collector
├── prometheus/
│   ├── rules/           # Alert rules
│   ├── alertmanager.yml # Alertmanager configuration
│   └── prometheus.yml   # Prometheus configuration
├── docker-compose.yml   # Container orchestration
└── Dockerfile          # Go service container build
```

## Monitoring Metrics

- HTTP request latency
- Request count
- Error rate
- Goroutine count

## Alert Rules

- Goroutine count exceeds threshold
- High HTTP request error rate
- Service response time anomaly

## License

MIT License

## Contributing

Issues and Pull Requests are welcome! 
