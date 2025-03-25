package main

import (
	"flag"
	"fmt"
	"log"

	"jsondumper/package/req"
)

func main() {
	configPath := flag.String("path", "config.yaml", "Path to the YAML configuration file")
	flag.Parse()

	query := "kube_node_info"

	response, err := req.QueryPrometheus(*configPath, query)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	fmt.Println("Prometheus Response:")
	fmt.Println(response)
}
