# JsonDumper

This project is a Go application that automates Prometheus monitoring queries. It reads a flexible YAML configuration file that defines query templates and monitoring targets (such as pods or nodes), dynamically builds valid PromQL expressions, sends them to a Prometheus server, and saves the responses in structured JSON files. It is designed for extensibility and can support any Prometheus-compatible metrics or targets.

---

## Features

- Dynamic query generation using YAML-based configuration
- Target support: pod, namespace, node, and regex patterns
- PromQL match operator support: `=`, `=~`, `!=`, `!~`
- Prometheus HTTP API integration
- JSON output for all responses
- Easy-to-extend query templates

---

## YAML Configuration

### Example: `config.yaml`

```yaml
prometheus:
  server_ip: 'YOUR_PROMETHEUS_SERVER_IP'
  port: 'YOUR_PROMETHEUS_SERVER_PORT'
  output_dir: './output'
  query_templates:
    - name: "EXAMPLE_NAME"
      expression: "YOUR_QUERY"
      labels:
        - key: "pod"
          value_from: "pod_name"
          match: "=~"
        - key: "namespace"
          value_from: "namespace"
          match: "="
      filename_suffix: "FILE_NAME"
  target:
    - type: "pod"
      pod_name: "TARGET_POD_NAME"
      namespace: "TARGET_NAMESPACE"
```
### Field Descriptions

📌***prometheus***

| Field | Type | Description |
|:--|:--|:--|
|server_ip |string |IP address of the Prometheus server|
|port |string |Port where Prometheus is running |
|output_dir |string |Directory path to save JSON output files|
|query_templates |list |A list of query template objects|
|target |list |A list of monitoring target|

📌***query_templates***

|Field |Type |Description|
|:--|:--:|:--|
|name |string| Logical name of the query|
|expression |string |PromQL template with %s for label selectors|
|labels |list |List of label definitions used in the selector|
|filename_suffix |string |Used to name output JSON files|

📌***labels (inside each query template)***

|Field| Type| Description|
|:--|:--|:--|
|key| string| Prometheus label key (e.g., pod, namespace)|
|value_from |string |Field name in target config to substitute (e.g., pod_name, namespace)|
|match |string| PromQL match operator: =, =~, !=, !~|

target

|Field| Type| Description|
|:--|:--|:--|
|type | string| Target type, e.g., pod or node|
|pod_name|string|Pod name or regex pattern (required if type is pod)|
|namespace|string|Kubernetes namespace (required for pod targets)|
|node_name|string|Node name (used if type is node)|

## Usage
### Build and Run
```bash
go build -o promquery
./promquery --path=config.yaml
```
Or run without building:
```bash
go run main.go --path=config.yaml
```