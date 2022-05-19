package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

var (
	// HttpserverRequestTotal 表示接收http请求总数
	HttpserverRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "httpserver_request_total",
		Help: "The Total number of httpserver requests",
	},
		// 设置标签：请求方法和路径
		[]string{"method", "endpoint"})

	HttpserverRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "httpserver_request_duration_seconds",
		Help:    "httpserver request duration distribution",
		Buckets: []float64{0.1, 0.3, 0.5, 0.7, 0.9, 1},
	},
		[]string{"method", "endpoint"})
)

// 注册监控指标
func init() {
	prometheus.MustRegister(HttpserverRequestTotal)
	prometheus.MustRegister(HttpserverRequestDuration)
}

func NewMetrics(router http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		router(w, r)
		duration := time.Since(start)
		// httpserverRequestTotal 记录
		HttpserverRequestTotal.With(prometheus.Labels{"method": r.Method, "endpoint": r.URL.Path}).Inc()
		// httpserverRequestDuration 记录
		HttpserverRequestDuration.With(prometheus.Labels{"method": r.Method, "endpoint": r.URL.Path}).Observe(duration.Seconds())
	}
}
