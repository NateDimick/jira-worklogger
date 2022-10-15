package main

import (
	"fmt"
	"strconv"
	"time"
)

type IssueUpdate struct {
	IssueKey string `json:"IssueKey"`
	Comment  string `json:"Comment"`
}

type Config struct {
	Server   string `json:"Server"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type CurrentIssue struct {
	IssueKey  string `json:"IssueKey"`
	Started   bool   `json:"Started"`
	StartTime string `json:"StartTime"`
}

type StartStopResult struct {
	Success bool   `json:"Success"`
	Time    string `json:"Time"`
}

func (a *App) UpdateConfig(update Config) bool {
	err := setFile(configFile, update)
	if err != nil {
		a.emitError(err)
		return false
	}
	return true
}

func (a *App) StartStopWorkTime(update IssueUpdate) StartStopResult {
	eventTime := time.Now().UTC()
	err := startStop(update, &eventTime)
	if err != nil {
		a.emitError(err)
		return *new(StartStopResult)
	}
	return StartStopResult{true, eventTime.Format(iso8601Layout)}
}

func (a *App) AbortCurrentIssue() bool {
	err := setFile(stateFile, new(CurrentIssue))
	if err != nil {
		a.emitError(err)
		return false
	}
	return true
}

func (a *App) ManualStopWork(update IssueUpdate, localTs string) bool {
	dtLocalLayout := "2006-01-02T15:04"
	t, err := time.Parse(dtLocalLayout, localTs)
	if err != nil {
		a.emitError(err)
		return false
	}
	t = t.UTC()
	err = startStop(update, &t)
	if err != nil {
		a.emitError(err)
		return false
	}
	return true
}

func (a *App) GetCurrentIssue() CurrentIssue {
	issue, err := getFile[CurrentIssue](stateFile)
	if err != nil {
		a.emitError(err)
		return *new(CurrentIssue)
	}
	return *issue
}

func (a *App) GetConfig() Config {
	config, err := getFile[Config](configFile)
	if err != nil {
		a.emitError(err)
		fmt.Println("error getting config in go", err.Error())
		return *new(Config)
	}
	return *config
}

func (a *App) ForceError() {
	_, err := strconv.Atoi("pizza")
	if err != nil {
		a.emitError(err)
	}
}

func startStop(update IssueUpdate, ts *time.Time) error {
	issue, err := getFile[CurrentIssue](stateFile)
	if err != nil {
		issue = &CurrentIssue{update.IssueKey, false, ts.Format(iso8601Layout)}
	}
	if issue.Started {
		// time to end it
		startTime, err := time.Parse(iso8601Layout, issue.StartTime)
		if err != nil {
			return err
		}
		config, err := getFile[Config](configFile)
		if err != nil {
			return err
		}
		err = addWorkLog(update, &startTime, ts, config)
		if err != nil {
			return err
		}
		err = setFile(stateFile, new(CurrentIssue))
		if err != nil {
			return err
		}
	} else {
		// started new task / work segment
		issue.IssueKey = update.IssueKey
		issue.Started = true
		issue.StartTime = ts.Format(iso8601Layout)
		err = setFile(stateFile, issue)
		if err != nil {
			return err
		}
	}
	return nil
}
