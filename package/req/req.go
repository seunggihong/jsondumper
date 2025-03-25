package req

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"jsondumper/package/yaml_reader"
)

func QueryPrometheus(configPath, promQL string) (string, error) {
	config, err := yaml_reader.LoadConfig(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	url := fmt.Sprintf("http://%s:%s/api/v1/query?query=%s", config.Prometheus.ServerIP, config.Prometheus.Port, promQL)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
