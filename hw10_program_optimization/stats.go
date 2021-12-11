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
	searchStr := "." + domain
	scanner := bufio.NewScanner(r)
	result := make(DomainStat)

	var email string

	for scanner.Scan() {
		email = fastjson.GetString(scanner.Bytes(), "Email")

		if strings.HasSuffix(email, searchStr) {
			num = result[strings.ToLower(strings.SplitN(email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(email, "@", 2)[1])] = num
		}
	}

	return result, nil
}
