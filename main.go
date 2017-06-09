package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.WriteString("hello world")
	})

	app.Run(iris.Addr(":8080"))
}
