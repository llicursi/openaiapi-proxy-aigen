package handlers

import (
	"log"
)

// ApiResponse structs
type ApiResponse struct {
	Message string `json:"message"`
	Model   string `json:"model"`
	Object  string `json:"object"`
}

// Config holds all configuration
type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Logging LoggingConfig `yaml:"logging"`
}

type ServerConfig struct {
	Port       int    `yaml:"port"`
	ForwardURL string `yaml:"forward_url"`
}

type LoggingConfig struct {
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

// Handler holds the configuration for request handling
type Handler struct {
	config Config
	logger *log.Logger
}
