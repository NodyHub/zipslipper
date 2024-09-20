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

## Usage example

```shell
(main) ~/git/zipslipper% make && ./zipslipper go.mod ../../go.mod test.zip
(main) ~/git/zipslipper% unzip -l test.zip
Archive:  test.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  09-20-2024 09:41   sub/
        3  09-20-2024 09:41   sub/root
        3  09-20-2024 09:41   sub/root/outside
        3  09-20-2024 09:41   sub/root/outside/0
        3  09-20-2024 09:41   sub/root/outside/0/1
      103  09-20-2024 08:39   sub/root/outside/0/1/go.mod
---------                     -------
      115                     6 files
```
