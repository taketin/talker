talker
======

## Description

This script is choose a talker.

Please use to study sessions presenter extraction, etc.

It uses the dropbox in storage, it does not overlap anyone is running.

## Usage

```
$ talker select

> taketin
```

if all members selected

![screenshots1](screenshots/next_round.png)

Other usage please look at the help in `talker -h`

## Install

```
$ go get github.com/taketin/talker
$ go get github.com/tools/godep
$ cd {GO_ROOT}/src/github.com/taketin/talker
$ godep get
$ mkdir -p $GOPATH/config/talker
$ mv config.tml $GOPATH/config/talker/
$ vi config.tml # edit your settings
$ go install
$ which talker
```

how using ghq.

```
$ ghq get github.com/taketin/talker
$ ghq get http://github.com/tools/godep
$ ghq look talker
$ godep get
$ mkdir -p $GOPATH/config/talker
$ mv config.tml $GOPATH/config/talker/
$ vi config.tml # edit your settings
$ go install
$ which talker
```

## Contribution

1. Fork it
1. Create your feature branch (git checkout -b my-new-feature)
1. Commit your changes (git commit -am 'Add some feature')
1. Push to the branch (git push origin my-new-feature)
1. Create new Pull Request

## Author

@taketin
