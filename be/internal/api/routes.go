package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/seprich/keycape/internal/util"
)

func ApiRoutes(r chi.Router) {
	r.Post(
		"/hello",
		util.PostHandler(
			util.NewJsonHandlerRules[HelloReqBody, HelloRespBody](),
			PostHelloHandler))
	r.Get(
		"/fut",
		util.GetHandler(util.NewJsonHandlerRules[util.NoBody, TestRespBody](),
			GetTestHandler))
}
