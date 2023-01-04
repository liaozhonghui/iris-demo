package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.Default()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Use(myMiddleware)

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})
	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userId, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userId)
	})

	app.Run(iris.Addr(":8080"))
}
func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
