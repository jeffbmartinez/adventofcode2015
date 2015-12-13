// puzzle at http://adventofcode.com/day/4

package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

var key string = "ckczppom"

func main() {

	for i := 0; ; i++ {

		if i%10001 == 10000 {
			fmt.Printf("i = %v...\n", i)
		}

		test := key + strconv.Itoa(i)
		md5String := MD5Hash(test)

		if md5String[:5] == "00000" {
			fmt.Printf("First hash with prefix '00000' is %v, from %v\n", md5String, test)
			break
		}
	}
}

func MD5Hash(s string) string {
	md5Hash := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", md5Hash)
}
