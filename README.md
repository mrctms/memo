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

You can use this command:
```
memo -h
Usage of memo:
  -a	[memo] | To add a memo
  -ash
    	[long memo] [shorted memo] | Add a shorted memo
  -d	[id] | To delete a memo
  -da
    	To delete all memo
  -m	[id] [new memo] | To edit a memo
  -msh
    	[id] [new shorted memo] | To edit the memo behind the shorted memo
  -r	[id] | Show the complete memo
  -s	To show all memo
```

For example,<br>
If you want add a memo:
```
memo -a "buy some cookies"
memo -ash "long string" "short string"
```
If you want delete a single memo:
```
memo -s
1 - buy some cookies
2 - https://github.com/MarckTomack/memo
3 - buy new car
4 - link README
memo -r 4
4 - https://github.com/MarckTomack/memo/README.md

memo -d 2
memo -s
1 - buy some cookies
3 - buy new car
4 - link README
```

To delete multiple memo:
```
memo -d 1 2 3 4
```

# Build from source

You need a <a href="https://golang.org/dl/">Go compiler</a>

```
git clone https://github.com/MarckTomack/memo.git
cd memo
go build 
```
