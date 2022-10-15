package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	iso8601Layout string = "2006-01-02T15:04:05.999Z"
)

type AddWorklogBody struct {
	Comment          string `json:"comment"`
	StartTime        string `json:"started"`
	TimeSpentSeconds int64  `json:"timeSpentSeconds"`
}

type StatusCodeError struct {
	CodeActual   int
	CodeExpected int
	Details      string
}

func (s *StatusCodeError) Error() string {
	return fmt.Sprintf("Wrong status code: expected %d, received %d", s.CodeExpected, s.CodeActual)
}

// referenced https://docs.atlassian.com/software/jira/docs/api/REST/7.1.2/#api/2/issue-addWorklog
func addWorkLog(update IssueUpdate, startTs, endTs *time.Time, config *Config) error {
	body, err := buildRequestBody(update.Comment, startTs, endTs)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://%s/rest/api/v2/issue/%s/worklog?adjustEstimate=auto", config.Server, update.IssueKey)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return err
	}
	// add auth header
	rawAuth := fmt.Sprintf("%s:%s", config.Username, config.Password)
	auth := base64.StdEncoding.EncodeToString([]byte(rawAuth))
	req.Header["Authorization"] = []string{fmt.Sprintf("Basic %s", auth)}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return &StatusCodeError{resp.StatusCode, http.StatusCreated, ""}
	}
	//fmt.Println("skipping request", req)
	return nil
}

func buildRequestBody(comment string, startTs, endTs *time.Time) (io.Reader, error) {
	tss := int64(endTs.Sub(*startTs).Seconds())
	stString := startTs.Format(iso8601Layout)
	reqBody := AddWorklogBody{comment, stString, tss}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(body), nil
}
