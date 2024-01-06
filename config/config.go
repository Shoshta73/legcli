package config

import "gopkg.in/ini.v1"

type ConfigData struct {
	Fullname       string
	DefaultLicence string
}

func GetConfigFileContents(file string) (*ConfigData, error) {
	cd := ConfigData{
		Fullname:       "",
		DefaultLicence: "",
	}

	cfg, err := ini.Load(file)
	if err != nil {
		return &cd, err
	}

	section := cfg.Section("Default")
	cd.Fullname = section.Key("fullname").String()
	cd.DefaultLicence = section.Key("default_licence").String()

	return &cd, nil
}
