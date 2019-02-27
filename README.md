# dead_simple_mixin_robot_golang
dead simple Mixin Network robot wrote in go lang

### download and install go in your computer
[Install go lang](https://golang.org/doc/install)

### check your install
```bash
$ go version
go version go1.10 darwin/amd64
```
### Check your $GOPATH
```bash
$ echo $GOPATH
/Users/linli/go

$ ls $GOPATH
bin pkg src

```

### clone all source code into $GOPATH/src/github.com/user
```bash
$ ls $GOPATH/src/github.com/user
dead_simple_mixin_robot_golang hello                          mixin-sdk-go                   mixinrobot                     stringutil
$ cd $GOPATH/src/github.com/user
$ git clone https://github.com/myrual/dead_simple_mixin_robot_golang.git
```

### list all clone file
```bash
$ ls $GOPATH/src/github.com/user/dead_simple_mixin_robot_golang
LICENSE   README.md config    main.go
```

### update config/config.go

### install library by go get
```bash
$ go get github.com/MooooonStar/mixin-sdk-go/messenger
```


### build local
```bash
$ cd $GOPATH/src/github.com/user/dead_simple_mixin_robot_golang
$ go build main.go
```

## build and install by go global build command
