package channel

var handlers = make(map[string]*Handler)

// use to load the current handler
func Load(host string) *Handler {
	handler := getHandler(host)
	if handler == nil {
		handler = createHandler(host)
		addHandler(handler)
	}
	return handler
}

func getHandler(host string) *Handler {
	return handlers[host]
}

func addHandler(handler *Handler) {
	handlers[handler.Host] = handler
}
