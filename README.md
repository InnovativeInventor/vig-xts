## vig-xts
An implementation of the Vigenere cipher with ciphertext stealing in Go (Golang).

## Running
To compile and run, type
```
go run vig-xts.go [options]
```
To run the precompiled binaries, go to `bin/` and execute the binary appropriate for your platform.

Example (for Linux-based 64 bit computers):
```
./crypto_linux_amd64
```

## Options
```
$ go run vig-xts.go

NAME:
   Vig-xts - Use xts with the Vigenere cipher

USAGE:
   vig-xts [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     encrypt, e  Encrypt text
     decrypt, d  Decrypt text
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```