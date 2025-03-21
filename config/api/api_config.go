package api

import (
	"TestProject/router"
	"golang.org/x/sync/errgroup"
)

func StartServer() {
	var g errgroup.Group

	// 執行 Rest Server
	g.Go(func() error {
		r := router.SetupRouter()
		err := r.Run(":8080")
		if err != nil {
		}
		return err
	})

	if err := g.Wait(); err != nil {
	}
}
