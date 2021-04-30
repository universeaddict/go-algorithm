package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func CountingValleys(n int, ar string) {
	count := 0
	lpos := 0
	cpos := 0
	for _, v := range ar {
		// U = 85
		// D = 68
		if v == 85 {
			cpos += 1
		} else {
			cpos -= 1
		}
		if cpos == 0 && lpos < 0 {
			count += 1
		}
		lpos = cpos
	}
	fmt.Println(count)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	reader1 := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	s1 := readLine(reader)
	v1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("input value 1 salah")
	}
	s2 := readLine(reader1)
	s2 = strings.Replace(s2, " ", "", -1)
	CountingValleys(v1, s2)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
