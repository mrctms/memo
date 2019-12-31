# memo

With this program you can save your memo from the command line.

# Installation on Linux

Download <a href="https://github.com/MarckTomack/memo/releases">here</a>

```
cd memo
chmod+x install.sh
sudo ./install.sh
```

# Usage

You can use this commands:
```
Usage:
  memo [command]

Available Commands:
  add         Add a memo
  delete      Delete a memo
  edit        Edit a memo
  help        Help about any command
  show        Show all memo

Flags:
  -h, --help   help for memo

Use "memo [command] --help" for more information about a command.

```
# Build from source

You need a <a href="https://golang.org/dl/">Go compiler</a>

```
git clone https://github.com/MarckTomack/memo.git
cd memo
go build 
```
