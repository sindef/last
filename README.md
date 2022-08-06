## About ##

A quick go script to print the last n lines of a file with an offset, given by o

## Usage ##

```
  -f string
        File to read
        
  -n int
        Number of lines (default 10)
  -o int
        Offset (default 0)
```
For example, you could run `last -f testfile -n 50 -o 5` to print the last 50 lines of a file, with an offset of 5 (i.e. starting at the 55th last line, and printing to the 5th last)

This is essentially `tail` piped into `head` but in Go

