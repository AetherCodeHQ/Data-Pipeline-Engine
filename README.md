# Data Pipeline Engine

![CI](https://github.com/Qyroxen/Data-Pipeline-Engine/actions/workflows/ci.yml/badge.svg) ![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![License](https://img.shields.io/badge/License-MIT-yellow.svg) ![Stars](https://img.shields.io/github/stars/Qyroxen/Data-Pipeline-Engine?style=social)

> Build and run data pipelines with a simple YAML configuration

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Data-Pipeline-Engine?style=social)](https://github.com/Qyroxen/Data-Pipeline-Engine/stargazers)

## What is it?

Data Pipeline Engine lets you define data transformations in YAML and run them as efficient Go pipelines.

## Why should you care?

ETL processes are complex. This tool makes them simple with YAML configuration.

## Demo

```bash
./data-pipeline run --config pipeline.yaml
```

**Output:**
```
Pipeline completed:
  - 10,000 records processed
  - 3 transformations applied
  - Output: output.csv
```

## Features

- YAML-based pipeline definition
- Parallel processing
- Error handling and retries
- Multiple data sources
- Monitoring and logging

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Data-Pipeline-Engine.git
cd Data-Pipeline-Engine
go build -o data-pipeline .

# Run
./data-pipeline --config pipeline.yaml
```

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--config` | Pipeline config file | `pipeline.yaml` |
| `--dry-run` | Validate without running | `false` |
| `--parallel` | Parallel workers | `4` |
| `--retry` | Retry failed steps | `3` |

## Examples

# Run pipeline
./data-pipeline run --config pipeline.yaml

# Dry run
./data-pipeline run --config pipeline.yaml --dry-run

# With parallelism
./data-pipeline run --config pipeline.yaml --parallel 8

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Data-Pipeline-Engine/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Data-Pipeline-Engine?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Data-Pipeline-Engine/network/members">
    <img src="https://img.shields.io/github/forks/Qyroxen/Data-Pipeline-Engine?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Data-Pipeline-Engine/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Data-Pipeline-Engine" alt="Issues">
  </a>
</p>
