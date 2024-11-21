package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpError struct {
	Status  int     // http status code for the response
	Code    string  // either the textual version of status code or pre-defined enumerated value returned in response
	Details *string // Attach to logs - do not send in response
}

func (e HttpError) Error() string {
	return fmt.Sprintf("%v %s %v", e.Status, e.Code, e.Details)
}

type ResultingHandler[T any] func(c *gin.Context) (T, error)

func ResultToResponse[T any](fn ResultingHandler[T]) gin.HandlerFunc {
	return func(context *gin.Context) {
		r, e := fn(context)
		toResponse(context, &r, e)
	}
}

func toResponse[T any](c *gin.Context, r *T, e any) {
	if e != nil || r == nil {
		switch t := e.(type) {
		case HttpError:
			tinyCode := GenerateTinyId()
			c.Set("error_code", t.Code)
			c.Set("error_details", t.Details)
			c.Set("error_ref", tinyCode)
			c.JSON(t.Status, gin.H{"error": t.Code, "ref": tinyCode})
		default:
			panic(e)
		}
	} else {
		c.JSON(http.StatusOK, r)
	}
}
