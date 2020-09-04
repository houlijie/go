package file

import "os"

func write() {
	filePath := "example.txt"
	file, err := os.OpenFile(filePath, O_WRONLY|O_CREATE, 0666)
	if err != nil {
		return
	}

	defer file.Close()

	str := "hello world"
}
