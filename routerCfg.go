package routeLoader

//the config struct
type RouteCfg struct {
	Schemes          []string
	Host             []string
	RateLimit        bool
	PrefixPath       string
	PrefixStaticPath string
}

//route description
type RouteDesc struct {
	URI      string
	Verbs    []string
	Handlers []RouteHandler
	//defualt value is the value of the PrefixStaticPath
	// in RouteCfg
	PrefixPath string
}

//route hqandler description
type RouteHandler struct {
	//class name
	InstanceName string
	//handler function name
	Handler []string
}
