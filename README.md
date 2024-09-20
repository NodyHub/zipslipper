# zipslipper
Create tar/zip archives that try to exploit zipslip vulnerability.

## CLI Tool

You can use this library on the command line with the `zipslipper` command.

### Installation

```cli
go install github.com/NodyHub/zipslipper@latest
```

### Manual Build and Installation

```cli
git clone git@github.com:NodyHub/zipslipper.git
cd zipslipper
make
make install
```

## Usage

Basic usage on cli:

```shell
% zipslipper -h
Usage: zipslipper <input> <relative-path> <output-file> [flags]

A utility to build tar/zip archives that performs a zipslip attack.

Arguments:
  <input>            Input file.
  <relative-path>    Relative extraction path.
  <output-file>      Output file.

Flags:
  -h, --help                  Show context-sensitive help.
  -t, --archive-type="zip"    Archive type. (tar, zip)
  -v, --verbose               Verbose logging.
  -V, --version               Print release version information.
```
