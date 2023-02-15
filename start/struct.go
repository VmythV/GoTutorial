package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

func main2() {

	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)

}

type location struct {
	Lat  float64
	Long float64
}

type dis struct {
	lat  float64
	long float64
}

func distance(loc1, loc2 location) dis {
	return dis{
		math.Abs(loc2.Lat - loc1.Lat), math.Abs(loc2.Long - loc1.Long),
	}
}

func main() {
	var spirit location
	//spirit.Lat = -14.5684
	//spirit.Long = 175.472636
	// 可以少些参数，这样少写的参数就是原本类型的零值
	spirit = location{Lat: -14.5684, Long: 175.472636}

	var opportunity location
	//opportunity.Lat = -1.9462
	//opportunity.Long = 354.4734
	opportunity = location{-1.9462, 354.4734}

	fmt.Println(spirit, opportunity)
	d := distance(spirit, opportunity)
	fmt.Println(d)
	fmt.Printf("%v\n", d)
	fmt.Printf("%+v\n", d)

	fmt.Println("==============")
	// 结构体是值传递
	oppo := opportunity
	oppo.Lat += 100
	fmt.Println(opportunity, oppo)

	fmt.Println("==============")
	// 将结构体转化成json
	// 这里有点坑，如果结构体中的变量首字母是小写的话，是打印不出的，必须是大写的才能打印出来
	// Go中 大写是可见的，小写是不可见
	// Marshal函数只会对struct中被导出的字段进行编码
	bytes, err := json.Marshal(oppo)
	exitOnError(err)
	fmt.Println(string(bytes))

	fmt.Println("==============")
	// 使用struct标签来自定义JSON
	// Go语言中的json包要求struct中的字段必须以大写字母开头，类似CamelCase驼峰型命名规范
	// 如果需要snake_case蛇形命名规范，可以为字段注明标签，使得json包在进行编码的时候能够按照标签里的样式修改字段名
	type location2 struct {
		Lat  float64 `json:"latitude"xml:"latitudexml"`
		Long float64 `json:"longitude"xml:"latitudexml"`
	}
	curiosity := location2{-4.5895, 137.4417}

	bytes, err = json.MarshalIndent(curiosity, "", " ")
	exitOnError(err)

	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
