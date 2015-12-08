# wso2Scripts

A collection of scripts which help with various WSO2 deployment/development/maintainence
tasks.

These scripts are basically bash scripts for basic repetitive tasks that I encounter
on a frequent basis. They are all written in Go as an experiment of sorts in using
go scripts over bash scripts as they are both easier to write and easier to read.

## Building

To build for your platform run `make`. This will build all scripts for your platform

To build for all supported platforms run `make cross`.

Currently the supported platforms are:
* linux/amd64
* linux/386
* darwin/amd64

If you have need for any other supported platforms please submit a pull request.
Currently the supported platforms are minimal only due to the fact that I personally
don't have any need for other platforms.

All binaries are found in `target/<os>/<arch>/` for the desired os and arch. 
