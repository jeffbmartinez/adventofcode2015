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

		if i%100001 == 100000 {
			fmt.Printf("i = %v...\n", i)
		}

		test := key + strconv.Itoa(i)
		md5String := MD5Hash(test)

		targetPrefix := "000000"
		if md5String[:len(targetPrefix)] == targetPrefix {
			fmt.Printf("First hash with prefix '000000' is %v, from %v\n", md5String, test)
			break
		}
	}
}

func MD5Hash(s string) string {
	md5Hash := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", md5Hash)
}
