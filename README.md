# genpass

genpass helps to quickly generate passwords in CLI.

By default it produces a 16 character long password consisting of uppercase and lowercase letters, numbers and symbols.

It can be customized using several optional parameters.

Usage: genpass [-lnrsu] [int]

-l : use lowercase\
-u : use uppercase\
-n : use numbers\
-s : use symbols\
-r : use redused set of symbols : !@#$%^&*\

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
