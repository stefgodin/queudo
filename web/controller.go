package web

import (
	"net/http"
	"strings"
)

type Controller interface {
	GetBasePath() string
	GetRoutes() []Route
}

type RouteHandler func(http.ResponseWriter, *http.Request, MatchedRoute)

type Route struct {
	Methods string
	Pattern string
	Handler RouteHandler
}

type MatchedRoute struct {
	Route  Route
	Params []string
}

type Router struct {
	URLMatcher *URLMatcher
}

func (rtr *Router) RegisterController(ctrl Controller) {
	routes := ctrl.GetRoutes()
	http.HandleFunc(ctrl.GetBasePath(), func(w http.ResponseWriter, r *http.Request) {
		for _, route := range routes {
			if !strings.Contains(route.Methods, r.Method) {
				continue
			}

			pattern := ctrl.GetBasePath() + "/" + route.Pattern
			params, err := rtr.URLMatcher.match(r.URL.Path, pattern)

			if err != nil {
				continue
			}

			route.Handler(w, r, MatchedRoute{Route: route, Params: params})
			return
		}

		http.NotFound(w, r)
	})
}

var DefaultRouter = &Router{
	URLMatcher: DefaultURLMatcher,
}
