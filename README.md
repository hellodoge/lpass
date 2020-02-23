# lpass
lpass is an easy-to-use alternative to using the same password for multiple websites or all your accounts.

## Installation
You need to have Go 1.10 or newer on your machine.
```
go get -u github.com/hellodoge/lpass
go get -u golang.org/x/crypto/pbkdf2
cd $HOME/go/src/github.com/hellodoge/lpass
go install
```

## Usage
```
Usage: lpass [master key] [token or url] [args]
  -0	do not use digit characters
  -A	do not use uppercase characters
  -_	do not use special characters
  -a	do not use lowercase characters
  -i int
    	iteration count (default 65536)
  -l int
    	password length (default 15)
```

## License
This code is released under an MIT License. You are free to use, modify, distribute, or sell it under those terms.
