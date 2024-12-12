package web

import (
	"fmt"
	"net/http"
	"strconv"
)

type ChannelCtrl struct{}

func (c *ChannelCtrl) GetBasePath() string {
	return "/channels/"
}

func (c *ChannelCtrl) GetRoutes() []Route {
	return []Route{
		{Methods: "GET", Pattern: "/", Handler: c.index},
		{Methods: "GET", Pattern: "/test/:int/test", Handler: c.notIndex},
	}
}

func (c *ChannelCtrl) index(w http.ResponseWriter, r *http.Request, rte MatchedRoute) {
	tokens := tokenize(r.URL.Path)
	fmt.Fprintln(w, tokens)
	fmt.Fprintln(w, len(tokens))
	fmt.Fprintln(w, "Bonjoure")
}

func (c *ChannelCtrl) notIndex(w http.ResponseWriter, r *http.Request, rte MatchedRoute) {
	fmt.Fprintln(w, "Hello")
	id, _ := strconv.Atoi(rte.Params[0])
	fmt.Fprintln(w, id)
}
