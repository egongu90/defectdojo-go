# defectdojo-go
DefectDojo Go Client

This library interact with DefectDojo API to upload scan results.

Tested on DefectDojo API v2.

Disclaimer:

This proyect was just roughly 5h of work for the first time coding in golang learning from 0.
It may and probably won't be maintained.

## Installation

Compiled release

```sh
VERSION="v1.0.0"
wget -O /usr/local/bin/defectdojo-go https://github.com/xNaaro/defectdojo-go/releases/download/$VERSION/binary-linux-amd64
chmod +x /usr/local/bin/defectdojo-go
```

With go install

```sh
go install github.com/xNaaro/defectdojo-go@latest
```

## Usage

Requires the following environment variables for authentication:

```sh
export DEFECTDOJO_URL="https://demo.defectdojo.org/api/v2"
export DEFECTDOJO_USERNAME="admin"
export DEFECTDOJO_PASSWORD="1Defectdojo@demo#appsec"
```

Command options

```sh
Import scan results to Defect Dojo.

Usage:
  defectdojo-go importScan [flags]

Aliases:
  importScan, imp

Flags:
      --active string               Active status. (default "true")
      --check_list string           Check list. (default "true")
      --close_old_findings string   Close old findings. (default "false")
      --engagement_name string      Engangement name to upload report.
      --file_name string            File name or absolute path to upload. (default "results.json")
  -h, --help                        help for importScan
      --minimum_severity string     Minimum severity. (default "Info")
      --product_name string         Product name to upload report.
      --push_to_jira string         Push to Jira. (default "false")
      --scan_date string            Scan date. (default "2025-03-22")
      --scan_type string            Scan type, one of the supported by DefectDojo. Case and space sensitive. (default "Bandit Scan")
      --status string               Status. (default "Not Started")
      --verified string             Verified scan. (default "true")
```

Example usage with bare minimum parameters required

```sh
$ defectdojo-go importScan --product_name "Apple Accounting Software" --engagement_name "test" --file_name 'results.json' --scan_type "Bandit Scan"       
Authenticating...
Import scan
Status code: 201
```