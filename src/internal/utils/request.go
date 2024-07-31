package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type RequestResult[T any] struct {
	Ok      bool
	Status  int
	Body    T
	Headers http.Header
}

func NewRequestResult[T any](ok bool, status int, body T, headers http.Header) *RequestResult[T] {
	return &RequestResult[T]{
		Ok:      ok,
		Status:  status,
		Body:    body,
		Headers: headers,
	}
}

func Request[ReqBody any, RespBody any](url string, method string, headers map[string]string, body ReqBody) (*RequestResult[RespBody], error) {
	defaultHeaders := map[string]string{
		"User-Agent":                   "GoGinTemplate/1.0.0",
		"Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept",
		"Cache-Control":                "no-cache",
	}

	var req *http.Request

	switch method {
	case "GET":
		request, err := http.NewRequest(
			http.MethodGet,
			url,
			nil,
		)
		if err != nil {
			return nil, err
		}
		req = request
		break
	case "PUT":
	case "PATCH":
	case "POST":
		requestBodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		request, err := http.NewRequest(
			http.MethodPost,
			url,
			bytes.NewBuffer(requestBodyBytes),
		)
		if err != nil {
			return nil, err
		}
		req = request
	default:
		return nil, errors.New("invalid method")
	}

	if req == nil {
		return nil, errors.New("request is nil")
	}

	for key, value := range defaultHeaders {
		req.Header.Set(key, value)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	ok := response.StatusCode <= 399

	switch response.Header.Get("Content-Type") {
	case "application/json":
		var responseBody RespBody
		err = json.Unmarshal(responseBodyBytes, &responseBody)
		if err != nil {
			return nil, err
		}
		return NewRequestResult[RespBody](ok, response.StatusCode, responseBody, response.Header), nil
	default:
		return nil, errors.New("invalid content type")
	}
}
