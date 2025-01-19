package main

import (
	"f-commerce/infra"
	"f-commerce/route"
	"fmt"
	"log"
	"net/http"
)

func main() {

	ctx, err := infra.NewIntegrateContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
		return
	}

	r := route.NewRoutes(ctx)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", ctx.Cfg.Port),
		Handler: r,
	}

	log.Println("Server Running On Port : " + ctx.Cfg.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
