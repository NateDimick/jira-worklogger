package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	iso8601Layout string = "2006-01-02T15:04:05.999Z"
	jiraLayout    string = "2006-01-02T15:04:05.000-0700"
)

type AddWorklogBody struct {
	Comment   string `json:"comment"`
	StartTime string `json:"started"`
	TimeSpent string `json:"timeSpent"`
}

type StatusCodeError struct {
	CodeActual   int
	CodeExpected int
	Details      string
}

func (s *StatusCodeError) Error() string {
	return fmt.Sprintf("Wrong status code: expected %d, received %d", s.CodeExpected, s.CodeActual)
}

// referenced https://docs.atlassian.com/software/jira/docs/api/REST/9.3.0/#api/2/issue-addWorklog
func addWorkLog(update IssueUpdate, startTs, endTs *time.Time, config *Config) error {
	body, err := buildRequestBody(update.Comment, startTs, endTs)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://%s/rest/api/2/issue/%s/worklog", config.Server, update.IssueKey)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}
	// query string
	qs := req.URL.Query()
	qs.Add("adjustEstimate", "auto")
	req.URL.RawQuery = qs.Encode()
	// add auth header
	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", config.Token)}
	req.Header["Content-Type"] = []string{"application/json"}
	fmt.Printf("request: %+v\n", req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return &StatusCodeError{resp.StatusCode, http.StatusCreated, ""}
	}
	return nil
}

func buildRequestBody(comment string, startTs, endTs *time.Time) (io.Reader, error) {
	tss := int(endTs.Sub(*startTs).Seconds())
	stString := startTs.Format(jiraLayout)
	reqBody := AddWorklogBody{comment, stString, buildTimeSpent(tss)}
	fmt.Printf("request body: %+v\n", reqBody)
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(body), nil
}

func buildTimeSpent(seconds int) string {
	minutes := seconds / 60
	mm := minutes % 60
	hh := minutes / 60
	var timeSpent string
	if hh != 0 {
		hours := float64(minutes) / 60.0
		timeSpent = fmt.Sprintf("%.1fh", hours)
	} else if mm != 0 {
		timeSpent = fmt.Sprintf("%dm", minutes)
	} else if hh == 0 && mm == 0 {
		timeSpent = "1m"
	}
	return timeSpent
}
