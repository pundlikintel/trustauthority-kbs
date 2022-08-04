/*
 * Copyright (C) 2022 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */

package http

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"intel/amber/kbs/v1/clients/constant"
	"intel/amber/kbs/v1/model"
	"intel/amber/kbs/v1/service"
	"net/http"
)

func setCreateAuthTokenHandler(svc service.Service, router *mux.Router, options []httpTransport.ServerOption, jwtAuth *model.JwtAuthz) error {

	createAuthTokenHandler := httpTransport.NewServer(
		makeCreateAuthTokenHttpEndpoint(svc, jwtAuth),
		decodeCreateAuthTokenHttpRequest,
		encodeCreateAuthTokenHttpResponse,
		options...,
	)

	router.Handle("/token", createAuthTokenHandler).Methods(http.MethodPost)

	return nil
}

func makeCreateAuthTokenHttpEndpoint(svc service.Service, jwtAuth *model.JwtAuthz) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.AuthTokenRequest)
		return svc.CreateAuthToken(ctx, req, jwtAuth)
	}
}

func decodeCreateAuthTokenHttpRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req model.AuthTokenRequest

	if r.Header.Get(constant.HTTPHeaderKeyContentType) != constant.HTTPHeaderValueApplicationJson {
		log.Error(ErrInvalidAcceptHeader.Error())
		return nil, ErrInvalidContentTypeHeader
	}
	if r.Header.Get(constant.HTTPHeaderKeyAccept) != constant.HTTPHeaderValueApplicationJwt {
		log.Error(ErrInvalidAcceptHeader.Error())
		return nil, ErrInvalidAcceptHeader
	}
	if r.ContentLength == 0 {
		log.Error(ErrEmptyRequestBody.Error())
		return nil, ErrEmptyRequestBody
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		log.WithError(err).Error(ErrJsonDecodeFailed.Error())
		return nil, ErrJsonDecodeFailed
	}

	return req, nil
}

func encodeCreateAuthTokenHttpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(string)
	w.Write([]byte(resp))
	return nil
}
