package smalltools

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAll(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}

func WriteToFile(fileName, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		fmt.Println("write succeed!")
		defer f.Close()
	}
	return err
}
