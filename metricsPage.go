package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func metricsPage(w http.ResponseWriter, r *http.Request) {
	var resp string
	var metrics []Metrics
	hosts, ok := r.URL.Query()["host"]

	if !ok || len(hosts[0]) < 0 {
		log.Println("Url Param 'host' is missing")
		resp = "No host query found"
	} else {
		targets := strings.Split(hosts[0], ",")
		targets = reworkURL(targets)
		for _, target := range targets {
			metrics = append(metrics, newMetrics(target))
		}
		resp = retAnswer(metrics)
	}
	w.Write([]byte(resp))
}

func retAnswer(metrics []Metrics) string {
	var grade string
	var score string
	var test string

	for _, metric := range metrics {
		grade += fmt.Sprintf("observatory_http_grade{host=%s} %d\n", metric.Target, gradeLetterToInt(metric.Grade))
		score += fmt.Sprintf("observatory_http_score{host=%s} %d\n", metric.Target, metric.Score)
		test += fmt.Sprintf("observatory_http_test{host=%s} %d\n", metric.Target, metric.TestPass)
	}

	resp := fmt.Sprintf("%s\n%s\n%s%s\n%s\n%s%s\n%s\n%s", helpGrade, typeGrade, grade, helpScore, typeScore, score, helpTest, typeTest, test)
	return resp
}

func gradeLetterToInt(str string) int {
	mapping := map[string]int{
		"A+": 12,
		"A":  11,
		"A-": 10,
		"B+": 9,
		"B":  8,
		"B-": 7,
		"C+": 6,
		"C":  5,
		"C-": 4,
		"D+": 3,
		"D":  2,
		"D-": 1,
		"F":  0,
	}

	str = strings.ToUpper(str)
	l, ok := mapping[str]
	if !ok {
		l = 0
	}
	return l
}

func reworkURL(str []string) []string {
	var ret []string
	for _, targetURL := range str {
		targetURL = strings.TrimPrefix(targetURL, "https://")
		targetURL = strings.TrimPrefix(targetURL, "http://")
		targetURL = "https://" + targetURL
		val, _ := url.Parse(targetURL)
		targetURL = val.Hostname()
		_, err := net.Dial("tcp", targetURL+":http")
		if err != nil {
			log.Println("unreachable, error:", err)
		} else {
			ret = append(ret, targetURL)
		}
	}
	return ret
}
