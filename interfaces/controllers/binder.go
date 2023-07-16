package controllers

type RouteBinder interface {
	Bind()
}

func BindRoutes(routes ...RouteBinder) {
	for _, route := range routes {
		route.Bind()
	}
}
