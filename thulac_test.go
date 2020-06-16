package thulacgo

import "fmt"

func ExampleTest() {

	lac := NewThulacgo("models", "", false, false, false, byte('_'))
	defer lac.Deinit()

	ret := lac.Seg("我爱北京天安门")

	fmt.Println(ret)

	//Output:
	//我_r 爱_v 北京_ns 天安门_ns
}
