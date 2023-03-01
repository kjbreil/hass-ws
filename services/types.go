package services

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Service interface {
	SetID(id *int)
	JSON() string
}
type Flash string

const (
	FlashLong  Flash = "long"
	FlashShort Flash = "short"
)

type Level string

const (
	LevelDebug    Level = "debug"
	LevelInfo     Level = "info"
	LevelWarning  Level = "warning"
	LevelError    Level = "error"
	LevelFatal    Level = "fatal"
	LevelCritical Level = "critical"
)

type Qos string

const (
	QosNull Qos = "null"
)

type Format string

const (
	FormatNull Format = "null"
)

type ColorName string

const (
	ColorNameNull ColorName = "null"
)

type Mode string

const (
	ModeDark  Mode = "dark"
	ModeLight Mode = "light"
)

type HvacMode string

const (
	HvacModeHeatCool HvacMode = "heat_cool"
	HvacModeHeat     HvacMode = "heat"
	HvacModeOff      HvacMode = "off"
	HvacModeAuto     HvacMode = "auto"
	HvacModeCool     HvacMode = "cool"
	HvacModeDry      HvacMode = "dry"
	HvacModeFanOnly  HvacMode = "fan_only"
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
