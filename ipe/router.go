// Copyright 2014, 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ipe

import (
	"net/http"

	"github.com/gorilla/mux"
)

type router struct {
	ctx    *applicationContext
	mux    *mux.Router
	routes map[string]contextHandler
}

func newRouter(ctx *applicationContext) *router {
	return &router{
		ctx: ctx,
		mux: mux.NewRouter().StrictSlash(true),
	}
}

func (a *router) GET(path string, handler contextHandler) {
	a.Handle("GET", path, handler)
}

func (a *router) POST(path string, handler contextHandler) {
	a.Handle("POST", path, handler)
}

func (a *router) Handle(method, path string, handler contextHandler) {
	a.mux.Methods(method).Path(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWithContext(a.ctx, params(mux.Vars(r)), w, r)
	})
}

func (a router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
