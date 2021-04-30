package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func PrintNumber(g string) {
	for i, v := range g {
		nol := ""
		total := len(g) - 1 - i
		for i := 0; i < total; i++ {
			nol = nol + "0"
		}
		fmt.Println(string(v) + nol)
	}
}

func main() {
	// total := "1.345.679"
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	s := readLine(reader)
	g := strings.Replace(s, ".", "", -1)
	_, err := strconv.Atoi(g)
	if err != nil {
		fmt.Println("input salah")
		return
	}
	PrintNumber(g)
}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
