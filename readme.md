# Extended Logger
Sample extended logger for Go with 3 log levels support and buffer as output destination.

## Description
LogExtended is a wrapper around Go's standard log package that adds log level functionality. 


## Features

3 log levels: Error, Warning, Info
Configurable filtering level


## Implementation Notes

The logger uses Go's standard log package internally
Log messages include file name and line number information
The filtering is done at the application level, not at the log package level
Each log level has an automatic prefix: "ERROR ", "WARN ", "INFO "

## Requirements

Go 1.11+ (for module support)
Standard library packages: bytes, fmt, log