package main

var ServiceNames = append([]string{
	"automation",
	"backup",
	"counter",
	"frontend",
	"group",
	"homeassistant",
	"input_boolean",
	"input_button",
	"input_datetime",
	"input_number",
	"input_select",
	"input_text",
	"logbook",
	"logger",
	"min_max",
	"mqtt",
	"notify",
	"persistent_notification",
	"person",
	"recorder",
	"scene",
	"schedule",
	"script",
	"shell_command",
	"system_log",
	"template",
	"timer",
	"tts",
	"wake_on_lan",
	"zone",
}, DeviceNames...)

type Service struct {
	Name string
}

func ServicesInit() (retval []Service) {
	for _, name := range ServiceNames {
		s := Service{
			Name: name,
		}
		retval = append(retval, s)
	}
	return retval
}
