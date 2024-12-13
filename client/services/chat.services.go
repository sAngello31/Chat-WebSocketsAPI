package services

import (
	"sort"
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID(userA, userB string) string {
	listNumbers := []string{userA, userB}
	sort.Strings(listNumbers)
	user_combined := strings.Join(listNumbers, "-")
	space := uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	return uuid.NewMD5(space, []byte(user_combined)).String()
}
