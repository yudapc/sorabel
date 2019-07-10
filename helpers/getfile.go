package helpers

import "os"

func GetFile(path string) string {
	dir, _ := os.Getwd()
	pathSchema := dir + path
	return "file:///" + pathSchema
}
