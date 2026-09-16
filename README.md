# genpass

Quick and simple CLI tool for password generation with clipboard support.

By default, genpass produces a 16 character password consisting of uppercase and lowercase letters, numbers and symbols, and automatically adds it to your clipboard.

It can be customized using several optional arguments.

## Motivation

Generating a secure password should be as simple as a short command.

## Quick Start

### Install

1. You need [Go](https://go.dev/doc/install) to compile (and to optionally download and install) the app.
2. Use `go install github.com/rafaelmikayelyan/genpass/cmd/genpass@latest` to install the app.

### Generate Password

Run `genpass` without any arguments to generate a 16-character password containing lowercase and uppercase letters, numbers, and symbols, which is automatically copied to your clipboard:
```
$ genpass
=> 7RqFM-X"55H\8hrE
```
Run `genpass #` where # is a number specifying the desired length:
```
$ genpass 5
=> P+N9w
```
Run `genpass --help` to see all available options.

## Usage

`genpass [flags] &lt;count&gt;`

#### Flags (optional)
-l : use lowercase\
-u : use uppercase\
-R : use alphanumeric set, excluding lowercase L (l), uppercase i (I) and uppercase o (O) to increase readability.\
-n : use numbers\
-s : use symbols\
-r : use reduced set of symbols `!@#$%^&*` which correspond to <kbd>SHIFT</kbd> + <kbd>1</kbd>...<kbd>8</kbd>\
-v, --version&nbsp;: get version (must be used alone)\
-h, --help&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;: get description of available functionality (must be used alone)

#### Count (optional)
0... : length of the password

#### Defaults

Running `genpass` without any arguments is equivalent to `genpass -lnsu 16`, generating a 16-character password that includes lowercase and uppercase letters, numbers, and symbols.

### Examples

```
$ genpass -n 4
=> 6790
```
```
$ genpass -ul
=> pvyzJpQRTPMeqzbg
```
```
$ genpass 20 -Rr
=> pLpC75q!%SVua496S05W
```

## Contributing

### Clone the repo

```bash
git clone https://github.com/rafaelmikayelyan/genpass
cd genpass
```

### Build the compiled binary

```bash
go build
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
