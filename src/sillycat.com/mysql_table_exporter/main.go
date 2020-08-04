package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
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
	mysql_table_active prometheus.GaugeVec
	mysql_table_counts prometheus.GaugeVec
}

func NewExporter(metricsPrefix string) *Exporter {
	mysql_table_active := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "mysql_table_active",
		Help:      "This is a gauga vece metric for table status"},
		[]string{"myLabel"})

	mysql_table_counts := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "mysql_table_counts",
		Help:      "This is a gauga vece metric for record count"},
		[]string{"myLabel"})

	return &Exporter{
		mysql_table_active: mysql_table_active,
		mysql_table_counts: mysql_table_counts,
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	// --------
	// logic
	e.mysql_table_active.WithLabelValues("subscriptions").Set(float64(1))
	e.mysql_table_counts.WithLabelValues("subscriptions").Set(float64(13))

	e.mysql_table_active.Collect(ch)
	e.mysql_table_counts.Collect(ch)
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.mysql_table_active.Describe(ch)
	e.mysql_table_counts.Describe(ch)
}

func main() {
	fmt.Println(`
        prometheus exporter mysql_table_exporter,
        metrics expose at http://localhost:18081/mysqltable/metrics
    `)

	// Define parameters
	metricsPath := "/mysqltable/metrics"
	listenAddress := "0.0.0.0:18081"
	metricsPrefix := "mysql_table"

	// Register exporter to Prometheus, call Collect
	exporter := NewExporter(metricsPrefix)
	prometheus.MustRegister(exporter)
	prometheus.MustRegister(version)

	// Launch http service
	http.Handle(metricsPath, promhttp.Handler())
	fmt.Println(http.ListenAndServe(listenAddress, nil))
}
