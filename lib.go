package go_try

var TryError error

func Try(v interface{}, e error) interface{} {
	if e == nil {
		return v
	}
	TryError = e
	panic(e)
}
