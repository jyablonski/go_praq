package main

import (
	"errors"
	"fmt"
)

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)

	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]

		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}

	return userMap, nil
}

type UserTwo struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]UserTwo, name string) (deleted bool, err error) {
	existingUser, ok := users[name]

	if !ok {
		return false, errors.New("not found")
	}

	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}

	return false, nil
}

func getCounts(userIDs []string) map[string]int {
	counts := make(map[string]int)

	for _, userID := range userIDs {
		count := counts[userID]
		count++
		counts[userID] = count
	}

	return counts
}

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstChar := rune(name[0])
		_, ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	return counts
}

func main() {
	fmt.Println("Fuq")
}
