package main

import (
	"fmt"
	"strings"
)

func main() {
	//判断是否以prefix开头和suffix结尾
	str := "This is an an example"
	fmt.Println(strings.HasPrefix(str, "Th"))
	fmt.Println(strings.HasSuffix(str, "ple"))
	//判断字符串是否包含
	fmt.Println(strings.Contains(str, "an"))
	//判断字符串首次/末次出现的位置
	fmt.Println(strings.Index(str, "an"))
	fmt.Println(strings.LastIndex(str, "an"))
	//替换
	fmt.Println(strings.Replace(str, "an", "two", -1))
	//统计字符串次数
	fmt.Println(strings.Count(str, "an"))
	//分割字符串，分割成slice
	str_ori := "www.newscctv.net"
	slice := strings.Split(str_ori, ".")
	fmt.Println(slice)
	//拼接字符串，slice->string
	fmt.Println(strings.Join(slice, "/"))
}
