package api

import (
	. "github.com/seprich/keycape/internal/logger"
	"github.com/seprich/keycape/internal/util"
)

type HelloReqBody struct {
	What bool `json:"what"`
}

type HelloRespBody struct {
	Meh bool `json:"meh"`
}

func PostHelloHandler(ctx util.RequestContext[HelloReqBody]) (HelloRespBody, error) {
	Logger.Info("Request Body: ", "requestBody", ctx.Body)
	return HelloRespBody{Meh: true}, nil
}
