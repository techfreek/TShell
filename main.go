package main

import (
	"fmt"
	//"strings"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("TShell")
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

func main() {
   fmt.Println("Hello world!")
   log.Info("Hello World");
}
