package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
	DataBaseOpCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "database_operations_total",
			Help: "Total number of database operations",
		},
		[]string{"operation"},
	)
)

func InitPrometheus(r *gin.Engine) {
	prometheus.MustRegister(HttpRequestsTotal)
	prometheus.MustRegister(DataBaseOpCount)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
