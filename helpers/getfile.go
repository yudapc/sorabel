package helpers

import "os"

func ProjectDirectory() string {
	dir, _ := os.Getwd()
	return dir
}

func GetFile(path string) string {
	dir := ProjectDirectory()
	pathSchema := dir + path
	return "file:///" + pathSchema
}
