package api

//Handle ..
type Handle struct {
}

//CreateHandle ..
func CreateHandle() *Handle {
	return &Handle{}
}

//GetHandle ..
func GetHandle() *Handle {
	handle := CreateHandle()

	return handle
}
