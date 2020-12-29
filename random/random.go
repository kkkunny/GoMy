package random

import (
	"math"
	"math/rand"
	"reflect"
	"time"
)

// 随机整数[start, end)
func Int(start int, end int)int{
	rand.Seed(time.Now().UnixNano())
	return start + rand.Intn(end-start)
}
// 随机浮点数[start, end)
func Float(start float64, end float64, round int)float64{
	temp := math.Pow(10, float64(round))
	startInt, endInt := int(start*temp), int(end*temp)
	resultInt := Int(startInt, endInt)
	return float64(resultInt) / temp
}
// 随机布尔值
func Bool()bool{
	temp := Int(0, 2)
	if temp == 0{
		return false
	}else{
		return true
	}
}
// 列表中随机取一个值（有重复）
func GetOne(slice interface{})interface{}{
	value := reflect.ValueOf(slice)
	i := Int(0, value.Len())
	return value.Index(i).Interface()
}
// 列表中根据比例随机取一个值（有重复）
func GetOneByRatio(slice interface{}, ratioes []int)interface{}{
	value := reflect.ValueOf(slice)
	length := value.Len()
	if len(ratioes) != length{
		panic("both of them have different lengths")
	}
	var temps []interface{}
	for i:=0; i<length; i++{
		for j:=0; j<ratioes[i]; j++{
			temps = append(temps, value.Index(i).Interface())
		}
	}
	return GetOne(temps)
}
// 随机顺序
func Sort(slice interface{})[]interface{}{
	value := reflect.ValueOf(slice)
	order := make(map[int]uint8)
	for i:=0; i<value.Len(); i++{
		order[i] = 0
	}
	var result []interface{}
	for k, _ := range order{
		result = append(result, value.Index(k).Interface())
	}
	return result
}