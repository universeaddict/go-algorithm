package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SockMerchant(n int, ar []int) {
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	pairs := 0
	count := 0
	last := 0
	for i, v := range ar {
		if last == v && i == len(ar)-1 {
			count += 1
			sum := count / 2
			pairs += sum
		} else if last == v {
			count = count + 1
		} else {
			sum := count / 2
			pairs += sum
			count = 1
		}
		last = v
	}
	fmt.Println(pairs)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	reader1 := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	s1 := readLine(reader)
	v1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("input value 1 salah")
		return
	}
	s2 := readLine(reader1)
	v2 := strings.Split(s2, " ")
	ints := make([]int, len(v2))
	for i, s := range v2 {
		ints[i], _ = strconv.Atoi(s)
	}
	SockMerchant(v1, ints)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
