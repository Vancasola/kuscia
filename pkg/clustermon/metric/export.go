// Package metric the function to export metrics to Prometheus
package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"strings"
)

func ProduceCounter(namespace string, name string, help string, labels map[string]string) prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace:   namespace,
			Name:        name,
			Help:        help,
			ConstLabels: labels,
		})
}

func ProduceGauge(namespace string, name string, help string, labels map[string]string) prometheus.Gauge {
	return prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace:   namespace,
			Name:        name,
			Help:        help,
			ConstLabels: labels,
		})
}

func ProduceHistogram(namespace string, name string, help string, labels map[string]string) prometheus.Histogram {
	return prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace:   namespace,
			Name:        name,
			Help:        help,
			ConstLabels: labels,
		})
}
func ProduceSummary(namespace string, name string, help string, labels map[string]string) prometheus.Summary {
	return prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace:   namespace,
			Name:        name,
			Help:        help,
			ConstLabels: labels,
		})
}

var counters = make(map[string]prometheus.Counter)
var gauges = make(map[string]prometheus.Gauge)
var histograms = make(map[string]prometheus.Histogram)
var summaries = make(map[string]prometheus.Summary)

func ProduceMetric(reg *prometheus.Registry,
	metricType string, nameSpace string, name string, help string, labels map[string]string) *prometheus.Registry {
	var metricId string
	if labels["type"] == "envoy" {
		metricId = "cluster__" + labels["cluster_name"] + "__" + name
	} else if labels["type"] == "ss" {
		metricId = labels["local_domain"] + "__" + labels["remote_domain"] + "__" + name + "__" + labels["aggregation_function"]
	}
	if metricType == "Counter" {
		counters[metricId] = ProduceCounter(nameSpace, name, help, labels)
		reg.MustRegister(counters[metricId])
	} else if metricType == "Gauge" {
		gauges[metricId] = ProduceGauge(nameSpace, name, help, labels)
		reg.MustRegister(gauges[metricId])
	} else if metricType == "Histogram" {
		histograms[metricId] = ProduceHistogram(nameSpace, name, help, labels)
		reg.MustRegister(histograms[metricId])
	} else if metricType == "Summary" {
		summaries[metricId] = ProduceSummary(nameSpace, name, help, labels)
		reg.MustRegister(summaries[metricId])
	}
	return reg
}
func Formalize(metric string) string {
	metric = strings.Replace(metric, "-", "_", -1)
	metric = strings.Replace(metric, ".", "__", -1)
	metric = strings.ToLower(metric)
	return metric
}
func ProduceMetrics(localDomainName string,
	clusterAddresses map[string][]string,
	netMetrics []string,
	cluMetrics []string,
	MetricTypes map[string]string,
	aggregationMetrics map[string]string) *prometheus.Registry {
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
	for clusterName := range clusterAddresses {
		destinationAddress := clusterAddresses[clusterName]
		clusterName = Formalize(clusterName)
		labels := make(map[string]string)
		labels["type"] = "ss"
		labels["local_domain"] = localDomainName
		for _, dstAddr := range destinationAddress {
			labels["remote_domain"] = Formalize(strings.Split(dstAddr, ":")[0])
			for _, metric := range netMetrics {
				metric = Formalize(metric)
				labels["aggregation_function"] = aggregationMetrics[metric]
				reg = ProduceMetric(reg, MetricTypes[metric], "ss", metric, metric+"with aggregation function "+aggregationMetrics[metric], labels)
			}
		}

		labels = make(map[string]string)
		labels["type"] = "envoy"
		for _, metric := range cluMetrics {
			metric = Formalize(metric)
			labels["cluster_name"] = clusterName
			reg = ProduceMetric(reg, MetricTypes[metric], "envoy", metric, metric, labels)
		}
	}
	return reg
}

func UpdateMetrics(clusterResults map[string]float64, MetricTypes map[string]string) {
		for metric, val := range clusterResults {
			metricId := Formalize(metric)
			metricTypeId := strings.Join(strings.Split(metric, ".")[2:3], "__")
			switch MetricTypes[metricTypeId] {
			case "Counter":
				counters[metricId].Add(val)
			case "Gauge":
				gauges[metricId].Set(val)
			case "Histogram":
				histograms[metricId].Observe(val)
			case "Summary":
				summaries[metricId].Observe(val)
			}
		}
}