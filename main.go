package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"os/exec"
	"strconv"
	// "math/rand"
	"net/http"
	"time"
	"strings"
)

type Collector struct {
	NumberOfProcessesMetric *prometheus.Desc
	NumberOfNetworkInterfaces *prometheus.Desc
}

func newCollector() *Collector {
	return &Collector{
		NumberOfProcessesMetric: prometheus.NewDesc("NumberOfProcessesMetric",
			"Shows number of processes",
			nil, nil,
		),
		NumberOfNetworkInterfaces: prometheus.NewDesc("NumberOfNetworkInterfaces",
			"Shows number of network interfaces",
			nil, nil,
		),
	}
}

func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.NumberOfProcessesMetric
	ch <- collector.NumberOfNetworkInterfaces
}

func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	m1 := prometheus.MustNewConstMetric(collector.NumberOfProcessesMetric, prometheus.GaugeValue, getNumberOfProcesses())
	m2 := prometheus.MustNewConstMetric(collector.NumberOfNetworkInterfaces, prometheus.GaugeValue, getNumberOfNetworkInterfaces())
	m1 = prometheus.NewMetricWithTimestamp(time.Now().Add(-time.Hour), m1)
	m2 = prometheus.NewMetricWithTimestamp(time.Now(), m2)
	ch <- m1
	ch <- m2
}

func getNumberOfProcesses() float64 {
	cmd := "ps -e | wc -l"
	commandOutput, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	converted, err := strconv.ParseFloat(strings.TrimSuffix(string(commandOutput), "\n"), 64)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return converted-3
}

func getNumberOfNetworkInterfaces() float64 {
	cmd := "ls /sys/class/net | wc -l"
	commandOutput, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	converted, err := strconv.ParseFloat(strings.TrimSuffix(string(commandOutput), "\n"), 64)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return converted
}

func main() {
	collector := newCollector()
	prometheus.MustRegister(collector)

	http.Handle("/console/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9101", nil))
}