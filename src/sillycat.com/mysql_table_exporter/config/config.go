package config

import (
	"os"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func GetIntEnv(key, defaultValue string) int {
	val := GetEnv(key, defaultValue)
	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("some error"))
	}
	return ret
}
