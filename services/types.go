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
	ColorNameAliceblue            ColorName = "aliceblue"
	ColorNameAntiquewhite         ColorName = "antiquewhite"
	ColorNameAqua                 ColorName = "aqua"
	ColorNameAquamarine           ColorName = "aquamarine"
	ColorNameAzure                ColorName = "azure"
	ColorNameBeige                ColorName = "beige"
	ColorNameBisque               ColorName = "bisque"
	ColorNameBlanchedalmond       ColorName = "blanchedalmond"
	ColorNameBlue                 ColorName = "blue"
	ColorNameBlueviolet           ColorName = "blueviolet"
	ColorNameBrown                ColorName = "brown"
	ColorNameBurlywood            ColorName = "burlywood"
	ColorNameCadetblue            ColorName = "cadetblue"
	ColorNameChartreuse           ColorName = "chartreuse"
	ColorNameChocolate            ColorName = "chocolate"
	ColorNameCoral                ColorName = "coral"
	ColorNameCornflowerblue       ColorName = "cornflowerblue"
	ColorNameCornsilk             ColorName = "cornsilk"
	ColorNameCrimson              ColorName = "crimson"
	ColorNameCyan                 ColorName = "cyan"
	ColorNameDarkblue             ColorName = "darkblue"
	ColorNameDarkcyan             ColorName = "darkcyan"
	ColorNameDarkgoldenrod        ColorName = "darkgoldenrod"
	ColorNameDarkgray             ColorName = "darkgray"
	ColorNameDarkgreen            ColorName = "darkgreen"
	ColorNameDarkgrey             ColorName = "darkgrey"
	ColorNameDarkkhaki            ColorName = "darkkhaki"
	ColorNameDarkmagenta          ColorName = "darkmagenta"
	ColorNameDarkolivegreen       ColorName = "darkolivegreen"
	ColorNameDarkorange           ColorName = "darkorange"
	ColorNameDarkorchid           ColorName = "darkorchid"
	ColorNameDarkred              ColorName = "darkred"
	ColorNameDarksalmon           ColorName = "darksalmon"
	ColorNameDarkseagreen         ColorName = "darkseagreen"
	ColorNameDarkslateblue        ColorName = "darkslateblue"
	ColorNameDarkslategray        ColorName = "darkslategray"
	ColorNameDarkslategrey        ColorName = "darkslategrey"
	ColorNameDarkturquoise        ColorName = "darkturquoise"
	ColorNameDarkviolet           ColorName = "darkviolet"
	ColorNameDeeppink             ColorName = "deeppink"
	ColorNameDeepskyblue          ColorName = "deepskyblue"
	ColorNameDimgray              ColorName = "dimgray"
	ColorNameDimgrey              ColorName = "dimgrey"
	ColorNameDodgerblue           ColorName = "dodgerblue"
	ColorNameFirebrick            ColorName = "firebrick"
	ColorNameFloralwhite          ColorName = "floralwhite"
	ColorNameForestgreen          ColorName = "forestgreen"
	ColorNameFuchsia              ColorName = "fuchsia"
	ColorNameGainsboro            ColorName = "gainsboro"
	ColorNameGhostwhite           ColorName = "ghostwhite"
	ColorNameGold                 ColorName = "gold"
	ColorNameGoldenrod            ColorName = "goldenrod"
	ColorNameGray                 ColorName = "gray"
	ColorNameGreen                ColorName = "green"
	ColorNameGreenyellow          ColorName = "greenyellow"
	ColorNameGrey                 ColorName = "grey"
	ColorNameHomeassistant        ColorName = "homeassistant"
	ColorNameHoneydew             ColorName = "honeydew"
	ColorNameHotpink              ColorName = "hotpink"
	ColorNameIndianred            ColorName = "indianred"
	ColorNameIndigo               ColorName = "indigo"
	ColorNameIvory                ColorName = "ivory"
	ColorNameKhaki                ColorName = "khaki"
	ColorNameLavender             ColorName = "lavender"
	ColorNameLavenderblush        ColorName = "lavenderblush"
	ColorNameLawngreen            ColorName = "lawngreen"
	ColorNameLemonchiffon         ColorName = "lemonchiffon"
	ColorNameLightblue            ColorName = "lightblue"
	ColorNameLightcoral           ColorName = "lightcoral"
	ColorNameLightcyan            ColorName = "lightcyan"
	ColorNameLightgoldenrodyellow ColorName = "lightgoldenrodyellow"
	ColorNameLightgray            ColorName = "lightgray"
	ColorNameLightgreen           ColorName = "lightgreen"
	ColorNameLightgrey            ColorName = "lightgrey"
	ColorNameLightpink            ColorName = "lightpink"
	ColorNameLightsalmon          ColorName = "lightsalmon"
	ColorNameLightseagreen        ColorName = "lightseagreen"
	ColorNameLightskyblue         ColorName = "lightskyblue"
	ColorNameLightslategray       ColorName = "lightslategray"
	ColorNameLightslategrey       ColorName = "lightslategrey"
	ColorNameLightsteelblue       ColorName = "lightsteelblue"
	ColorNameLightyellow          ColorName = "lightyellow"
	ColorNameLime                 ColorName = "lime"
	ColorNameLimegreen            ColorName = "limegreen"
	ColorNameLinen                ColorName = "linen"
	ColorNameMagenta              ColorName = "magenta"
	ColorNameMaroon               ColorName = "maroon"
	ColorNameMediumaquamarine     ColorName = "mediumaquamarine"
	ColorNameMediumblue           ColorName = "mediumblue"
	ColorNameMediumorchid         ColorName = "mediumorchid"
	ColorNameMediumpurple         ColorName = "mediumpurple"
	ColorNameMediumseagreen       ColorName = "mediumseagreen"
	ColorNameMediumslateblue      ColorName = "mediumslateblue"
	ColorNameMediumspringgreen    ColorName = "mediumspringgreen"
	ColorNameMediumturquoise      ColorName = "mediumturquoise"
	ColorNameMediumvioletred      ColorName = "mediumvioletred"
	ColorNameMidnightblue         ColorName = "midnightblue"
	ColorNameMintcream            ColorName = "mintcream"
	ColorNameMistyrose            ColorName = "mistyrose"
	ColorNameMoccasin             ColorName = "moccasin"
	ColorNameNavajowhite          ColorName = "navajowhite"
	ColorNameNavy                 ColorName = "navy"
	ColorNameNavyblue             ColorName = "navyblue"
	ColorNameOldlace              ColorName = "oldlace"
	ColorNameOlive                ColorName = "olive"
	ColorNameOlivedrab            ColorName = "olivedrab"
	ColorNameOrange               ColorName = "orange"
	ColorNameOrangered            ColorName = "orangered"
	ColorNameOrchid               ColorName = "orchid"
	ColorNamePalegoldenrod        ColorName = "palegoldenrod"
	ColorNamePalegreen            ColorName = "palegreen"
	ColorNamePaleturquoise        ColorName = "paleturquoise"
	ColorNamePalevioletred        ColorName = "palevioletred"
	ColorNamePapayawhip           ColorName = "papayawhip"
	ColorNamePeachpuff            ColorName = "peachpuff"
	ColorNamePeru                 ColorName = "peru"
	ColorNamePink                 ColorName = "pink"
	ColorNamePlum                 ColorName = "plum"
	ColorNamePowderblue           ColorName = "powderblue"
	ColorNamePurple               ColorName = "purple"
	ColorNameRed                  ColorName = "red"
	ColorNameRosybrown            ColorName = "rosybrown"
	ColorNameRoyalblue            ColorName = "royalblue"
	ColorNameSaddlebrown          ColorName = "saddlebrown"
	ColorNameSalmon               ColorName = "salmon"
	ColorNameSandybrown           ColorName = "sandybrown"
	ColorNameSeagreen             ColorName = "seagreen"
	ColorNameSeashell             ColorName = "seashell"
	ColorNameSienna               ColorName = "sienna"
	ColorNameSilver               ColorName = "silver"
	ColorNameSkyblue              ColorName = "skyblue"
	ColorNameSlateblue            ColorName = "slateblue"
	ColorNameSlategray            ColorName = "slategray"
	ColorNameSlategrey            ColorName = "slategrey"
	ColorNameSnow                 ColorName = "snow"
	ColorNameSpringgreen          ColorName = "springgreen"
	ColorNameSteelblue            ColorName = "steelblue"
	ColorNameTan                  ColorName = "tan"
	ColorNameTeal                 ColorName = "teal"
	ColorNameThistle              ColorName = "thistle"
	ColorNameTomato               ColorName = "tomato"
	ColorNameTurquoise            ColorName = "turquoise"
	ColorNameViolet               ColorName = "violet"
	ColorNameWheat                ColorName = "wheat"
	ColorNameWhite                ColorName = "white"
	ColorNameWhitesmoke           ColorName = "whitesmoke"
	ColorNameYellow               ColorName = "yellow"
	ColorNameYellowgreen          ColorName = "yellowgreen"
)

type CommandType string

const (
	CommandTypeIr CommandType = "ir"
	CommandTypeRf CommandType = "rf"
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
	FormatHls Format = "hls"
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
	Qos0 Qos = "0"
	Qos1 Qos = "1"
	Qos2 Qos = "2"
)

type Repeat string

const (
	RepeatAll Repeat = "all"
	RepeatOff Repeat = "off"
	RepeatOne Repeat = "one"
)

type RepeatType string

const (
	RepeatTypePause  RepeatType = "pause"
	RepeatTypeRepeat RepeatType = "repeat"
	RepeatTypeSingle RepeatType = "single"
)

type SignatureScheme string

const (
	SignatureSchemeHmacSha256 SignatureScheme = "hmac-sha256"
)

type Transport string

const (
	TransportTcp Transport = "tcp"
	TransportUdp Transport = "udp"
)
