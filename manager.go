package utils

import "github.com/goplit/utils/files"

func ReadFile(path string) files.FileProduct {
	return files.ReadFileToProduct(path)
}
