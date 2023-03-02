package services

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////

type Service interface {
	SetID(id *int)
	JSON() string
}
type ColorName string

const (
	ColorNameNull ColorName = "null"
)

type CommandType string

const (
	CommandTypeNull CommandType = "null"
)

type Direction string

const (
	DirectionForward Direction = "forward"
	DirectionReverse Direction = "reverse"
)

type Enqueue string

const (
	EnqueueAdd     Enqueue = "add"
	EnqueueNext    Enqueue = "next"
	EnqueuePlay    Enqueue = "play"
	EnqueueReplace Enqueue = "replace"
)

type Flash string

const (
	FlashLong  Flash = "long"
	FlashShort Flash = "short"
)

type Format string

const (
	FormatNull Format = "null"
)

type HvacMode string

const (
	HvacModeAuto     HvacMode = "auto"
	HvacModeCool     HvacMode = "cool"
	HvacModeDry      HvacMode = "dry"
	HvacModeFanOnly  HvacMode = "fan_only"
	HvacModeHeat     HvacMode = "heat"
	HvacModeHeatCool HvacMode = "heat_cool"
	HvacModeOff      HvacMode = "off"
)

type Level string

const (
	LevelCritical Level = "critical"
	LevelDebug    Level = "debug"
	LevelError    Level = "error"
	LevelFatal    Level = "fatal"
	LevelInfo     Level = "info"
	LevelWarning  Level = "warning"
)

type Mode string

const (
	ModeDark  Mode = "dark"
	ModeLight Mode = "light"
)

type Qos string

const (
	QosNull Qos = "null"
)

type Repeat string

const (
	RepeatAll Repeat = "all"
	RepeatOff Repeat = "off"
	RepeatOne Repeat = "one"
)

type RepeatType string

const (
	RepeatTypeNull RepeatType = "null"
)

type SignatureScheme string

const (
	SignatureSchemeNull SignatureScheme = "null"
)

type Transport string

const (
	TransportNull Transport = "null"
)
