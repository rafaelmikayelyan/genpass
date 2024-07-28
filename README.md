# genpass

Quick and simple CLI tool for password generation with clipboard support.

By default, genpass produces a 16 character long password consisting of uppercase and lowercase letters, numbers and symbols, and automatically adds it to your clipboard.

It can be customized using several optional arguments.

## Usage

genpass [-lnrRsu] [int]

-l : use lowercase\
-u : use uppercase\
-R : use alphanumeric set, excluding lowercase L (l), uppercase i (I) and uppercase o (O) to increase readability.\
-n : use numbers\
-s : use symbols\
-r : use redused set of symbols `!@#$%^&*\` which correspond to <kbd>SHIFT</kbd> + <kbd>1</kbd>...<kbd>8</kbd>

int : length of the password

Additionally, you can use:\
-v, --version&nbsp;: get version\
-h, --help&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;: get description of available functionality

### Examples
```
$ genpass
=> 7RqFM-X"55H\8hrE
```
```
$ genpass -n 4
=> 6790
```
```
$ genpass -ul
=> pvyzJpQRTPMeqzbg
```
```
$ genpass 5
=> P+N=n
```
```
$ genpass 20 -Rr
=> pLpC75q!%SVua496S05W
```
