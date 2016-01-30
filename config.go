package venom

import (
	"os"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	File string `yaml:"file"`
	URL  string `yaml:"url" mapstructure:"url"`
}

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
}

func NewConfig() (*Config, error) {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	c := Config{}
	flags.StringVarP(&c.URL, "url", "u", "", "the url")
	viper.BindPFlag("url", flags.Lookup("url"))

	// This doesn't need to be in the Config struct, because we're just using it to override viper.
	file := flags.StringP("file", "f", "", "name of the config file")

	// Parse the command line args into the flag set, ignoring the command name.
	flags.Parse(os.Args[1:])

	if *file != "" {
		viper.SetConfigFile(*file)
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
