package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	var num int
	var key string
	var email string
	searchStr := "." + domain
	scanner := bufio.NewScanner(r)
	result := make(DomainStat, 100000)

	for scanner.Scan() {
		email = fastjson.GetString(scanner.Bytes(), "Email")
		if strings.HasSuffix(email, searchStr) {
			key = strings.ToLower(strings.SplitN(email, "@", 2)[1])
			num = result[key]
			num++
			result[key] = num
		}
	}

	return result, nil
}
