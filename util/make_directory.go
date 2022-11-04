package util

import (
	"fmt"
	"os"
)

// ディレクトリ作成
func MakeDirectory(directoryName string) error {
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		os.Mkdir(directoryName, 0777)
		return nil
	} else if err != nil {
		return fmt.Errorf("Failed to create directory : %v", err)
	} else {
		return nil
	}
}
