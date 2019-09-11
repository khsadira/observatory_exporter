package main

import (
	"log"
	"net/http"
)

type Metrics struct {
	Target     string `json:"target"`
	Grade      string `json:"grade"`
	TestPass   int    `json:"tests_passed"`
	Score      int    `json:"score"`
	TlsEnable  float64
	TlsIsTrust float64
	TlsLevel   float64
	TlsGrade   float64
	TlsScore   float64
	CertExpire float64
	CertStart  float64
}

func main() {
	http.HandleFunc("/metrics/", metricsPage)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
<head><title>Observatory Exporter</title></head>
			 <body>
			 <h1>Observatory Exporter</h1>
			 <p><a href='/metrics/'>observatory metrics</a></p>
			 </body>
			 </html>`))
	})
	log.Fatal(http.ListenAndServe(":9230", nil))
}
