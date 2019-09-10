package main

import (
	"encoding/json"
	"time"
)

type Scan struct {
	ID               int64                  `json:"id"`
	Timestamp        time.Time              `json:"timestamp"`
	Target           string                 `json:"target"`
	Replay           int                    `json:"replay"` //hours or days
	Has_tls          bool                   `json:"has_tls"`
	Cert_id          int64                  `json:"cert_id"`
	Trust_id         int64                  `json:"trust_id"`
	Is_valid         bool                   `json:"is_valid"`
	Validation_error string                 `json:"validation_error,omitempty"`
	ScanError        string                 `json:"scan_error,omitempty"`
	Complperc        int                    `json:"completion_perc"`
	AnalysisResults  Analyses               `json:"analysis,omitempty"`
	Ack              bool                   `json:"ack"`
	Attempts         int                    `json:"attempts"` //number of retries
	AnalysisParams   map[string]interface{} `json:"analysis_params"`
}

type Analysis struct {
	ID       int64           `json:"id"`
	Analyzer string          `json:"analyzer"`
	Result   json.RawMessage `json:"result"`
	Success  bool            `json:"success"`
}

type Analyses []Analysis

type Certificate struct {
	Validity Validity `json:"validity"`
}

type Validity struct {
	NotBefore time.Time `json:"notBefore"`
	NotAfter  time.Time `json:"notAfter"`
}
