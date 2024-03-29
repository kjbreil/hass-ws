// Code generated by jsonenums -type=ColorName; DO NOT EDIT.

package services

import (
	"fmt"
	"github.com/goccy/go-json"
)

var (
	_ColorNameNameToValue = map[string]ColorName{
		"ColorNamealiceblue":            ColorNamealiceblue,
		"ColorNameantiquewhite":         ColorNameantiquewhite,
		"ColorNameaqua":                 ColorNameaqua,
		"ColorNameaquamarine":           ColorNameaquamarine,
		"ColorNameazure":                ColorNameazure,
		"ColorNamebeige":                ColorNamebeige,
		"ColorNamebisque":               ColorNamebisque,
		"ColorNameblanchedalmond":       ColorNameblanchedalmond,
		"ColorNameblue":                 ColorNameblue,
		"ColorNameblueviolet":           ColorNameblueviolet,
		"ColorNamebrown":                ColorNamebrown,
		"ColorNameburlywood":            ColorNameburlywood,
		"ColorNamecadetblue":            ColorNamecadetblue,
		"ColorNamechartreuse":           ColorNamechartreuse,
		"ColorNamechocolate":            ColorNamechocolate,
		"ColorNamecoral":                ColorNamecoral,
		"ColorNamecornflowerblue":       ColorNamecornflowerblue,
		"ColorNamecornsilk":             ColorNamecornsilk,
		"ColorNamecrimson":              ColorNamecrimson,
		"ColorNamecyan":                 ColorNamecyan,
		"ColorNamedarkblue":             ColorNamedarkblue,
		"ColorNamedarkcyan":             ColorNamedarkcyan,
		"ColorNamedarkgoldenrod":        ColorNamedarkgoldenrod,
		"ColorNamedarkgray":             ColorNamedarkgray,
		"ColorNamedarkgreen":            ColorNamedarkgreen,
		"ColorNamedarkgrey":             ColorNamedarkgrey,
		"ColorNamedarkkhaki":            ColorNamedarkkhaki,
		"ColorNamedarkmagenta":          ColorNamedarkmagenta,
		"ColorNamedarkolivegreen":       ColorNamedarkolivegreen,
		"ColorNamedarkorange":           ColorNamedarkorange,
		"ColorNamedarkorchid":           ColorNamedarkorchid,
		"ColorNamedarkred":              ColorNamedarkred,
		"ColorNamedarksalmon":           ColorNamedarksalmon,
		"ColorNamedarkseagreen":         ColorNamedarkseagreen,
		"ColorNamedarkslateblue":        ColorNamedarkslateblue,
		"ColorNamedarkslategray":        ColorNamedarkslategray,
		"ColorNamedarkslategrey":        ColorNamedarkslategrey,
		"ColorNamedarkturquoise":        ColorNamedarkturquoise,
		"ColorNamedarkviolet":           ColorNamedarkviolet,
		"ColorNamedeeppink":             ColorNamedeeppink,
		"ColorNamedeepskyblue":          ColorNamedeepskyblue,
		"ColorNamedimgray":              ColorNamedimgray,
		"ColorNamedimgrey":              ColorNamedimgrey,
		"ColorNamedodgerblue":           ColorNamedodgerblue,
		"ColorNamefirebrick":            ColorNamefirebrick,
		"ColorNamefloralwhite":          ColorNamefloralwhite,
		"ColorNameforestgreen":          ColorNameforestgreen,
		"ColorNamefuchsia":              ColorNamefuchsia,
		"ColorNamegainsboro":            ColorNamegainsboro,
		"ColorNameghostwhite":           ColorNameghostwhite,
		"ColorNamegold":                 ColorNamegold,
		"ColorNamegoldenrod":            ColorNamegoldenrod,
		"ColorNamegray":                 ColorNamegray,
		"ColorNamegreen":                ColorNamegreen,
		"ColorNamegreenyellow":          ColorNamegreenyellow,
		"ColorNamegrey":                 ColorNamegrey,
		"ColorNamehomeassistant":        ColorNamehomeassistant,
		"ColorNamehoneydew":             ColorNamehoneydew,
		"ColorNamehotpink":              ColorNamehotpink,
		"ColorNameindianred":            ColorNameindianred,
		"ColorNameindigo":               ColorNameindigo,
		"ColorNameivory":                ColorNameivory,
		"ColorNamekhaki":                ColorNamekhaki,
		"ColorNamelavender":             ColorNamelavender,
		"ColorNamelavenderblush":        ColorNamelavenderblush,
		"ColorNamelawngreen":            ColorNamelawngreen,
		"ColorNamelemonchiffon":         ColorNamelemonchiffon,
		"ColorNamelightblue":            ColorNamelightblue,
		"ColorNamelightcoral":           ColorNamelightcoral,
		"ColorNamelightcyan":            ColorNamelightcyan,
		"ColorNamelightgoldenrodyellow": ColorNamelightgoldenrodyellow,
		"ColorNamelightgray":            ColorNamelightgray,
		"ColorNamelightgreen":           ColorNamelightgreen,
		"ColorNamelightgrey":            ColorNamelightgrey,
		"ColorNamelightpink":            ColorNamelightpink,
		"ColorNamelightsalmon":          ColorNamelightsalmon,
		"ColorNamelightseagreen":        ColorNamelightseagreen,
		"ColorNamelightskyblue":         ColorNamelightskyblue,
		"ColorNamelightslategray":       ColorNamelightslategray,
		"ColorNamelightslategrey":       ColorNamelightslategrey,
		"ColorNamelightsteelblue":       ColorNamelightsteelblue,
		"ColorNamelightyellow":          ColorNamelightyellow,
		"ColorNamelime":                 ColorNamelime,
		"ColorNamelimegreen":            ColorNamelimegreen,
		"ColorNamelinen":                ColorNamelinen,
		"ColorNamemagenta":              ColorNamemagenta,
		"ColorNamemaroon":               ColorNamemaroon,
		"ColorNamemediumaquamarine":     ColorNamemediumaquamarine,
		"ColorNamemediumblue":           ColorNamemediumblue,
		"ColorNamemediumorchid":         ColorNamemediumorchid,
		"ColorNamemediumpurple":         ColorNamemediumpurple,
		"ColorNamemediumseagreen":       ColorNamemediumseagreen,
		"ColorNamemediumslateblue":      ColorNamemediumslateblue,
		"ColorNamemediumspringgreen":    ColorNamemediumspringgreen,
		"ColorNamemediumturquoise":      ColorNamemediumturquoise,
		"ColorNamemediumvioletred":      ColorNamemediumvioletred,
		"ColorNamemidnightblue":         ColorNamemidnightblue,
		"ColorNamemintcream":            ColorNamemintcream,
		"ColorNamemistyrose":            ColorNamemistyrose,
		"ColorNamemoccasin":             ColorNamemoccasin,
		"ColorNamenavajowhite":          ColorNamenavajowhite,
		"ColorNamenavy":                 ColorNamenavy,
		"ColorNamenavyblue":             ColorNamenavyblue,
		"ColorNameoldlace":              ColorNameoldlace,
		"ColorNameolive":                ColorNameolive,
		"ColorNameolivedrab":            ColorNameolivedrab,
		"ColorNameorange":               ColorNameorange,
		"ColorNameorangered":            ColorNameorangered,
		"ColorNameorchid":               ColorNameorchid,
		"ColorNamepalegoldenrod":        ColorNamepalegoldenrod,
		"ColorNamepalegreen":            ColorNamepalegreen,
		"ColorNamepaleturquoise":        ColorNamepaleturquoise,
		"ColorNamepalevioletred":        ColorNamepalevioletred,
		"ColorNamepapayawhip":           ColorNamepapayawhip,
		"ColorNamepeachpuff":            ColorNamepeachpuff,
		"ColorNameperu":                 ColorNameperu,
		"ColorNamepink":                 ColorNamepink,
		"ColorNameplum":                 ColorNameplum,
		"ColorNamepowderblue":           ColorNamepowderblue,
		"ColorNamepurple":               ColorNamepurple,
		"ColorNamered":                  ColorNamered,
		"ColorNamerosybrown":            ColorNamerosybrown,
		"ColorNameroyalblue":            ColorNameroyalblue,
		"ColorNamesaddlebrown":          ColorNamesaddlebrown,
		"ColorNamesalmon":               ColorNamesalmon,
		"ColorNamesandybrown":           ColorNamesandybrown,
		"ColorNameseagreen":             ColorNameseagreen,
		"ColorNameseashell":             ColorNameseashell,
		"ColorNamesienna":               ColorNamesienna,
		"ColorNamesilver":               ColorNamesilver,
		"ColorNameskyblue":              ColorNameskyblue,
		"ColorNameslateblue":            ColorNameslateblue,
		"ColorNameslategray":            ColorNameslategray,
		"ColorNameslategrey":            ColorNameslategrey,
		"ColorNamesnow":                 ColorNamesnow,
		"ColorNamespringgreen":          ColorNamespringgreen,
		"ColorNamesteelblue":            ColorNamesteelblue,
		"ColorNametan":                  ColorNametan,
		"ColorNameteal":                 ColorNameteal,
		"ColorNamethistle":              ColorNamethistle,
		"ColorNametomato":               ColorNametomato,
		"ColorNameturquoise":            ColorNameturquoise,
		"ColorNameviolet":               ColorNameviolet,
		"ColorNamewheat":                ColorNamewheat,
		"ColorNamewhite":                ColorNamewhite,
		"ColorNamewhitesmoke":           ColorNamewhitesmoke,
		"ColorNameyellow":               ColorNameyellow,
		"ColorNameyellowgreen":          ColorNameyellowgreen,
	}

	_ColorNameValueToName = map[ColorName]string{
		ColorNamealiceblue:            "ColorNamealiceblue",
		ColorNameantiquewhite:         "ColorNameantiquewhite",
		ColorNameaqua:                 "ColorNameaqua",
		ColorNameaquamarine:           "ColorNameaquamarine",
		ColorNameazure:                "ColorNameazure",
		ColorNamebeige:                "ColorNamebeige",
		ColorNamebisque:               "ColorNamebisque",
		ColorNameblanchedalmond:       "ColorNameblanchedalmond",
		ColorNameblue:                 "ColorNameblue",
		ColorNameblueviolet:           "ColorNameblueviolet",
		ColorNamebrown:                "ColorNamebrown",
		ColorNameburlywood:            "ColorNameburlywood",
		ColorNamecadetblue:            "ColorNamecadetblue",
		ColorNamechartreuse:           "ColorNamechartreuse",
		ColorNamechocolate:            "ColorNamechocolate",
		ColorNamecoral:                "ColorNamecoral",
		ColorNamecornflowerblue:       "ColorNamecornflowerblue",
		ColorNamecornsilk:             "ColorNamecornsilk",
		ColorNamecrimson:              "ColorNamecrimson",
		ColorNamecyan:                 "ColorNamecyan",
		ColorNamedarkblue:             "ColorNamedarkblue",
		ColorNamedarkcyan:             "ColorNamedarkcyan",
		ColorNamedarkgoldenrod:        "ColorNamedarkgoldenrod",
		ColorNamedarkgray:             "ColorNamedarkgray",
		ColorNamedarkgreen:            "ColorNamedarkgreen",
		ColorNamedarkgrey:             "ColorNamedarkgrey",
		ColorNamedarkkhaki:            "ColorNamedarkkhaki",
		ColorNamedarkmagenta:          "ColorNamedarkmagenta",
		ColorNamedarkolivegreen:       "ColorNamedarkolivegreen",
		ColorNamedarkorange:           "ColorNamedarkorange",
		ColorNamedarkorchid:           "ColorNamedarkorchid",
		ColorNamedarkred:              "ColorNamedarkred",
		ColorNamedarksalmon:           "ColorNamedarksalmon",
		ColorNamedarkseagreen:         "ColorNamedarkseagreen",
		ColorNamedarkslateblue:        "ColorNamedarkslateblue",
		ColorNamedarkslategray:        "ColorNamedarkslategray",
		ColorNamedarkslategrey:        "ColorNamedarkslategrey",
		ColorNamedarkturquoise:        "ColorNamedarkturquoise",
		ColorNamedarkviolet:           "ColorNamedarkviolet",
		ColorNamedeeppink:             "ColorNamedeeppink",
		ColorNamedeepskyblue:          "ColorNamedeepskyblue",
		ColorNamedimgray:              "ColorNamedimgray",
		ColorNamedimgrey:              "ColorNamedimgrey",
		ColorNamedodgerblue:           "ColorNamedodgerblue",
		ColorNamefirebrick:            "ColorNamefirebrick",
		ColorNamefloralwhite:          "ColorNamefloralwhite",
		ColorNameforestgreen:          "ColorNameforestgreen",
		ColorNamefuchsia:              "ColorNamefuchsia",
		ColorNamegainsboro:            "ColorNamegainsboro",
		ColorNameghostwhite:           "ColorNameghostwhite",
		ColorNamegold:                 "ColorNamegold",
		ColorNamegoldenrod:            "ColorNamegoldenrod",
		ColorNamegray:                 "ColorNamegray",
		ColorNamegreen:                "ColorNamegreen",
		ColorNamegreenyellow:          "ColorNamegreenyellow",
		ColorNamegrey:                 "ColorNamegrey",
		ColorNamehomeassistant:        "ColorNamehomeassistant",
		ColorNamehoneydew:             "ColorNamehoneydew",
		ColorNamehotpink:              "ColorNamehotpink",
		ColorNameindianred:            "ColorNameindianred",
		ColorNameindigo:               "ColorNameindigo",
		ColorNameivory:                "ColorNameivory",
		ColorNamekhaki:                "ColorNamekhaki",
		ColorNamelavender:             "ColorNamelavender",
		ColorNamelavenderblush:        "ColorNamelavenderblush",
		ColorNamelawngreen:            "ColorNamelawngreen",
		ColorNamelemonchiffon:         "ColorNamelemonchiffon",
		ColorNamelightblue:            "ColorNamelightblue",
		ColorNamelightcoral:           "ColorNamelightcoral",
		ColorNamelightcyan:            "ColorNamelightcyan",
		ColorNamelightgoldenrodyellow: "ColorNamelightgoldenrodyellow",
		ColorNamelightgray:            "ColorNamelightgray",
		ColorNamelightgreen:           "ColorNamelightgreen",
		ColorNamelightgrey:            "ColorNamelightgrey",
		ColorNamelightpink:            "ColorNamelightpink",
		ColorNamelightsalmon:          "ColorNamelightsalmon",
		ColorNamelightseagreen:        "ColorNamelightseagreen",
		ColorNamelightskyblue:         "ColorNamelightskyblue",
		ColorNamelightslategray:       "ColorNamelightslategray",
		ColorNamelightslategrey:       "ColorNamelightslategrey",
		ColorNamelightsteelblue:       "ColorNamelightsteelblue",
		ColorNamelightyellow:          "ColorNamelightyellow",
		ColorNamelime:                 "ColorNamelime",
		ColorNamelimegreen:            "ColorNamelimegreen",
		ColorNamelinen:                "ColorNamelinen",
		ColorNamemagenta:              "ColorNamemagenta",
		ColorNamemaroon:               "ColorNamemaroon",
		ColorNamemediumaquamarine:     "ColorNamemediumaquamarine",
		ColorNamemediumblue:           "ColorNamemediumblue",
		ColorNamemediumorchid:         "ColorNamemediumorchid",
		ColorNamemediumpurple:         "ColorNamemediumpurple",
		ColorNamemediumseagreen:       "ColorNamemediumseagreen",
		ColorNamemediumslateblue:      "ColorNamemediumslateblue",
		ColorNamemediumspringgreen:    "ColorNamemediumspringgreen",
		ColorNamemediumturquoise:      "ColorNamemediumturquoise",
		ColorNamemediumvioletred:      "ColorNamemediumvioletred",
		ColorNamemidnightblue:         "ColorNamemidnightblue",
		ColorNamemintcream:            "ColorNamemintcream",
		ColorNamemistyrose:            "ColorNamemistyrose",
		ColorNamemoccasin:             "ColorNamemoccasin",
		ColorNamenavajowhite:          "ColorNamenavajowhite",
		ColorNamenavy:                 "ColorNamenavy",
		ColorNamenavyblue:             "ColorNamenavyblue",
		ColorNameoldlace:              "ColorNameoldlace",
		ColorNameolive:                "ColorNameolive",
		ColorNameolivedrab:            "ColorNameolivedrab",
		ColorNameorange:               "ColorNameorange",
		ColorNameorangered:            "ColorNameorangered",
		ColorNameorchid:               "ColorNameorchid",
		ColorNamepalegoldenrod:        "ColorNamepalegoldenrod",
		ColorNamepalegreen:            "ColorNamepalegreen",
		ColorNamepaleturquoise:        "ColorNamepaleturquoise",
		ColorNamepalevioletred:        "ColorNamepalevioletred",
		ColorNamepapayawhip:           "ColorNamepapayawhip",
		ColorNamepeachpuff:            "ColorNamepeachpuff",
		ColorNameperu:                 "ColorNameperu",
		ColorNamepink:                 "ColorNamepink",
		ColorNameplum:                 "ColorNameplum",
		ColorNamepowderblue:           "ColorNamepowderblue",
		ColorNamepurple:               "ColorNamepurple",
		ColorNamered:                  "ColorNamered",
		ColorNamerosybrown:            "ColorNamerosybrown",
		ColorNameroyalblue:            "ColorNameroyalblue",
		ColorNamesaddlebrown:          "ColorNamesaddlebrown",
		ColorNamesalmon:               "ColorNamesalmon",
		ColorNamesandybrown:           "ColorNamesandybrown",
		ColorNameseagreen:             "ColorNameseagreen",
		ColorNameseashell:             "ColorNameseashell",
		ColorNamesienna:               "ColorNamesienna",
		ColorNamesilver:               "ColorNamesilver",
		ColorNameskyblue:              "ColorNameskyblue",
		ColorNameslateblue:            "ColorNameslateblue",
		ColorNameslategray:            "ColorNameslategray",
		ColorNameslategrey:            "ColorNameslategrey",
		ColorNamesnow:                 "ColorNamesnow",
		ColorNamespringgreen:          "ColorNamespringgreen",
		ColorNamesteelblue:            "ColorNamesteelblue",
		ColorNametan:                  "ColorNametan",
		ColorNameteal:                 "ColorNameteal",
		ColorNamethistle:              "ColorNamethistle",
		ColorNametomato:               "ColorNametomato",
		ColorNameturquoise:            "ColorNameturquoise",
		ColorNameviolet:               "ColorNameviolet",
		ColorNamewheat:                "ColorNamewheat",
		ColorNamewhite:                "ColorNamewhite",
		ColorNamewhitesmoke:           "ColorNamewhitesmoke",
		ColorNameyellow:               "ColorNameyellow",
		ColorNameyellowgreen:          "ColorNameyellowgreen",
	}
)

