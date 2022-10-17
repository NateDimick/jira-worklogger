package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	appDir     string = ".worklogger"
	configFile string = "config.json"
	stateFile  string = "issue.json"
)

func getAppDirPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, appDir), nil
}

func createAppDir() error {
	appDirPath, err := getAppDirPath()
	if err != nil {
		return err
	}
	if _, err := os.Stat(appDirPath); os.IsNotExist(err) {
		fmt.Println("creating app directory")
		err = os.Mkdir(appDirPath, os.ModeDir)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("skipping app directory creation")
	}
	return nil
}

func getFile[T any](fileName string) (*T, error) {
	appDirPath, err := getAppDirPath()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(appDirPath, fileName)
	fBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := new(T)
	err = json.Unmarshal(fBytes, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func setFile(fileName string, data interface{}) error {
	appDirPath, err := getAppDirPath()
	if err != nil {
		return err
	}
	path := filepath.Join(appDirPath, fileName)
	fBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, fBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
