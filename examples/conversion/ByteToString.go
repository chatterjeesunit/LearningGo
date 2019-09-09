package main

import "unsafe"

func main() {

	b := []byte("Don''t forget to call rand.Seed() to properly initialize it before you use the math/rand package")
	//length := len(b)
	arr := make([]*string, 99999999)
	//builder := new(strings.Builder)
	//
	//for i:=0; i< 99999999; i++ {
	//	for j:=0; j< 100; j++ {
	//		builder.Write(b)
	//		s :=  builder.String()
	//		arr[j] = &s
	//		builder.Reset()
	//	}
	//
	//}
	for i:=0; i< 99999999; i++ {
		for j:=0; j< 99999999; j++ {
			//var s string
			//stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
			//stringHeader.Data = uintptr(unsafe.Pointer(&b[0]))
			//stringHeader.Len = length
			//println(s)
			arr[j] = (*string)(unsafe.Pointer(&b))
			println(*arr[j], len(*arr[j]))

		}

	}
}
