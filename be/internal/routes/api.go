package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/seprich/keycape/internal/util"
)

func RegisterAPIRoutes(r *gin.RouterGroup) {
	r.GET("/hello", util.ResultToResponse(helloHandler))
}

type HelloResponse struct {
	Some string `json:"some"`
}

func helloHandler(c *gin.Context) (HelloResponse, error) {
	d := "blaah"
	return HelloResponse{Some: "hello"}, util.HttpError{400, "what", &d}
	//return HelloResponse{Some: "hi"}, fmt.Errorf("Yeah gaah")
	//panic("aargh!!!")
	//return HelloResponse{Some: "hello"}, nil
}
