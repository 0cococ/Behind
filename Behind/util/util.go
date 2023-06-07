package util

import "math/rand"

func RandomSring(n int) string {
	letters := []byte("aewifdjpoaekjfvpijam210139120939")
	result := make([]byte, n)
	//rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
