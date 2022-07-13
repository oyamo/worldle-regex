package config

import (
	"log"
	"os"
	"regexp"
)

const (
	DEV = iota
	PROD
)

type (
	Server struct {
		Addr string
		Port string
	}

	File struct {
		Path string
	}

	Templates struct {
		Path string
	}

	Config struct {
		File      File
		Templates Templates
		Server    Server
		Log       log.Logger
	}
)

func LoadConfig(path string, version int) (*Config, error) {
	// load config from env file
	// If router is in production mode, load from os env
	// If router is in development mode, load from local file
	cfg := &Config{
		File: File{
			Path: os.Getenv("WORD_FILE"),
		},
		Templates: Templates{
			Path: os.Getenv("TEMPLATE_PATH"),
		},
		Server: Server{
			Addr: os.Getenv("SERVER_ADDR"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Log: *log.New(os.Stdout, "- reegle -", log.LstdFlags),
	}

	if version == DEV {
		// load config from local env file

		file, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		// parse config from local env file
		content := string(file)
		wordFileRegx := regexp.MustCompile(`WORD_FILE=(.*)`)
		templatePathRegx := regexp.MustCompile(`TEMPLATE_PATH=(.*)`)
		serverAddrRegx := regexp.MustCompile(`SERVER_ADDR=(.*)`)
		serverPortRegx := regexp.MustCompile(`SERVER_PORT=(.*)`)

		// set config from local env file
		cfg.File.Path = wordFileRegx.FindStringSubmatch(content)[1]
		cfg.Templates.Path = templatePathRegx.FindStringSubmatch(content)[1]
		cfg.Server.Addr = serverAddrRegx.FindStringSubmatch(content)[1]
		cfg.Server.Port = serverPortRegx.FindStringSubmatch(content)[1]

	}

	return cfg, nil
}
