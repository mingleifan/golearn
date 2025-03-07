package main

import (
	"fmt"
	"maps"
)

func main() {

	cityMap := make(map[string]string)

	//添加
	cityMap["CHINA"] = "BEIJING"
	cityMap["JAPAN"] = "TOKYO"
	cityMap["USA"] = "NEW_YORK"

	//遍历
	printMap(cityMap)

	//删除
	delete(cityMap, "JAPAN")

	fmt.Println("---------")

	//修改
	cityMap["USA"] = "WASHINGTON"
	changeMapV(cityMap, "USA")

	printMap(cityMap)

	fmt.Println("-----copied----")
	copiedCityMap := make(map[string]string)
	maps.Copy(copiedCityMap, cityMap)
	printMap(copiedCityMap)

}

func printMap(cityMap map[string]string) {
	//引用传递，指向同一个内存空间
	for k, v := range cityMap {
		fmt.Printf("key = %s, value = %s\n", k, v)
	}
}

func changeMapV(cityMap map[string]string, key string) {
	cityMap[key] = "WASHINGTON_BAK"
}
