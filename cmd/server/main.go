package main

import (
	"encoding/json"
	"fmt"
	"io"
	"monitoring-demo/internal/metrics"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Alert struct {
	Status      string            `json:"status"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

func main() {
	// 初始化metrics收集器
	metrics.Init()

	r := gin.Default()

	// 添加prometheus metrics中间件
	r.Use(metrics.PrometheusMiddleware())

	// metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	type AlertWebhook struct {
		Version           string            `json:"version"`
		GroupKey          string            `json:"groupKey"`
		Status            string            `json:"status"`
		Receiver          string            `json:"receiver"`
		GroupLabels       map[string]string `json:"groupLabels"`
		CommonLabels      map[string]string `json:"commonLabels"`
		CommonAnnotations map[string]string `json:"commonAnnotations"`
		ExternalURL       string            `json:"externalURL"`
		Alerts            []Alert           `json:"alerts"`
	}

	// 添加告警webhook处理器
	r.POST("/webhook", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "读取请求体失败"})
			return
		}

		var webhook AlertWebhook
		if err := json.Unmarshal(body, &webhook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "解析告警数据失败"})
			return
		}

		// 打印告警信息到stdout
		fmt.Printf("\n=== 收到新的告警 ===\n")
		fmt.Printf("状态: %s\n", webhook.Status)
		for _, alert := range webhook.Alerts {
			fmt.Printf("\n告警详情:\n")
			fmt.Printf("状态: %s\n", alert.Status)
			fmt.Printf("标签: %v\n", alert.Labels)
			fmt.Printf("注释: %v\n", alert.Annotations)
		}
		fmt.Println("==================")

		c.JSON(http.StatusOK, gin.H{"message": "告警接收成功"})
	})

	// 示例API
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/error", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	})

	r.Run(":8080")
}
