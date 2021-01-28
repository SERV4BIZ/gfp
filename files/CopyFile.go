package files

import (
	"io/ioutil"
)

// CopyFile is copy file from src to dst
func CopyFile(src string, dst string) error {
	//Read all the contents of the  original file
	bytesRead, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	//Copy all the contents to the desitination file
	err = ioutil.WriteFile(dst, bytesRead, 0777)
	if err != nil {
		return err
	}

	return nil
}
