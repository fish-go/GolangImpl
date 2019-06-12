package reflect

import "reflect"

func typeof(data interface{}) string {
	return reflect.TypeOf(data).String()
}
