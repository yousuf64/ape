package main

import (
	"github.com/yousuf64/shift"
	"net/http"
)

func groupV2(g *shift.Group) {
	g.GET("/abc", func(w http.ResponseWriter, r *http.Request, route shift.Route) error {
		_, err := w.Write([]byte("v2.go file: abc"))
		return err
	})
	g.GET("/xyz", func(w http.ResponseWriter, r *http.Request, route shift.Route) error {
		_, err := w.Write([]byte("v2.go file: xyz"))
		return err
	})
}
