package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)                                                             //显示当前完整时间，2020-06-02 11:41:54.636445 +0800 CST m=+0.000264250
	fmt.Println(t.UTC())                                                       //显示世界统一时间，2020-06-02 06:41:53.889064 +0000 UTC
	fmt.Println(t.Unix())                                                      //显示时间戳，1591080113
	fmt.Printf("%d-%02d-%02d 星期%d", t.Year(), t.Month(), t.Day(), t.Weekday()) //获取当天的年月日，2020-06-02 星期2
	fmt.Println(t.Weekday())                                                   //星期，Tuesday
	fmt.Println(t.Month())                                                     //月份，June
	fmt.Println(time.Duration(1000 * 1e9))                                     //时间间隔，16m40s
	fmt.Println(time.Duration(1000 * time.Second))                             //时间间隔，16m40s
	week := 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) //当前时间+一周，2020-06-09 14:30:55.134045 +0800 CST m=+604800.000501034

	//日期转时间戳
	loc, err := time.LoadLocation("Local")                                 //参数可以是Local,UTC
	fmt.Println(loc, err)                                                  //返回时区,(Local,nil)
	timeLayout := "2006-01-02 15:04:05"                                    //转化所需模板
	tmp, _ := time.ParseInLocation(timeLayout, "2020-06-02 15:10:12", loc) //将时间字符串转化为特定时区里的时间，2020-06-02 15:10:12 +0800 CST
	fmt.Println(tmp)
	fmt.Println(tmp.Unix()) //将时间转化为unix时间戳

	//时间戳转日期
	datetime := time.Unix(t.Unix(), 0).Format(timeLayout)
	fmt.Println(datetime)

	//时间延后
	tchan := time.After(3 * time.Second) //3秒后从另一个线程到此线程输出时间
	fmt.Printf("tchan type=%T\n", tchan)
	fmt.Println("mark 1")
	fmt.Println("tchan=", <-tchan)
	fmt.Println("mark 2")
	time.Sleep(2 * time.Second)
	fmt.Println("mark 3")
	fmt.Println(time.Now())
}
