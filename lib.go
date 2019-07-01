package go_try

import "reflect"

var TryError error

func Try(params ...interface{}) []interface{} {
	l := len(params)
	if len(params) == 0 {
		return nil
	}
	e := params[l-1]
	if !isError(e) {
		return params[:l-1]
	}
	TryError = e.(error)
	panic(e)
}

func isError(e interface{}) bool {
	ef := reflect.TypeOf(e)
	if ef == nil {
		return false
	}
	return ef.Implements(reflect.TypeOf((*error)(nil)).Elem())
}
