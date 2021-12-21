package main

import (
	"common"
	"common/module"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()
	engine.Use(metricMiddle())
	engine.GET("/metrics", prometheusHandler())
	engine.POST(common.PathDataJson, func(c *gin.Context) {
		var data module.JsonDataReq
		err := c.ShouldBindJSON(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.ServiceJsonData(&data))
	})
	engine.Run()
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func metricMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}
		start := time.Now()
		c.Next()
		common.PromSummaryRequestLatency.Observe(float64(time.Since(start)))
	}
}
