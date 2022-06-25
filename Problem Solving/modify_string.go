package main

// import (
//     "bufio"
//     "fmt"
//     "io"
//     "os"
//     "strings"
// )



/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as parameter.
 */



func ModifyString(str string) string{
	answer := ""

	for i:=len(str)-1; i >= 0; i-- {
		if str[i]==32{
			continue
		} else if str[i]>=48 && str[i]<=57{
			continue
		} else {
			answer += string(str[i])
		}
	}

	return answer

}

func main() {
	var x string = "  0he219ll  o9 "
	print(ModifyString(x))
}