package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"shrink-url/internal/util"
)

type ShrinkUrlSettings struct {
	BaseUrl   string `json:"BaseUrl" validate:"url"`
	MaxLength int    `json:"MaxLength"`
}

type Config struct {
	ShrinkUrlSettings ShrinkUrlSettings `json:"ShrinkUrlSettings"`
	AllowedHosts      string            `json:"AllowedHosts"`
}

const (
	SettingsFile = "appsettings.json"
)

var Cfg = Config{
	ShrinkUrlSettings: ShrinkUrlSettings{
		BaseUrl:   "https://example.co/",
		MaxLength: 6,
	},
	AllowedHosts: "*",
}

func LoadConfig() {
	jsonFile, err := os.Open(SettingsFile)
	if err != nil {
		log.Printf("failed to load %s: %v\n", SettingsFile, err)
		return
	}
	log.Printf("Successfully loaded %s\n", SettingsFile)
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var cfg Config

	json.Unmarshal(byteValue, &cfg)

	if cfg.AllowedHosts != "" {
		Cfg.AllowedHosts = cfg.AllowedHosts
	}

	if util.ValidateUrl(cfg.ShrinkUrlSettings.BaseUrl) {
		Cfg.ShrinkUrlSettings.BaseUrl = cfg.ShrinkUrlSettings.BaseUrl
	}

	if cfg.ShrinkUrlSettings.MaxLength != 0 {
		Cfg.ShrinkUrlSettings.MaxLength = cfg.ShrinkUrlSettings.MaxLength
	}
}
