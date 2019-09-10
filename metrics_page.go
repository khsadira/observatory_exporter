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

	if !ok || len(hosts) < 0 {
		log.Println("Url Param 'host' is missing")
		resp = "No host query found"
	} else {
		hosts = reworkURL(hosts)
		for i, target := range hosts {
			metrics = append(metrics, newHTTPMetrics(target))
			metrics[i] = addTLSMetrics(metrics[i], target)
		}
		resp = retAnswer(metrics)
	}
	w.Write([]byte(resp))
}

func retAnswer(metrics []Metrics) string {
	var grade string
	var score string
	var test string
	var tlse string
	var trust string
	var level string
	var tlsgrade string
	var tlsscore string
	var expire string
	var start string

	for _, metric := range metrics {
		grade += fmt.Sprintf("observatory_http_grade{host=\"%s\"} %d", metric.Target, gradeLetterToInt(metric.Grade))
		score += fmt.Sprintf("observatory_http_score{host=\"%s\"} %d", metric.Target, metric.Score)
		test += fmt.Sprintf("observatory_http_tests_passed{host=\"%s\"} %d", metric.Target, metric.TestPass)

		tlse += fmt.Sprintf("observatory_tls_enable{host=\"%s\"} %f", metric.Target, metric.TlsEnable)
		trust += fmt.Sprintf("observatory_tls_is_valid{host=\"%s\"} %f", metric.Target, metric.TlsIsTrust)
		level += fmt.Sprintf("observatory_tls_level{host=\"%s\"} %f", metric.Target, metric.TlsLevel)
		tlsgrade += fmt.Sprintf("observatory_tls_grade{host=\"%s\"} %f", metric.Target, metric.TlsLevel)
		tlsscore += fmt.Sprintf("observatory_tls_score{host=\"%s\"} %f", metric.Target, metric.TlsScore)

		expire += fmt.Sprintf("observatory_cert_expire{host=\"%s\"} %f", metric.Target, metric.CertExpire)
		start += fmt.Sprintf("observatory_cert_start{host=\"%s\"} %f", metric.Target, metric.CertStart)
	}

	resp := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n", helpGrade, typeGrade, grade, helpScore, typeScore, score, helpTest, typeTest, test)
	resp += fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n", helptlse, typetlse, tlse, helptrust, typetrust, trust, helplevel, typelevel, level, helptlsgrade, typetlsgrade, tlsgrade, helptlsscore, typetlsscore, tlsscore, helpexpire, typeexpire, expire, helpstart, typestart, start)
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
