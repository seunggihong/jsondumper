package main

import (
	"flag"
	"fmt"
	"log"

	"jsondumper/package/req"
	"jsondumper/package/yaml_reader"
)

func main() {
	config_path := flag.String("path", "config.yaml", "Path to config YAML")
	flag.Parse()

	config, err := yaml_reader.LoadConfig(*config_path)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	for _, target := range config.Prometheus.Target {
		for _, template := range config.Prometheus.QueryTemplates {
			query := req.BuildQuery(template, target)
			resp, err := req.QueryPrometheus(config.Prometheus.ServerIP, config.Prometheus.Port, query)
			if err != nil {
				log.Printf("Query failed: %v", err)
				continue
			}

			filename := fmt.Sprintf("%s_%s.json", template.FilenameSuffix, TargetName(target))
			err = req.SaveJSONResponse(resp, config.Prometheus.OutputDir, filename)
			if err != nil {
				log.Printf("Failed to save JSON: %v", err)
			}
		}
	}
}

func TargetName(t yaml_reader.TargetConfig) string {
	if t.Type == "pod" {
		return fmt.Sprintf("%s_%s", t.Namespace, t.PodName)
	} else if t.Type == "node" {
		return t.NodeName
	}
	return "unknown"
}
