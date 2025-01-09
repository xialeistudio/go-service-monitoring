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

func main() {
	// Initialize metrics collector
	metrics.Init()

	r := gin.Default()

	// Add prometheus metrics middleware
	r.Use(metrics.PrometheusMiddleware())

	// Metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Add alert webhook handler
	r.POST("/webhook", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}

		var webhook AlertWebhook
		if err := json.Unmarshal(body, &webhook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse alert data"})
			return
		}

		// Print alert information to stdout
		fmt.Printf("\n=== Received New Alert ===\n")
		fmt.Printf("Status: %s\n", webhook.Status)
		for _, alert := range webhook.Alerts {
			fmt.Printf("\nAlert Details:\n")
			fmt.Printf("Status: %s\n", alert.Status)
			fmt.Printf("Labels: %v\n", alert.Labels)
			fmt.Printf("Annotations: %v\n", alert.Annotations)
		}
		fmt.Println("==================")

		c.JSON(http.StatusOK, gin.H{"message": "Alert received successfully"})
	})

	// Example APIs
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
