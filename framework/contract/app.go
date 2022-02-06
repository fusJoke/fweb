package contract


const AppKey = "hade:app"

type App interface {
	Version() string
	BaseFolder() string
	ConfigFolder() string
	LogFolder() string
	Provider() string
	MiddlewareFolder() string
	CommandFolder() string
	RuntimeFolder() string
	TestFolder() string
}