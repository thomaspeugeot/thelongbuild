# The Long Build

This repo is for the reproduction of a long build time for 1500 lines of go.

## Measuring the build time

It is not straightfoward to reproduce a build time. The go compiler seems to be full
of caches and the go clean *package name* does not seems to work.

Therefore, the procedure is to refactor the name of a function/struct and run the 
command `time go build -v github.com/thomaspeugeot/thelongbuild/go/tree` and then discard the 
change.

On a MacBook Pro 2019

```
% time go build -v github.com/thomaspeugeot/thelongbuild/go/tree
github.com/thomaspeugeot/thelongbuild/go/tree
go build -v github.com/thomaspeugeot/thelongbuild/go/tree  44.29s user 2.28s system 376% cpu 12.380 total
```

45'' for 1500 lines of code seems a lot for go build

## Analysis

Sincere apologies for:

- not being able to reproduce the issue on a smaller synthetic code base. I tried for many hours to construe a similar case but to no avail.
- having obfuscated the code. The model is proprietary therefore it was a requirement for sharing the issue.

The package of interest is in the directory "go/tree". It is one file of 1500 lines of code.

The issue arised when coding this package. First, the build time went strangely to 11'', then after  more coding to 22', then with more complexity to 57'' then with even more coding, it went down to 44'' (I expected to build time to rise even more).

Fast build time is a wonderful feature of go. I hope this report can help in this direction.