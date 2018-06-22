package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
)

func init() {
	prometheus.MustRegister(requests)
	prometheus.MustRegister(checks)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, foo, how are you?")
	requests.Inc()
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IMOK")
	checks.Inc()
}

func main() {
	h := http.NewServeMux()
	h.Handle("/metrics", promhttp.Handler())
	h.HandleFunc("/", handler)
	h.HandleFunc("/health", health)
	log.Fatal(http.ListenAndServe(":8080", h))
}
