package config

import "github.com/joho/godotenv"

// Global `env` map
var env map[string]string = initEnv()

// Initialize the global `env` variable
func initEnv() map[string]string {
	e, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	return e
}

// Retrieves the whole global `env`
// Runs `initEnv()` if needed
func GetEnv() map[string]string {
	if len(env) == 0 {
		env = initEnv()
	}

	return env
}

// Retrieves an Env property by its name
func GetEnvItem(item string) string {
	e := GetEnv()
	return e[item]
}
