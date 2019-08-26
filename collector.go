package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func newMetrics(target string) Metrics {
	var metrics Metrics

	json, err := getJSONID(target)
	if err != nil {
		return metrics
	}

	metrics = createMetrics(json)
	metrics.Target = target
	return metrics
}

func createMetrics(js []byte) Metrics {
	var metrics Metrics

	json.Unmarshal(js, &metrics)
	return metrics
}

type scan struct {
	State string `json:"state"`
}

func getJSONID(target string) ([]byte, error) {
	var scan scan
	resp, err := http.Post("https://http-observatory.security.mozilla.org/api/v1/analyze?host="+target+"&rescan=true", "", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &scan)
	if scan.State != "FINISHED" {
		time.Sleep(time.Second / 2)
		return getJSONID(target)
	}
	return buf, err
}
