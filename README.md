# memo

With this program you can save your memo from the command line.

# Installation on Linux

Download <a href="https://github.com/MarckTomack/memo/releases/tag/v1.0">here</a> and extract the folder.

<pre>
<code>
cd memo
chmod+x install.sh
sudo ./install.sh
</code>
</pre>

# Usage

You can use this command:
<pre>
<code>
a - To add a memo
d position number- To delete a memo
da - To delete all memo
s - To show all memo
</code>
</pre>

For example,<br>
If you want add a memo:
<pre>
<code>
memo a "buy some cookies"
</code>
</pre>
If you want delete a single memo:
<pre>
<code>
memo s
1 - buy some cookies
2 - https://github.com/MarckTomack/memo
3 - buy new car
<br>
memo d 2
memo s
1 - buy some cookies
2 - buy new car
</code>
</pre>

# Build from source

<pre>
go get github.com/mattn/go-sqlite3
go get github.com/mitchellh/go-homedir
go build memo.go
</pre>
