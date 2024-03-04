# timex

Time helper for Golang.

## Install

```shell
go get github.com/innotechdevops/timex
```

## How to use

- Parse by date ISO

```go
datetime, _ := timex.ParseBy("2023-12-03T13:51:06.474Z", timex.DateTimeFormatISO)
```

- Parse by GMT7

```go
date, _ := timex.ParseByGMT7(datetime, timex.TimeFormatDash1)
```