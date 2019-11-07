package main

//https://www.quora.com/How-does-base64-encoding-work
import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const Base64Map = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func toChar(b []string) string {
	s := ""
	for i := 0; i < len(b); i++ {
		n, _ := strconv.ParseInt(b[i], 2, 64)
		fmt.Println(b[i], "  --  ", n, " - ", string(byte(n)))
		s = s + string(byte(n))
	}
	return s
}
func toBin(b []int) string {
	s := ""
	for i := 0; i < len(b); i++ {
		s = s + fmt.Sprintf("%06v", (strconv.FormatInt(int64(b[i]), 2)))
	}
	return s
}
func toInt(s string) []int {
	i := make([]int, 0)
	fmt.Println(s)
	for _i := 0; _i < len(s); _i++ {
		i = append(i, strings.Index(Base64Map, string(s[_i])))
	}
	return i
}
func deMultiplex(s []int8) string {
	i := ""
	for _i := 0; _i < len(s); _i++ {
		i = i + string(Base64Map[s[_i]])
	}
	return i
}
func main() {
	f, _ := os.Open("4.txt")
	byts, _ := ioutil.ReadAll(f)
	//ind := 0 //strings.Index(string(byts), "base64,") + 7
	//byts = byts[ind:]
	code := string(byts)
	ints := toInt(code)
	fmt.Println(ints)
	sbin := toBin(ints)
	rgx, _ := regexp.Compile("........")
	SubMatchedBytes := rgx.FindAllString(sbin, -1)
	characters := toChar(SubMatchedBytes)
	fmt.Println(characters)
	// sbyts := toBin(byts)
	// if len(sbyts)%6 != 0 {
	// 	sbyts = sbyts + strings.Repeat("0", 6-len(sbyts)%6)
	// }
	// rgx, _ := regexp.Compile("......")
	// SubMatchedBytes := rgx.FindAllString(sbyts, -1)
	// SubMatchInts := toInt(SubMatchedBytes)
	// fmt.Println("{}", SubMatchedBytes[len(SubMatchedBytes)-1], len(sbyts)%6)
	// //fmt.Println(SubMatchedBytes[0:3], SubMatchInts[:3])
	// fmt.Println(deMultiplex(SubMatchInts))
	// data:image/jpeg;base64, ... ==
}
