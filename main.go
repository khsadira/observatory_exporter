package main

import (
	"log"
	"net/http"
)

type Metrics struct {
	Target   string `json:"target"`
	Grade    string `json:"grade"`
	TestPass int    `json:"tests_passed"`
	Score    int    `json:"score"`
}

const (
	helpGrade = "# HELP observatory_http_grade Grade representation of score, A+=12, A=11, A-=10, B+=9, B=8, B-=7, C+=6, C=5, C-=4, D+=3, D=2, D-=1, F=0"
	typeGrade = "# TYPE observatory_http_grade gauge"
	helpScore = "# HELP observatory_http_score Http score"
	typeScore = "# TYPE observatory_http_score gauge"
	helpTest  = "# HELP observatory_http_tests_passed Number of test passed"
	typeTest  = "# TYPE observatory_http_tests_passed gauge"
)

func main() {
	http.HandleFunc("/metrics/", metricsPage)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
<head><title>Observatory Exporter</title></head>
			 <body>
			 <h1>Observatory Exporter</h1>
			 <p><a href='/metrics/'>observatory-http metrics</a></p>
			 </body>
			 </html>`))
	})
	log.Fatal(http.ListenAndServe(":9230", nil))
}
