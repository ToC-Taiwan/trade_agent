// Package config package config
package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"trade_agent/pkg/log"

	"gopkg.in/yaml.v2"
)

var (
	globalConfig *Config
	once         sync.Once
)

// Config Config
type Config struct {
	basePath string
	Server   Server   `json:"server,omitempty" yaml:"server"`
	Database Database `json:"database,omitempty" yaml:"database"`
	Schedule Schedule `json:"schedule,omitempty" yaml:"schedule"`
	MQTT     MQTT     `json:"mqtt,omitempty" yaml:"mqtt"`
	Trade    Trade    `json:"trade,omitempty" yaml:"trade"`
}

// Server Server
type Server struct {
	RunMode        string `json:"run_mode,omitempty" yaml:"run_mode"`
	HTTPPort       string `json:"http_port,omitempty" yaml:"http_port"`
	SinopacSRVHost string `json:"sinopac_srv_host,omitempty" yaml:"sinopac_srv_host"`
	SinopacSRVPort string `json:"sinopac_srv_port,omitempty" yaml:"sinopac_srv_port"`
}

// Database Database
type Database struct {
	DBHost     string `json:"db_host,omitempty" yaml:"db_host"`
	DBPort     string `json:"db_port,omitempty" yaml:"db_port"`
	DBUser     string `json:"db_user,omitempty" yaml:"db_user"`
	DBPass     string `json:"db_pass,omitempty" yaml:"db_pass"`
	Database   string `json:"database,omitempty" yaml:"database"`
	DBEncode   string `json:"db_encode,omitempty" yaml:"db_encode"`
	DBTimeZone string `json:"db_time_zone,omitempty" yaml:"db_time_zone"`
}

// MQTT MQTT
type MQTT struct {
	Host     string `json:"host,omitempty" yaml:"host"`
	Port     string `json:"port,omitempty" yaml:"port"`
	User     string `json:"user,omitempty" yaml:"user"`
	Password string `json:"password,omitempty" yaml:"password"`
	ClientID string `json:"client_id,omitempty" yaml:"client_id"`
	CAPath   string `json:"ca_path,omitempty" yaml:"ca_path"`
	CertPath string `json:"cert_path,omitempty" yaml:"cert_path"`
	KeyPath  string `json:"key_path,omitempty" yaml:"key_path"`
}

// Trade Trade
type Trade struct {
	HistoryPeriod   int64  `json:"history_period,omitempty" yaml:"history_period"`
	TargetCondition string `json:"target_condition,omitempty" yaml:"target_condition"`
	BlackStock      string `json:"black_stock,omitempty" yaml:"black_stock"`
	BlackCategory   string `json:"black_category,omitempty" yaml:"black_category"`
}

// Schedule Schedule
type Schedule struct {
	CleaneventCron     string `json:"cleanevent_cron,omitempty" yaml:"cleanevent_cron"`
	RestartSinopacCron string `json:"restart_sinopac_cron,omitempty" yaml:"restart_sinopac_cron"`
}

// parseConfig parseConfig
func parseConfig() {
	if globalConfig != nil {
		return
	}
	ex, err := os.Executable()
	if err != nil {
		log.Get().Panic(err)
	}
	exPath := filepath.Clean(filepath.Dir(ex) + "/configs/config.yaml")
	yamlFile, err := ioutil.ReadFile(exPath)
	if err != nil {
		log.Get().Panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &globalConfig)
	if err != nil {
		log.Get().Panic(err)
	}
	globalConfig.basePath = filepath.Clean(filepath.Dir(ex))
}

// Get Get
func Get() (config Config, err error) {
	if globalConfig != nil {
		return *globalConfig, err
	}
	once.Do(parseConfig)
	return *globalConfig, err
}

// GetServerConfig GetServerConfig
func (c Config) GetServerConfig() Server {
	return c.Server
}

// GetDBConfig GetDBConfig
func (c Config) GetDBConfig() Database {
	return c.Database
}

// GetScheduleConfig GetScheduleConfig
func (c Config) GetScheduleConfig() Schedule {
	return c.Schedule
}

// GetTradeConfig GetTradeConfig
func (c Config) GetTradeConfig() Trade {
	return c.Trade
}

// TargetCond TargetCond
type TargetCond struct {
	LimitPriceLow  float64 `json:"limit_price_low,omitempty" yaml:"limit_price_low"`
	LimitPriceHigh float64 `json:"limit_price_high,omitempty" yaml:"limit_price_high"`
	LimitVolume    int64   `json:"limit_volume,omitempty" yaml:"limit_volume"`
}

// GetTradeTargetCondtion GetTradeTargetCondtion
func (c Config) GetTradeTargetCondtion() *TargetCond {
	var cond TargetCond
	err := json.Unmarshal([]byte(c.Trade.TargetCondition), &cond)
	if err != nil {
		log.Get().Panic(err)
	}
	return &cond
}

// GetMQConfig GetMQConfig
func (c Config) GetMQConfig() MQTT {
	c.MQTT.CAPath = c.basePath + c.MQTT.CAPath
	c.MQTT.KeyPath = c.basePath + c.MQTT.KeyPath
	c.MQTT.CertPath = c.basePath + c.MQTT.CertPath
	return c.MQTT
}
