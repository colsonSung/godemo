package main

import (
	"examples/sqldemo"
	"examples/apidemo"
)

func main() {
	sqldemo.MysqlTest()
	apidemo.ControllerTest()
}
