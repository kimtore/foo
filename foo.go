package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var g float64

var (
	requests = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "foo",
		Subsystem: "main",
		Name:      "root_requests",
		Help:      "Number of requests made.",
	})
	checks = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "foo",
		Subsystem: "main",
		Name:      "health_requests",
		Help:      "Number of health checks made.",
	})
	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "foo",
		Subsystem: "main",
		Name:      "gauge",
		Help:      "Arbitrary number",
	})
)

func init() {
	prometheus.MustRegister(requests)
	prometheus.MustRegister(checks)
	prometheus.MustRegister(gauge)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, foo, how are you?")
	requests.Inc()
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IMOK")
	checks.Inc()
}

func gaugeHandler(w http.ResponseWriter, r *http.Request) {
	if g == 0 {
		g = 100
	} else {
		g = 0
	}
	gauge.Set(g)
	fmt.Fprintf(w, "%.2f", g)
}

func main() {
	h := http.NewServeMux()
	h.Handle("/metrics", promhttp.Handler())
	h.HandleFunc("/", handler)
	h.HandleFunc("/health", health)
	h.HandleFunc("/gauge", gaugeHandler)
	log.Fatal(http.ListenAndServe(":8080", h))
}
