package api

import (
	"regexp"
	"strconv"
)

func getToken(message string) uint64 {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	stringToken := re.FindAllString(message, -1)
	token, _ := strconv.Atoi(stringToken[0])
	return uint64(token)
}
