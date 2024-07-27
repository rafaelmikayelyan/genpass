# genpass

Quick and simple CLI tool for password generation with clipboard support.

By default, genpass produces a 16 character long password consisting of uppercase and lowercase letters, numbers and symbols, and automatically adds it to your clipboard.

It can be customized using several optional parameters.

## Usage

genpass [-lnrsu] [int]

-l : use lowercase\
-u : use uppercase\
-n : use numbers\
-s : use symbols\
-r : use redused set of symbols `!@#$%^&*\` which correspond to `SHIFT`+`12345678`.

int : length of the password

For example:
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
