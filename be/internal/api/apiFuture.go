package api

import (
	"github.com/seprich/go-future/async"
	. "github.com/seprich/keycape/internal/logger"
	"github.com/seprich/keycape/internal/util"
	"math/rand/v2"
	"time"
)

type TestRespBody struct {
	Meh string `json:"geh"`
}

func GetTestHandler(ctx util.RequestContext[util.NoBody]) (TestRespBody, error) {

	// With light Future abstraction:
	fut1 := async.NewFuture1(asyncResultFn, "Future1")
	fut2 := async.NewFuture1(asyncResultFn, "Future2")
	res1, err1 := fut1.Await()
	res2, err2 := fut2.Await()
	if err1 != nil || err2 != nil {
		panic("unexpected")
	}
	Logger.Info("Results from executing futures", "res1", res1, "res2", res2)

	return TestRespBody{Meh: "Done"}, nil
}

func asyncResultFn(name string) (string, error) {
	Logger.Info("Start " + name)
	time.Sleep(time.Duration(rand.IntN(1000)) * time.Millisecond)
	Logger.Info("Done " + name)

	return "Return " + name, nil
}
