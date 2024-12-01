package main

type API struct{}

// NewAPI creates a new API instance
func NewAPI() *API {
	return &API{}
}

// Echo returns a test message
func (a *API) Echo(message string) string {
	return "Echo: " + message
}
