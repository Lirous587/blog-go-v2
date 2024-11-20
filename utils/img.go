package utils

import "strings"

// SanitizedFileName 过滤文件名
func SanitizedFileName(fileName string) string {
	sanitizedFileName := strings.ReplaceAll(fileName, " ", "")
	sanitizedFileName = strings.ReplaceAll(sanitizedFileName, "-", "")
	sanitizedFileName = strings.ReplaceAll(sanitizedFileName, "_", "")
	return sanitizedFileName
}
