package config

// Config is the type of config
type config struct {
	Token   string
	Root    string
	Secret  string
	Git     string `default:"git"`
	GitSign bool   `default:"true"`
	GitRepo string
	Debug   bool
}

var Config config
