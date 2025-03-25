package req

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"jsondumper/package/yaml_reader"
)

func BuildLabelSelector(labels []yaml_reader.LabelDefinition, target yaml_reader.TargetConfig) string {
	var selectors []string
	for _, l := range labels {
		var value string
		switch l.ValueFrom {
		case "pod_name":
			value = target.PodName
		case "namespace":
			value = target.Namespace
		case "node_name":
			value = target.NodeName
		}
		selectors = append(selectors, fmt.Sprintf(`%s%s"%s"`, l.Key, l.Match, value))
	}
	return strings.Join(selectors, ", ")
}

func BuildQuery(template yaml_reader.QueryTemplate, target yaml_reader.TargetConfig) string {
	labels := BuildLabelSelector(template.Labels, target)
	return fmt.Sprintf(template.Expression, labels)
}

func QueryPrometheus(serverIP, port, query string) ([]byte, error) {
	escaped_query := url.QueryEscape(query)
	url := fmt.Sprintf("http://%s:%s/api/v1/query?query=%s", serverIP, port, escaped_query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func SaveJSONResponse(data []byte, outputDir, filename string) error {
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}
	path := filepath.Join(outputDir, filename)
	return ioutil.WriteFile(path, data, 0644)
}
