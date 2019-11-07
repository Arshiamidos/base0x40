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

func toBin(b []byte) string {
	s := ""
	for i := 0; i < len(b); i++ {
		s = s + fmt.Sprintf("%08v", (strconv.FormatInt(int64(b[i]), 2)))
	}
	return s
}
func toInt(s []string) []int8 {
	i := make([]int8, 0)
	for _i := 0; _i < len(s); _i++ {
		n, _ := strconv.ParseInt(s[_i], 2, 8)
		i = append(i, int8(n))
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
	f, _ := os.Open("3.txt")
	byts, _ := ioutil.ReadAll(f)
	//ind := 0 //strings.Index(string(byts), "base64,") + 7
	//byts = byts[ind:]
	sbyts := toBin(byts)
	if len(sbyts)%6 != 0 {
		sbyts = sbyts + strings.Repeat("0", 6-len(sbyts)%6)
	}
	rgx, _ := regexp.Compile("......")
	SubMatchedBytes := rgx.FindAllString(sbyts, -1)
	SubMatchInts := toInt(SubMatchedBytes)
	fmt.Println("{}", SubMatchedBytes[len(SubMatchedBytes)-1], len(sbyts)%6)
	//fmt.Println(SubMatchedBytes[0:3], SubMatchInts[:3])
	fmt.Println(deMultiplex(SubMatchInts))
	// data:image/jpeg;base64, ... ==
}
