package services

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Service interface {
	SetID(id *int)
	JSON() string
}
type Level string

const (
	LevelError    Level = "error"
	LevelFatal    Level = "fatal"
	LevelCritical Level = "critical"
	LevelDebug    Level = "debug"
	LevelInfo     Level = "info"
	LevelWarning  Level = "warning"
)

type Format string

const (
	FormatNull Format = "null"
)

type ColorName string

const (
	ColorNameNull ColorName = "null"
)

type Flash string

const (
	FlashLong  Flash = "long"
	FlashShort Flash = "short"
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

type HvacMode string

const (
	HvacModeAuto     HvacMode = "auto"
	HvacModeCool     HvacMode = "cool"
	HvacModeDry      HvacMode = "dry"
	HvacModeFanOnly  HvacMode = "fan_only"
	HvacModeHeatCool HvacMode = "heat_cool"
	HvacModeHeat     HvacMode = "heat"
	HvacModeOff      HvacMode = "off"
)

type Direction string

const (
	DirectionForward Direction = "forward"
	DirectionReverse Direction = "reverse"
)

type CommandType string

const (
	CommandTypeNull CommandType = "null"
)
