// Checking if the required file is present in the filepath

package pkg

import (
	"fmt"
	"os"
)

func CheckFileExists(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does dot exist! %s", filePath)
	}

	return err
}
