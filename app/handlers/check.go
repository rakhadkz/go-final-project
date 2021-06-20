package handlers

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"mizu2281/mizus-goservice/foundation/web"
	"net/http"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
		// panic("forcing panic")
	}

	//localhost:3000/readiness/10
	mp := web.Params(r)
	n := mp["id"]
	log.Println(n)

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
