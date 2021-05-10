package main

import (
	"fmt"	
	"time"
)


func dBCtoSBC(s string) string {
	retstr := ""
	for _, i := range s {
		insideCode := i
		if insideCode == 12288 {
			insideCode = 32
		} else {
			insideCode -= 65248
		}
		if insideCode < 32 || insideCode > 126 {
			retstr += string(i)
		} else {
			retstr += string(insideCode)
		}
	}
	return retstr
}

func testToSingle(s string) {
	fmt.Printf("%s to single is %s\n", s, dBCtoSBC(s))
}

func testTime() {
	now := time.Now()
	fmt.Printf("%d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	now.Sub(now)
}

// 获取当前时间的学习周 (当前第几周，该学期共多少周, 年级第几学期)
func getCurrentLearnWeek() (int, int, int) {
	now := time.Now()
	const shortForm = "2006-Jan-02"
	periodIndex := 1
	weekCount := 17
	var weekIndex int
	if now.Month() >= 9 || now.Month() < 3 {
		// 上学期
		periodIndex = 0
		startYear := now.Year()
		if now.Month() < 3 {
			startYear--
		}
		periodStartTime, _ := time.Parse(shortForm, fmt.Sprintf("%d-Sep-01", startYear))
		intervalDay := int(now.Sub(periodStartTime).Hours() / 24)
		if now.Month() < 3 || now.Month() > 10 {
			// 去除国庆3天
			intervalDay -= 3
		} else if now.Month() == 10 {
			if now.Day() > 3 {
				intervalDay -= 3
			} else {
				intervalDay -= now.Day()
			}
		}
		weekIndex = intervalDay / 7
	} else {
		periodStartTime, _ := time.Parse(shortForm, fmt.Sprintf("%d-Mar-01", now.Year()))
		intervalDay := int(now.Sub(periodStartTime).Hours() / 24)
		if now.Month() >= 5 {
			// 减去五一1天
			intervalDay--
		}
		weekIndex = intervalDay / 7
	}
	return weekIndex, weekCount, periodIndex
}

func main() {
	str := "我是hello小熊"
	strRune := []rune(str)
	endIdx := len(strRune) - 1
	fmt.Printf("sub %s\n", string(strRune[:endIdx]))

	
	testToSingle(str)
	testToSingle("123324kjsdf反抗军地方．，．．..,,")
	testToSingle("13234１２３４５６７８９０可接受的。。，，")
	testTime()

	wi, wc, pi := getCurrentLearnWeek()
	fmt.Printf("第%d学期 第%d周 共%d周\n", pi, wi, wc)

	md5Str := "kjsdfwewefwefwefkwjfwjewfwef"
	arr := []rune(md5Str)
	path := fmt.Sprintf("%s/%s/%s/%s/%s", string(arr[0:2]), string(arr[2:4]), string(arr[4:6]), string(arr[6:8]), md5Str)
	fmt.Println(path)
}