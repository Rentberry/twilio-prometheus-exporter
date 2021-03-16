package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
)

type collector struct {
	sid       string
	apiKey    string
	accountId string
}

func (c collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- count
	ch <- usage
	ch <- price
}

func (c collector) Collect(ch chan<- prometheus.Metric) {
	info, err := c.getUsageRecords()
	if err != nil {
		log.Print(err)
		return
	}

	for _, v := range info.UsageRecords {
		var err error
		countValue, err := strconv.ParseFloat(v.Count, 64)
		if err != nil {
			log.Println(err)
		}

		usageValue, err := strconv.ParseFloat(v.Usage, 64)
		if err != nil {
			log.Println(err)
		}

		priceValue, err := strconv.ParseFloat(v.Price, 64)
		if err != nil {
			log.Println(err)
		}

		if countValue == 0 && usageValue == 0 && priceValue == 0 {
			continue
		}

		ch <- prometheus.MustNewConstMetric(count, prometheus.CounterValue, countValue, c.accountId, v.Category)
		ch <- prometheus.MustNewConstMetric(usage, prometheus.CounterValue, usageValue, c.accountId, v.Category)
		ch <- prometheus.MustNewConstMetric(price, prometheus.CounterValue, priceValue, c.accountId, v.Category)
	}
}
