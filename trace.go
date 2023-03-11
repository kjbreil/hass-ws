package hass_ws

type (
	Logger interface {
		Println(v ...interface{})
		Printf(format string, v ...interface{})
	}
	NOOPLogger struct{}
)

func (NOOPLogger) Println(v ...interface{})               {}
func (NOOPLogger) Printf(format string, v ...interface{}) {}

var (
	INFO     Logger = NOOPLogger{}
	ERROR    Logger = NOOPLogger{}
	CRITICAL Logger = NOOPLogger{}
	WARN     Logger = NOOPLogger{}
	DEBUG    Logger = NOOPLogger{}
)
