package main

import (
	fume "github.com/fumeapp/fiber"
	"github.com/tcampbppu/server/app"
)

func main() {
	fume.Start(app.Init(), fume.Options{})
}
