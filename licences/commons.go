package licences

import (
	"os"
	"path"
	"strconv"
	"time"
)

var year string = strconv.Itoa(time.Now().Year())

func WriteFile(content []byte) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filename := path.Join(wd, "LICENCE")
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		panic(err)
	}
}
