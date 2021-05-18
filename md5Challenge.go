//WAP to match md5 signature of files with their original md5, if no match print error

package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func matchMd5(str string) {
	md5Gen := strings.Fields(str)[0]
	fileName := strings.Fields(str)[1]
	//fmt.Printf("%s %s\n",md5Gen,fileName)
	f, err := os.Open("data/Ch06/06_07/" + fileName)
	defer f.Close()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	hsh := md5.New()
	if _, err := io.Copy(hsh, f); err != nil {
		fmt.Errorf("%s\n", err)
	}
	result, _ := ioutil.ReadAll(f)
	hshStr := hex.EncodeToString(hsh.Sum([]byte(string(result))))
	if md5Gen == hshStr {
		
		fmt.Printf("%s : %s Same value of MD5 found!\n", fileName, md5Gen)
	} else {
		errStr := "ERROR: different MD5 checksums found"
		err := fmt.Errorf("%s : %s %s %s", errStr, fileName, md5Gen, hshStr)
		fmt.Println(err.Error())
	}

}

func main() {
	path := "data/Ch06/06_07/md5sum.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	defer file.Close()
	out := make(chan string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go matchMd5(scanner.Text())
	}

}
