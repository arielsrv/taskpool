# golang-toolkit
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-83.4%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)

## Developer tools

- [Golang Lint](https: //golangci-lint.run/)
- [Golang Task](https: //taskfile.dev/)
- [Golang Dependencies Update](https://github.com/oligot/go-mod-upgrade)
- [jq](https:                        //stedolan.github.io/jq/)

### For macOs

```shell
$ brew install go-task/tap/go-task
$ brew install golangci-lint
$ go install github.com/oligot/go-mod-upgrade@latest
$ brew install jq
```

## Table of contents

# Installation

```sh
go get -u github.com/arielsrv/taskpool
```

# ⚡️ Quickstart


```go
package main

import (
	"github.com/arielsrv/golang-toolkit/task"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var future1, future2, future3 *task.Task[int]

	tb := &task.Builder{
		MaxWorkers: runtime.NumCPU() - 1,
	}

	start := time.Now()
	tb.ForkJoin(func(c *task.Awaitable) {
		future1 = task.Await[int](c, GetNumber)
		future2 = task.Await[int](c, GetNumber)
		future3 = task.Await[int](c, GetNumber)
	})

	log.Println(future1.Result)
	log.Println(future2.Result)
	log.Println(future3.Result)

	end := time.Since(start)
	log.Println(end)
}

func GetNumber() (int, error) {
	value := rand.Int()
	time.Sleep(time.Millisecond * 1000)
	log.Println("done ...")
	return value, nil
}


```