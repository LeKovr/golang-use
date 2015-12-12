package main

import (
  "gopkg.in/macaron.v1"
  "github.com/skratchdot/open-golang/open"
)

func main() {
    m := macaron.Classic()
    m.Use(macaron.Renderer())

    m.Get("/", func(ctx *macaron.Context) {
        ctx.Data["Name"] = "jeremy"
        ctx.HTML(200, "hello") // 200 is the response code.
    })

    open.Run("http://localhost:4000/")
    m.Run()
}
