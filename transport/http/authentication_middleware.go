/*
 * Copyright(C) 2023 Intel Corporation. All Rights Reserved.
 */

package http

import (
	"github.com/shaj13/go-guardian/v2/auth"
	log "github.com/sirupsen/logrus"
	"intel/amber/kbs/v1/model"

	"net/http"
)

func authMiddleware(next http.Handler, authz *model.JwtAuthz) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strategy := authz.AuthZStrategy
		user, err := strategy.Authenticate(r.Context(), r)
		if err != nil {
			log.WithError(err).Error("Request unauthorized")
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		r = auth.RequestWithUser(user, r)
		next.ServeHTTP(w, r)
	})
}
