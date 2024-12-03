package services

import (
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID(userA int, userB int) string {
	listNumbers := []string{strconv.Itoa(userA), strconv.Itoa(userB)}
	sort.Strings(listNumbers)
	users_combined := strings.Join(listNumbers, "-")
	result := uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	return uuid.NewMD5(result, []byte(users_combined)).String()
}
