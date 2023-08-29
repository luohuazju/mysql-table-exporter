package main

import (
	"fmt"
	"mysql_table_exporter/config"
	"mysql_table_exporter/database"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mysql_table_exporter_version",
		Help: "Version information about this binary",
		ConstLabels: map[string]string{
			"version": "v1.6",
		},
	})
)

type Exporter struct {
	//mysql_table_active prometheus.GaugeVec
	mysql_table_counts prometheus.GaugeVec
}

func NewExporter(metricsPrefix string) *Exporter {
	// mysql_table_active := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
	// 	Namespace: metricsPrefix,
	// 	Name:      "mysql_table_active",
	// 	Help:      "This is a gauga vece metric for table status"},
	// 	[]string{"myLabel"})

	mysql_table_counts := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "mysql_table_counts",
		Help:      "This is a gauga vece metric for record count"},
		[]string{"myLabel"})

	return &Exporter{
		//mysql_table_active: mysql_table_active,
		mysql_table_counts: mysql_table_counts,
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	// --------
	// logic
	//subscriptions_flag := database.GetTableStatus("subscriptions", 10)
	weekly_ads_count := database.GetTableCreatedCount("weekly_ads", 1440)

	//e.mysql_table_active.WithLabelValues("subscriptions").Set(float64(subscriptions_flag))
	e.mysql_table_counts.WithLabelValues("weekly_ads").Set(float64(weekly_ads_count))

	//e.mysql_table_active.Collect(ch)
	e.mysql_table_counts.Collect(ch)
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	//e.mysql_table_active.Describe(ch)
	e.mysql_table_counts.Describe(ch)
}

func main() {
	// Define parameters
	httpPort := config.GetEnv("HTTP_PORT", "18081")
	httpHost := config.GetEnv("HTTP_HOST", "localhost")
	metricsPath := config.GetEnv("METRICS_PATH", "/mysqltable/metrics")
	metricsPrefix := config.GetEnv("METRICS_PREFIX", "mysql_table")
	listenAddress := httpHost + ":" + httpPort

	fmt.Printf(`
        prometheus exporter mysql_table_exporter,
        metrics expose at http://%s:%s%s
	`, httpHost, httpPort, metricsPath)
	fmt.Println("\n")

	exporter := NewExporter(metricsPrefix)
	registry := prometheus.NewRegistry()
	registry.MustRegister(exporter)
	registry.MustRegister(version)
	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	// [...] update metrics within a goroutine

	http.Handle(metricsPath, handler)
	http.ListenAndServe(listenAddress, nil)
	// exporter := NewExporter(metricsPrefix)
	// prometheus.MustRegister(exporter)
	// prometheus.MustRegister(version)

	// http.Handle(metricsPath, promhttp.Handler())
	// http.ListenAndServe(listenAddress, nil)

}
