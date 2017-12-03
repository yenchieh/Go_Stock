package config

import "fmt"

type Environment struct {
	Debug           bool
	AlphaVantageKey string
}

var Env Environment

func SetupEnv(env Environment) {
	Env = env
	fmt.Printf("\n Environment Variable: %#v \n", Env)
}
