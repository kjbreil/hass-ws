package entities

import (
	"fmt"
	"reflect"
	"strings"
)

func fillFields(object interface{}, attributes map[string]interface{}) {

	v := reflect.ValueOf(object).Elem()
	t := reflect.TypeOf(object).Elem()
	for i := 0; i < v.NumField(); i++ {
		//d := v.Field(i)
		//n := t.Field(i).name
		tg := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
		if d, ok := attributes[tg]; ok {
			if !v.Field(i).CanSet() {
				continue
			}
			ty := v.Field(i).Type().Elem().Kind()
			switch ty {
			case reflect.Int:
				ins := toInt(d)
				v.Field(i).Set(reflect.ValueOf(&ins))
			case reflect.Float64:
				fl := toFloat(d)
				v.Field(i).Set(reflect.ValueOf(&fl))
			case reflect.String:
				st := d.(string)
				v.Field(i).Set(reflect.ValueOf(&st))
			case reflect.Slice:
				//ss := toStringSlice(d)
				//v.Field(i).Set(reflect.ValueOf(&ss))
			default:
				panic(fmt.Errorf("not setup to handle %s in fillFields", ty))
			}
			delete(attributes, tg)
		}
	}
}

func toStringSlice(d interface{}) []string {
	ds := d.([]interface{})
	rtn := make([]string, len(ds))
	for i := range ds {
		switch ds[i].(type) {
		case string:
			rtn[i] = ds[i].(string)

		}
	}
	return rtn
}

func toInt(d interface{}) int {
	switch d.(type) {
	case float64:
		return int(d.(float64))
	default:
		return d.(int)
	}
}

func toFloat(d interface{}) float64 {
	switch d.(type) {
	case float64:
		return d.(float64)
	default:
		return 0.0
	}
}
