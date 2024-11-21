package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HttpError struct {
	status  int
	code    string
	details *string
}

func (e HttpError) Error() string {
	return fmt.Sprintf("%v %s %v", e.status, e.code, e.details)
}

type RequestContext[T any] struct {
	Request *http.Request
	Body    *T
}

type HandlerRules[T any, R any] struct {
	requestValidator func(r *http.Request) error
	requestParser    func(r *http.Request) (RequestContext[T], error)
	responseWriter   func(w http.ResponseWriter, value R, err error)
}

type NoBody struct{}

func NewJsonHandlerRules[T any, R any](validators ...func(r *http.Request) error) HandlerRules[T, R] {
	return HandlerRules[T, R]{
		requestValidator: func(r *http.Request) error {
			errs := ""
			for _, v := range validators {
				errs = fmt.Sprintf("%s %v", errs, v(r))
			}
			if len(errs) > 0 {
				return fmt.Errorf(errs)
			}
			return nil
		},
		requestParser:  jsonParseRequest[T],
		responseWriter: jsonResponseWriter[R],
	}
}

func GetHandler[R any](rules HandlerRules[NoBody, R], handler func(ctx RequestContext[NoBody]) (R, error)) http.HandlerFunc {
	return outHandler(rules, handler)
}

func PostHandler[T any, R any](rules HandlerRules[T, R], handler func(ctx RequestContext[T]) (R, error)) http.HandlerFunc {
	return inOutHandler(rules, handler)
}

func inOutHandler[T any, R any](rules HandlerRules[T, R], handler func(ctx RequestContext[T]) (R, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res R
		var err error
		err = rules.requestValidator(r)
		if err != nil {
			rules.responseWriter(w, res, err)
			return
		}
		var ctx RequestContext[T]
		ctx, err = rules.requestParser(r)
		if err != nil {
			rules.responseWriter(w, res, err)
			return
		}
		res, err = handler(ctx)
		if err != nil {
			rules.responseWriter(w, res, err)
			return
		}
		rules.responseWriter(w, res, nil)
	}
}

func outHandler[R any](rules HandlerRules[NoBody, R], handler func(ctx RequestContext[NoBody]) (R, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res R
		var err error
		err = rules.requestValidator(r)
		if err != nil {
			rules.responseWriter(w, res, err)
			return
		}
		var ctx RequestContext[NoBody]
		res, err = handler(ctx)
		if err != nil {
			rules.responseWriter(w, res, err)
			return
		}
		rules.responseWriter(w, res, nil)
	}
}

func jsonParseRequest[T any](r *http.Request) (RequestContext[T], error) {
	var body T
	err := json.NewDecoder(r.Body).Decode(&body)
	res := RequestContext[T]{Request: r, Body: &body}
	if err != nil {
		details := fmt.Sprint(err)
		return res, HttpError{status: http.StatusBadRequest, code: "Body conversion failed", details: &details}
	} else {
		return res, nil
	}
}

func jsonResponseWriter[R any](w http.ResponseWriter, obj R, err error) {
	if err != nil {
		var httpError HttpError
		switch {
		case errors.As(err, &httpError):
			w.WriteHeader(httpError.status)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		err2 := json.NewEncoder(w).Encode(obj)
		if err2 != nil {
			panic(err2)
		}
	}
}
