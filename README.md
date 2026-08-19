# Data Pipeline Engine

Build and manage data pipelines with ease - ETL, streaming, and batch processing.

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![CI](https://github.com/Qyroxen/data-pipeline-engine/actions/workflows/ci.yml/badge.svg)](https://github.com/Qyroxen/data-pipeline-engine/actions/workflows/ci.yml)

> Build and manage data pipelines with ease - ETL, streaming, and batch processing.

## What is it?

Data Pipeline Engine is a command-line tool built with Go that helps developers build and manage data pipelines with ease - etl, streaming, and batch processing. It's designed to be fast, reliable, and easy to use.

## Why?

Every developer needs data pipeline engine — but existing tools are either too complex, too slow, or require cloud dependencies. We built Data Pipeline Engine to be:
- **Fast** — Written in Go for maximum performance
- **Offline** — No cloud dependencies, your data stays on your machine
- **Simple** — Clean CLI interface with sensible defaults
- **Extensible** — Easy to customize and integrate into your workflow

## Features

- Visual pipeline builder
- Multiple data source support
- Transformation engine
- Scheduling and orchestration
- Monitoring and alerting
- CLI management

## Quick Start

### Prerequisites

- Go 1.23 or later

### Install

```bash
# Install with go install
go install github.com/Qyroxen/data-pipeline-engine@latest

# Or build from source
git clone https://github.com/Qyroxen/data-pipeline-engine.git
cd data-pipeline-engine
go build -o data-pipeline-engine .
```

### Usage

```bash
# Basic usage
.data-pipeline-engine --help

# Example
./data-pipeline-engine create --name etl-pipeline --source csv --dest postgres
```

## Output

```
Data Pipeline Engine v1.0.0

Scanning...

✓ Analysis complete
✓ Results ready

{
  "status": "success",
  "results": [...]
}
```

## Configuration

Create a `.config.yaml` file in your project root:

```yaml
# Configuration options
verbose: true
output: json
timeout: 30s
```

## CLI Flags

```
data pipeline engine [command]

Flags:
  --path string      Target path (default ".")
  --format string    Output format: json, text (default "text")
  --verbose          Enable verbose output
  --config string    Config file path
  --output string    Output file path
```

## Examples

### Basic Example

```bash
.data-pipeline-engine --path ./src
```

### Advanced Example

```bash
.data-pipeline-engine --path ./src --format json --output report.json --verbose
```

### CI/CD Integration

```yaml
# .github/workflows/ci.yml
- name: Run Data Pipeline Engine
  run: |
    go install github.com/Qyroxen/data-pipeline-engine@latest
    data-pipeline-engine --path . --format json --output report.json
```

## Documentation

- [Getting Started](docs/getting-started.md)
- [Configuration](docs/configuration.md)
- [API Reference](docs/api-reference.md)
- [Examples](examples/)
- [Contributing](CONTRIBUTING.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Qyroxen** - [GitHub](https://github.com/Qyroxen)

---

**Found this useful?** Give it a ⭐ on GitHub!
