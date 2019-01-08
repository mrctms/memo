# memo

With this program you can save your memo from the command line.

# Installation on Linux

Download <a href="https://github.com/MarckTomack/memo/releases/">here</a> and extract the folder.

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
memo -h
Usage of memo:
  -a	[memo] | To add a memo
  -ash
    	[long memo] [shorted memo] | Add a shorted memo
  -d	[position number] | To delete a memo
  -da
    	To delete all memo
  -m	[position number] | To edit a memo
  -msh
    	[position number] | To edit the memo behind the shorted memo
  -r	[position number] | Show the complete memo
  -s	To show all memo
</code>
</pre>

For example,<br>
If you want add a memo:
<pre>
<code>
memo -a "buy some cookies"
memo -ash "long string" "short string"
</code>
</pre>
If you want delete a single memo:
<pre>
<code>
memo -s
1 - buy some cookies
2 - https://github.com/MarckTomack/memo
3 - buy new car
4 - link README
memo -r 4
4 - https://github.com/MarckTomack/memo/edit/master/README.md
<br>
memo -d 2
memo -s
1 - buy some cookies
3 - buy new car
4 - link README
</code>
</pre>

# Build from source

<pre>
go get github.com/mattn/go-sqlite3
go build memo.go
</pre>