func init() {
	var v ColorName
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ColorNameNameToValue = map[string]ColorName{
			interface{}(ColorNamealiceblue).(fmt.Stringer).String():            ColorNamealiceblue,
			interface{}(ColorNameantiquewhite).(fmt.Stringer).String():         ColorNameantiquewhite,
			interface{}(ColorNameaqua).(fmt.Stringer).String():                 ColorNameaqua,
			interface{}(ColorNameaquamarine).(fmt.Stringer).String():           ColorNameaquamarine,
			interface{}(ColorNameazure).(fmt.Stringer).String():                ColorNameazure,
			interface{}(ColorNamebeige).(fmt.Stringer).String():                ColorNamebeige,
			interface{}(ColorNamebisque).(fmt.Stringer).String():               ColorNamebisque,
			interface{}(ColorNameblanchedalmond).(fmt.Stringer).String():       ColorNameblanchedalmond,
			interface{}(ColorNameblue).(fmt.Stringer).String():                 ColorNameblue,
			interface{}(ColorNameblueviolet).(fmt.Stringer).String():           ColorNameblueviolet,
			interface{}(ColorNamebrown).(fmt.Stringer).String():                ColorNamebrown,
			interface{}(ColorNameburlywood).(fmt.Stringer).String():            ColorNameburlywood,
			interface{}(ColorNamecadetblue).(fmt.Stringer).String():            ColorNamecadetblue,
			interface{}(ColorNamechartreuse).(fmt.Stringer).String():           ColorNamechartreuse,
			interface{}(ColorNamechocolate).(fmt.Stringer).String():            ColorNamechocolate,
			interface{}(ColorNamecoral).(fmt.Stringer).String():                ColorNamecoral,
			interface{}(ColorNamecornflowerblue).(fmt.Stringer).String():       ColorNamecornflowerblue,
			interface{}(ColorNamecornsilk).(fmt.Stringer).String():             ColorNamecornsilk,
			interface{}(ColorNamecrimson).(fmt.Stringer).String():              ColorNamecrimson,
			interface{}(ColorNamecyan).(fmt.Stringer).String():                 ColorNamecyan,
			interface{}(ColorNamedarkblue).(fmt.Stringer).String():             ColorNamedarkblue,
			interface{}(ColorNamedarkcyan).(fmt.Stringer).String():             ColorNamedarkcyan,
			interface{}(ColorNamedarkgoldenrod).(fmt.Stringer).String():        ColorNamedarkgoldenrod,
			interface{}(ColorNamedarkgray).(fmt.Stringer).String():             ColorNamedarkgray,
			interface{}(ColorNamedarkgreen).(fmt.Stringer).String():            ColorNamedarkgreen,
			interface{}(ColorNamedarkgrey).(fmt.Stringer).String():             ColorNamedarkgrey,
			interface{}(ColorNamedarkkhaki).(fmt.Stringer).String():            ColorNamedarkkhaki,
			interface{}(ColorNamedarkmagenta).(fmt.Stringer).String():          ColorNamedarkmagenta,
			interface{}(ColorNamedarkolivegreen).(fmt.Stringer).String():       ColorNamedarkolivegreen,
			interface{}(ColorNamedarkorange).(fmt.Stringer).String():           ColorNamedarkorange,
			interface{}(ColorNamedarkorchid).(fmt.Stringer).String():           ColorNamedarkorchid,
			interface{}(ColorNamedarkred).(fmt.Stringer).String():              ColorNamedarkred,
			interface{}(ColorNamedarksalmon).(fmt.Stringer).String():           ColorNamedarksalmon,
			interface{}(ColorNamedarkseagreen).(fmt.Stringer).String():         ColorNamedarkseagreen,
			interface{}(ColorNamedarkslateblue).(fmt.Stringer).String():        ColorNamedarkslateblue,
			interface{}(ColorNamedarkslategray).(fmt.Stringer).String():        ColorNamedarkslategray,
			interface{}(ColorNamedarkslategrey).(fmt.Stringer).String():        ColorNamedarkslategrey,
			interface{}(ColorNamedarkturquoise).(fmt.Stringer).String():        ColorNamedarkturquoise,
			interface{}(ColorNamedarkviolet).(fmt.Stringer).String():           ColorNamedarkviolet,
			interface{}(ColorNamedeeppink).(fmt.Stringer).String():             ColorNamedeeppink,
			interface{}(ColorNamedeepskyblue).(fmt.Stringer).String():          ColorNamedeepskyblue,
			interface{}(ColorNamedimgray).(fmt.Stringer).String():              ColorNamedimgray,
			interface{}(ColorNamedimgrey).(fmt.Stringer).String():              ColorNamedimgrey,
			interface{}(ColorNamedodgerblue).(fmt.Stringer).String():           ColorNamedodgerblue,
			interface{}(ColorNamefirebrick).(fmt.Stringer).String():            ColorNamefirebrick,
			interface{}(ColorNamefloralwhite).(fmt.Stringer).String():          ColorNamefloralwhite,
			interface{}(ColorNameforestgreen).(fmt.Stringer).String():          ColorNameforestgreen,
			interface{}(ColorNamefuchsia).(fmt.Stringer).String():              ColorNamefuchsia,
			interface{}(ColorNamegainsboro).(fmt.Stringer).String():            ColorNamegainsboro,
			interface{}(ColorNameghostwhite).(fmt.Stringer).String():           ColorNameghostwhite,
			interface{}(ColorNamegold).(fmt.Stringer).String():                 ColorNamegold,
			interface{}(ColorNamegoldenrod).(fmt.Stringer).String():            ColorNamegoldenrod,
			interface{}(ColorNamegray).(fmt.Stringer).String():                 ColorNamegray,
			interface{}(ColorNamegreen).(fmt.Stringer).String():                ColorNamegreen,
			interface{}(ColorNamegreenyellow).(fmt.Stringer).String():          ColorNamegreenyellow,
			interface{}(ColorNamegrey).(fmt.Stringer).String():                 ColorNamegrey,
			interface{}(ColorNamehomeassistant).(fmt.Stringer).String():        ColorNamehomeassistant,
			interface{}(ColorNamehoneydew).(fmt.Stringer).String():             ColorNamehoneydew,
			interface{}(ColorNamehotpink).(fmt.Stringer).String():              ColorNamehotpink,
			interface{}(ColorNameindianred).(fmt.Stringer).String():            ColorNameindianred,
			interface{}(ColorNameindigo).(fmt.Stringer).String():               ColorNameindigo,
			interface{}(ColorNameivory).(fmt.Stringer).String():                ColorNameivory,
			interface{}(ColorNamekhaki).(fmt.Stringer).String():                ColorNamekhaki,
			interface{}(ColorNamelavender).(fmt.Stringer).String():             ColorNamelavender,
			interface{}(ColorNamelavenderblush).(fmt.Stringer).String():        ColorNamelavenderblush,
			interface{}(ColorNamelawngreen).(fmt.Stringer).String():            ColorNamelawngreen,
			interface{}(ColorNamelemonchiffon).(fmt.Stringer).String():         ColorNamelemonchiffon,
			interface{}(ColorNamelightblue).(fmt.Stringer).String():            ColorNamelightblue,
			interface{}(ColorNamelightcoral).(fmt.Stringer).String():           ColorNamelightcoral,
			interface{}(ColorNamelightcyan).(fmt.Stringer).String():            ColorNamelightcyan,
			interface{}(ColorNamelightgoldenrodyellow).(fmt.Stringer).String(): ColorNamelightgoldenrodyellow,
			interface{}(ColorNamelightgray).(fmt.Stringer).String():            ColorNamelightgray,
			interface{}(ColorNamelightgreen).(fmt.Stringer).String():           ColorNamelightgreen,
			interface{}(ColorNamelightgrey).(fmt.Stringer).String():            ColorNamelightgrey,
			interface{}(ColorNamelightpink).(fmt.Stringer).String():            ColorNamelightpink,
			interface{}(ColorNamelightsalmon).(fmt.Stringer).String():          ColorNamelightsalmon,
			interface{}(ColorNamelightseagreen).(fmt.Stringer).String():        ColorNamelightseagreen,
			interface{}(ColorNamelightskyblue).(fmt.Stringer).String():         ColorNamelightskyblue,
			interface{}(ColorNamelightslategray).(fmt.Stringer).String():       ColorNamelightslategray,
			interface{}(ColorNamelightslategrey).(fmt.Stringer).String():       ColorNamelightslategrey,
			interface{}(ColorNamelightsteelblue).(fmt.Stringer).String():       ColorNamelightsteelblue,
			interface{}(ColorNamelightyellow).(fmt.Stringer).String():          ColorNamelightyellow,
			interface{}(ColorNamelime).(fmt.Stringer).String():                 ColorNamelime,
			interface{}(ColorNamelimegreen).(fmt.Stringer).String():            ColorNamelimegreen,
			interface{}(ColorNamelinen).(fmt.Stringer).String():                ColorNamelinen,
			interface{}(ColorNamemagenta).(fmt.Stringer).String():              ColorNamemagenta,
			interface{}(ColorNamemaroon).(fmt.Stringer).String():               ColorNamemaroon,
			interface{}(ColorNamemediumaquamarine).(fmt.Stringer).String():     ColorNamemediumaquamarine,
			interface{}(ColorNamemediumblue).(fmt.Stringer).String():           ColorNamemediumblue,
			interface{}(ColorNamemediumorchid).(fmt.Stringer).String():         ColorNamemediumorchid,
			interface{}(ColorNamemediumpurple).(fmt.Stringer).String():         ColorNamemediumpurple,
			interface{}(ColorNamemediumseagreen).(fmt.Stringer).String():       ColorNamemediumseagreen,
			interface{}(ColorNamemediumslateblue).(fmt.Stringer).String():      ColorNamemediumslateblue,
			interface{}(ColorNamemediumspringgreen).(fmt.Stringer).String():    ColorNamemediumspringgreen,
			interface{}(ColorNamemediumturquoise).(fmt.Stringer).String():      ColorNamemediumturquoise,
			interface{}(ColorNamemediumvioletred).(fmt.Stringer).String():      ColorNamemediumvioletred,
			interface{}(ColorNamemidnightblue).(fmt.Stringer).String():         ColorNamemidnightblue,
			interface{}(ColorNamemintcream).(fmt.Stringer).String():            ColorNamemintcream,
			interface{}(ColorNamemistyrose).(fmt.Stringer).String():            ColorNamemistyrose,
			interface{}(ColorNamemoccasin).(fmt.Stringer).String():             ColorNamemoccasin,
			interface{}(ColorNamenavajowhite).(fmt.Stringer).String():          ColorNamenavajowhite,
			interface{}(ColorNamenavy).(fmt.Stringer).String():                 ColorNamenavy,
			interface{}(ColorNamenavyblue).(fmt.Stringer).String():             ColorNamenavyblue,
			interface{}(ColorNameoldlace).(fmt.Stringer).String():              ColorNameoldlace,
			interface{}(ColorNameolive).(fmt.Stringer).String():                ColorNameolive,
			interface{}(ColorNameolivedrab).(fmt.Stringer).String():            ColorNameolivedrab,
			interface{}(ColorNameorange).(fmt.Stringer).String():               ColorNameorange,
			interface{}(ColorNameorangered).(fmt.Stringer).String():            ColorNameorangered,
			interface{}(ColorNameorchid).(fmt.Stringer).String():               ColorNameorchid,
			interface{}(ColorNamepalegoldenrod).(fmt.Stringer).String():        ColorNamepalegoldenrod,
			interface{}(ColorNamepalegreen).(fmt.Stringer).String():            ColorNamepalegreen,
			interface{}(ColorNamepaleturquoise).(fmt.Stringer).String():        ColorNamepaleturquoise,
			interface{}(ColorNamepalevioletred).(fmt.Stringer).String():        ColorNamepalevioletred,
			interface{}(ColorNamepapayawhip).(fmt.Stringer).String():           ColorNamepapayawhip,
			interface{}(ColorNamepeachpuff).(fmt.Stringer).String():            ColorNamepeachpuff,
			interface{}(ColorNameperu).(fmt.Stringer).String():                 ColorNameperu,
			interface{}(ColorNamepink).(fmt.Stringer).String():                 ColorNamepink,
			interface{}(ColorNameplum).(fmt.Stringer).String():                 ColorNameplum,
			interface{}(ColorNamepowderblue).(fmt.Stringer).String():           ColorNamepowderblue,
			interface{}(ColorNamepurple).(fmt.Stringer).String():               ColorNamepurple,
			interface{}(ColorNamered).(fmt.Stringer).String():                  ColorNamered,
			interface{}(ColorNamerosybrown).(fmt.Stringer).String():            ColorNamerosybrown,
			interface{}(ColorNameroyalblue).(fmt.Stringer).String():            ColorNameroyalblue,
			interface{}(ColorNamesaddlebrown).(fmt.Stringer).String():          ColorNamesaddlebrown,
			interface{}(ColorNamesalmon).(fmt.Stringer).String():               ColorNamesalmon,
			interface{}(ColorNamesandybrown).(fmt.Stringer).String():           ColorNamesandybrown,
			interface{}(ColorNameseagreen).(fmt.Stringer).String():             ColorNameseagreen,
			interface{}(ColorNameseashell).(fmt.Stringer).String():             ColorNameseashell,
			interface{}(ColorNamesienna).(fmt.Stringer).String():               ColorNamesienna,
			interface{}(ColorNamesilver).(fmt.Stringer).String():               ColorNamesilver,
			interface{}(ColorNameskyblue).(fmt.Stringer).String():              ColorNameskyblue,
			interface{}(ColorNameslateblue).(fmt.Stringer).String():            ColorNameslateblue,
			interface{}(ColorNameslategray).(fmt.Stringer).String():            ColorNameslategray,
			interface{}(ColorNameslategrey).(fmt.Stringer).String():            ColorNameslategrey,
			interface{}(ColorNamesnow).(fmt.Stringer).String():                 ColorNamesnow,
			interface{}(ColorNamespringgreen).(fmt.Stringer).String():          ColorNamespringgreen,
			interface{}(ColorNamesteelblue).(fmt.Stringer).String():            ColorNamesteelblue,
			interface{}(ColorNametan).(fmt.Stringer).String():                  ColorNametan,
			interface{}(ColorNameteal).(fmt.Stringer).String():                 ColorNameteal,
			interface{}(ColorNamethistle).(fmt.Stringer).String():              ColorNamethistle,
			interface{}(ColorNametomato).(fmt.Stringer).String():               ColorNametomato,
			interface{}(ColorNameturquoise).(fmt.Stringer).String():            ColorNameturquoise,
			interface{}(ColorNameviolet).(fmt.Stringer).String():               ColorNameviolet,
			interface{}(ColorNamewheat).(fmt.Stringer).String():                ColorNamewheat,
			interface{}(ColorNamewhite).(fmt.Stringer).String():                ColorNamewhite,
			interface{}(ColorNamewhitesmoke).(fmt.Stringer).String():           ColorNamewhitesmoke,
			interface{}(ColorNameyellow).(fmt.Stringer).String():               ColorNameyellow,
			interface{}(ColorNameyellowgreen).(fmt.Stringer).String():          ColorNameyellowgreen,
		}
	}
}

// MarshalJSON is generated so ColorName satisfies json.Marshaler.
func (r ColorName) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ColorNameValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ColorName: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ColorName satisfies json.Unmarshaler.
func (r *ColorName) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ColorName should be a string, got %s", data)
	}
	v, ok := _ColorNameNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ColorName %q", s)
	}
	*r = v
	return nil
}
