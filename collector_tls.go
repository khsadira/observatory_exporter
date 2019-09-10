package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func addTLSMetrics(metrics Metrics, target string) Metrics {
	tlsID, err := getTLSID(target)
	if err != nil {
		log.Println(err)
	}
	buf, certID, err := getTLSjson(tlsID)
	if err != nil {
		log.Println(err)
	}
	buf2, err := getCERT(certID)
	if err != nil {
		log.Println(err)
		return metrics
	}

	metrics = createTLSMetrics(buf, buf2, metrics)
	return metrics
}

func getTLSID(target string) (int64, error) {
	var scan scanTLS
	query := fmt.Sprintf("https://tls-observatory.services.mozilla.com/api/v1/scan?target=%s", target)
	resp, err := http.Post(query, "", nil)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(http.StatusText(resp.StatusCode))
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &scan)
	if scan.ID == 0 {
		return 0, errors.New("Wrong ScanID return\n")
	}
	return scan.ID, nil
}

func getTLSjson(tlsID int64) ([]byte, int64, error) {
	var scan scanCERT
	query := fmt.Sprintf("https://tls-observatory.services.mozilla.com/api/v1/results?id=%d", tlsID)
	resp, err := http.Get(query)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, 0, errors.New(http.StatusText(resp.StatusCode))
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &scan)
	return buf, scan.ID, nil
}

func getCERT(certID int64) ([]byte, error) {
	query := fmt.Sprintf("https://tls-observatory.services.mozilla.com/api/v1/certificate?id=%d", certID)
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	return buf, nil
}

type scanTLS struct {
	ID int64 `json:"scan_id"`
}

type mozillaEvalData struct {
	Level string `json:"level"`
}

type mozillaGradeData struct {
	Score       float64 `json:"grade"`
	LetterGrade string  `json:"lettergrade"`
}

type scanCERT struct {
	ID int64 `json:"cert_id"`
}

func levelToInt(str string) float64 {
	mapping := map[string]float64{
		"bad":           0,
		"non compliant": 1,
		"old":           2,
		"intermediate":  3,
		"modern":        4,
	}

	str = strings.ToLower(str)
	l, ok := mapping[str]
	if !ok {
		l = -1
	}
	return l
}

func gradeLetterToInt2(str string) float64 {
	mapping := map[string]float64{
		"A": 4,
		"B": 3,
		"C": 2,
		"D": 1,
		"F": 0,
	}

	str = strings.ToUpper(str)
	l, ok := mapping[str]
	if !ok {
		l = 0
	}
	return l
}

func boolToFloat(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func createTLSMetrics(buf []byte, buf2 []byte, metrics Metrics) Metrics {
	var res Scan
	var cert Certificate

	json.Unmarshal(buf, &res)
	json.Unmarshal(buf2, &cert)

	metrics.CertExpire = float64(cert.Validity.NotAfter.Unix())
	metrics.CertStart = float64(cert.Validity.NotBefore.Unix())

	metrics.TlsEnable = boolToFloat(res.Has_tls)
	metrics.TlsIsTrust = boolToFloat(res.Is_valid)

	for _, a := range res.AnalysisResults {
		if a.Success {
			switch a.Analyzer {
			case "mozillaEvaluationWorker":
				var d mozillaEvalData
				err := json.Unmarshal(a.Result, &d)
				if err != nil {
					log.Printf("Failed to unmarshal analyzer 'mozillaEvaluationWorker': %s", err)
					continue
				}
				metrics.TlsLevel = levelToInt(d.Level)

			case "mozillaGradingWorker":
				var d mozillaGradeData
				err := json.Unmarshal(a.Result, &d)
				if err != nil {
					log.Printf("Failed to unmarshal analyzer 'mozillaGradingWorker': %s", err)
					continue
				}
				metrics.TlsScore = float64(d.Score)
				metrics.TlsGrade = gradeLetterToInt2(d.LetterGrade)
			}
		}
	}

	return metrics
}
