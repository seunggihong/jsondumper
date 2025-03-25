package yaml_reader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Prometheus PrometheusConfig `yaml:"prometheus"`
}

type PrometheusConfig struct {
	ServerIP     string         `yaml:"server_ip"`
	Port         string         `yaml:"port"`
	Query        []string       `yaml:"query"`
	Target       []TargetConfig `yaml:"target"`
	Interval     int            `yaml:"interval"`
	Step         int            `yaml:"step"`
	StoragePath  string         `yaml:"storage_path"`
	DatafileName string         `yaml:"datafile_name"`
}

type TargetConfig struct {
	Type      string `yaml:"type"`
	PodName   string `yaml:"pod_name"`
	Namespace string `yaml:"namespace"`
	NodeName  string `yaml:"node_name"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}
