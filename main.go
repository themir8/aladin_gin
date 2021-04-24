package main

import (
	"github.com/mirsaid-mirzohidov/aladin_gin/api"
)

func main() {

	apiserver := api.New()

	apiserver.Run(":8000") // listen and serve on 0.0.0.0:8000 (for windows "localhost:8000")
}
