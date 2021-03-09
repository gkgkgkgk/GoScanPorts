# GoScanPorts
A very simple golang port scanner.

## How to use:
Build the code with `go build portscanner.go` or use the prebuilt binary. Run the code: `./portscanner`

## Flags:
- -host (string):
        Specified Host (default "127.0.0.1")
- -p (string):
        Specified Ports (default "0:1024")
- -t (int):
        Specified Timeout (default 50)
