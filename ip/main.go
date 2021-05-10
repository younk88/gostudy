package main

import (
	"fmt"
	"net"
)

func flatAnswers(s [][]string) [][]string {
	blankCount := len(s)
	if blankCount == 0 {
		return s
	}
	maxMethods := 1
	for _, blankAnswers := range s {
		if len(blankAnswers) > maxMethods {
			maxMethods = len(blankAnswers)
		}
	}
	result := make([][]string, maxMethods)
	for mi := 0; mi < maxMethods; mi++ {
		result[mi] = make([]string, blankCount)
	}
	for bi, blankAnswers := range s {
		bALen := len(blankAnswers)
		lastA := ""
		for ai, a := range blankAnswers {
			result[ai][bi] = a
			lastA = a
		}
		// 不足的答案用最后值组合
		for i := bALen; i < maxMethods; i++ {
			result[i][bi] = lastA
		}
	}
	return result
}

func main() {
	conn, err := net.Dial("udp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	host, port, err := net.SplitHostPort(localAddr)
	host, port, err = net.SplitHostPort(host)
	fmt.Println(host, port, err)	

	answers := [][]string {
		{"1", "2"},
		{"2", "1"},
		{"3"},
	}
	result := flatAnswers(answers)
	fmt.Println(answers)
	fmt.Println(result)
}