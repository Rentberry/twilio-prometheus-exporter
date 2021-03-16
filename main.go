package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	count = prometheus.NewDesc("twilio_count", "Count by category", []string{"account", "category"}, nil)
	usage = prometheus.NewDesc("twilio_usage", "Usage by category", []string{"account", "category"}, nil)
	price = prometheus.NewDesc("twilio_price", "Price by category", []string{"account", "category"}, nil)
)

type config struct {
	Addr      string `default:":9153"`
	AccountId string `envconfig:"TWILIO_ACCOUNT_ID" required:"true"`
	Sid       string `envconfig:"TWILIO_SID" required:"true"`
	ApiKey    string `envconfig:"TWILIO_API_KEY" required:"true"`
}

func main() {
	var conf config
	err := envconfig.Process("", &conf)
	if err != nil {
		log.Fatal(err)
	}

	mc := collector{
		accountId: conf.AccountId,
		sid:       conf.Sid,
		apiKey:    conf.ApiKey,
	}

	prometheus.MustRegister(mc)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Twilio statistics exporter</title></head>
             <body>
             <h1>Twilio statistics exporter</h1>
             <p><a href='metrics'>Metrics</a></p>
             </body>
             </html>`))
	})
	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(conf.Addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
