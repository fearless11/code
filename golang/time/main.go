package main

import (
	"fmt"
	"time"
)

// https://godoc.org/time

// 时间、时区、计算时间差值

/*
时间点：Time
时间段：Duration
时区： Location
格式化：Format
解析： Parse
*/

func main() {
	location()
	format()
	parse()
	calculate()
}

// 时区：默认采用UTC(unix标准时间)
// 时间计量系统： GTM（基于地球自转得出'世界时'）、UTC（基于原子震荡周期确定'原子时'）
// GTM(格林尼治标准时间)： 本初子午线在格林尼治
// UTC(协调世界时)：中国的本地时间比UTC快8小时，就会写作UTC+8
/*
2020-04-11 10:30:49.522686 +0000 UTC
2020-04-11 18:30:49.522686 +0800 CST
2020-04-11 03:30:49.522686 -0700 PDT
*/
func location() {
	now := time.Now()
	local1, err := time.LoadLocation("")
	if err != nil {
		fmt.Println("1ocal1:", err)
	}
	local2, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println("1ocal1:", err)
	}
	local3, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("1ocal1:", err)
	}

	fmt.Println(now.In(local1))
	fmt.Println(now.In(local2))
	fmt.Println(now.In(local3))
}

// 格式化
// 01、02、03、04、05、06、07 对应  月、日、时、分、秒、年、时差
/*
  | 01/JAN | 02 | 03/15 | 04 | 05 | 06 | -07[00][:00] |
  月：   01或Jan都可以
  小时： 03表示12小时制，15表示24小时制
  时差： -07，后边可以增加00或:00,表示更进一步的分秒时差
  上下午: 使用PM
  顺序： 随意，甚至重复都可以
*/
func format() {
	format12 := "2006-01-02 03:04:05"
	format24 := "2006-01-02 15:04:05"
	now := time.Now()
	fmt.Println("format: ", now.Format(format12))
	fmt.Println("format: ", now.Format(format24))
	fmt.Println("format: ", now.Format(time.RFC3339))
}

// 解析
func parse() {
	str := "2018-06-02 13:29:30.6354936 +0800 CST"
	format := "2006-01-02 15:04:05 +0800 CST"
	t, err := time.Parse(format, str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
}

func calculate() {
	lastmonth := time.Now().AddDate(0, -1, 0).Format("200601")
	fmt.Println("lastmonth:", lastmonth)

	format := "2006-01-02T15:04:05+08:00"
	now := time.Now()
	m, _ := time.ParseDuration(fmt.Sprintf("-%vm", 5))
	startT := now.Add(m).Format(format)
	endT := now.Format(format)
	fmt.Println("last 5 minutes", startT, endT)

	now = time.Now()
	time.Sleep(10 * time.Second)
	fmt.Println("since ", time.Since(now))
}
