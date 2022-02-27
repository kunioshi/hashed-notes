package config

import "github.com/joho/godotenv"

var env map[string]string

func GetEnv() map[string]string {
	if len(env) == 0 {
		newEnv, err := godotenv.Read()
		if err != nil {
			panic(err)
		}

		env = newEnv
	}

	return env
}

func GetEnvItem(item string) string {
	curEnv := GetEnv()
	return curEnv[item]
}
