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
	ServerIP       string          `yaml:"server_ip"`
	Port           string          `yaml:"port"`
	OutputDir      string          `yaml:"output_dir"`
	QueryTemplates []QueryTemplate `yaml:"query_templates"`
	Target         []TargetConfig  `yaml:"target"`
}

type QueryTemplate struct {
	Name           string            `yaml:"name"`
	Expression     string            `yaml:"expression"`
	Labels         []LabelDefinition `yaml:"labels"`
	FilenameSuffix string            `yaml:"filename_suffix"`
}

type LabelDefinition struct {
	Key       string `yaml:"key"`
	ValueFrom string `yaml:"value_from"`
	Match     string `yaml:"match"`
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
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	return &config, nil
}
