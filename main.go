package main

import (
	"customresourceexporter/pkg/opa"
	"flag"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port      = flag.String("port", ":9191", "Address to listen on metrics")
	metricsEndpoint   = flag.String("metrics", "/metrics", "Path Where Metric is exposed")
	inCluster = flag.Bool("incluster", false, "Cluster is running time")
	ticker    *time.Ticker
	done      = make(chan bool)
	metrics = []prometheus.metrics
)

type Exporter struct {
	
}

func newExporter() *Exporter  {
	return  &Exporter{}
}
func (e *Exporter)Describe(ch chan <- *prometheus.Desc)  {
	ch <- opa.Up
	ch <- opa.ConstraintViolation
	ch <- opa.ConstraintInformation
}