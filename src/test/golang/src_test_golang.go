package main

import (
	"github.com/bitwormhole/gitlib/src/test/golang/cfg"
	"github.com/bitwormhole/starter"
)

func main() {
	starter.InitApp().Use(cfg.ExportModule()).Run()
}
