## vig-xts
An implementation of the Vigenere cipher with ciphertext stealing in Go (Golang).

## Running
To compile and run, type
```
go run main.go [options]
```
To run the precompiled binaries, go to `bin/` and execute the binary appropriate for your platform.

Example (for Linux-based 64 bit computers):
```
./vig-xts_linux_amd64
```

## Options
```
$ go run main.go

NAME:
   Vig-xts - Use xts with the Vigenere cipher

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.1.5

COMMANDS:
     encrypt, e  Encrypt text
     decrypt, d  Decrypt text
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```