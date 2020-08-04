package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

var (
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mysql_table_exporter_version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": "v1.0",
		},
	})
)

type Exporter struct {
	gauge    prometheus.Gauge
	gaugeVec prometheus.GaugeVec
}

func NewExporter(metricsPrefix string) *Exporter {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "gauge_metric",
		Help:      "This is a gauge metric"})

	gaugeVec := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "gauge_vec_metric",
		Help:      "This is a gauga vece metric"},
		[]string{"myLabel"})

	return &Exporter{
		gauge:    gauge,
		gaugeVec: gaugeVec,
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	// --------
	// 业务逻辑
	timestamp := time.Now().Unix()
	fmt.Println("Timestamp: ", timestamp)
	rand.Seed(timestamp)
	ranint := rand.Intn(10000)
	fmt.Println("Random: ", ranint)
	e.gauge.Set(float64(timestamp))
	e.gaugeVec.WithLabelValues("helloworld").Set(float64(ranint))
	// --------
	// Called use a concurrency safe way
	e.gauge.Collect(ch)
	e.gaugeVec.Collect(ch)
}

// metric 描述, 可被重写
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.gauge.Describe(ch)
	e.gaugeVec.Describe(ch)
}

func main() {
	fmt.Println(`
        prometheus exporter example,
        metrics expose at http://:18081/metrics
    `)

	// Define parameters
	metricsPath := "/metrics"
	listenAddress := "0.0.0.0:18081"
	metricsPrefix := "fake"

	// Register exporter to Prometheus, call Collect
	exporter := NewExporter(metricsPrefix)
	prometheus.MustRegister(exporter)
	prometheus.MustRegister(version)

	// Launch http service
	http.Handle(metricsPath, promhttp.Handler())
	fmt.Println(http.ListenAndServe(listenAddress, nil))
}
