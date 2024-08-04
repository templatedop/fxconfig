package fxconfig

import (
	//"gotemplate/logger"
	//"os"
	"sync"

	//"github.com/spf13/viper"
)

var (
	once sync.Once
	//instance Econfig
)

// type Config struct {
// 	appName            string
// 	AppEnv             string
// 	DBConnection       string
// 	TokenSymmetricKey  string
// 	HttpUrl            string
// 	HttpPort           string
// 	DBHost             string
// 	DBPort             string
// 	DBdatabase         string
// 	DBUsername         string
// 	DBPassword         string
// 	TokenDuration      string
// 	RedisServer        string
// 	RedisPassword      string
// 	HttpAllowedOrigins string
// 	Loglevel           string
// 	ShutDownTime       string
// 	ShutDowntype       string

// 	MaxConns           int
// 	MinConns           int
// 	MaxConnLifetime    int
// 	MaxConnIdleTime    int
// 	HealthCheckPeriod  int
// 	HealthCheckTimeout int
// }

type Econfig struct {
	AppName            string `mapstructure:"AppName"`
	AppEnv             string `mapstructure:"AppEnv"`
	DBConnection       string `mapstructure:"DBConnection"`
	TokenSymmetricKey  string `mapstructure:"TokenSymmetricKey"`
	HttpUrl            string `mapstructure:"HttpUrl"`
	HttpPort           string `mapstructure:"HttpPort"`
	DBHost             string `mapstructure:"DBHost"`
	DBPort             string `mapstructure:"DBPort"`
	DBDatabase         string `mapstructure:"DBdatabase"`
	DBUsername         string `mapstructure:"DBUsername"`
	DBPassword         string `mapstructure:"DBPassword"`
	TokenDuration      string `mapstructure:"TokenDuration"`
	RedisServer        string `mapstructure:"RedisServer"`
	RedisPassword      string `mapstructure:"RedisPassword"`
	HttpAllowedOrigins string `mapstructure:"HttpAllowedOrigins"`
	LogLevel           string `mapstructure:"Loglevel"`
	ShutDownTime       string `mapstructure:"ShutDownTime"`
	ShutDowntype       string `mapstructure:"ShutDowntype"`
	MaxConns           int    `mapstructure:"MaxConns"`
	MinConns           int    `mapstructure:"MinConns"`
	MaxConnLifetime    int    `mapstructure:"MaxConnLifetime"`
	MaxConnIdleTime    int    `mapstructure:"MaxConnIdleTime"`
	HealthCheckPeriod  int    `mapstructure:"HealthCheckPeriod"`
	HealthCheckTimeout int    `mapstructure:"HealthCheckTimeout"`
}

// func NewConfig(c Config) Config {
// 	return Config{
// 		appName:            c.AppName,
// 		appEnv:             c.AppEnv,
// 		dBConnection:       c.DBConnection,
// 		tokenSymmetricKey:  c.TokenSymmetricKey,
// 		httpUrl:            c.HttpUrl,
// 		httpPort:           c.HttpPort,
// 		dBHost:             c.DBHost,
// 		dBPort:             c.DBPort,
// 		dBdatabase:         c.dBdatabase,
// 		dBUsername:         c.DBUsername,
// 		dBPassword:         c.DBPassword,
// 		tokenDuration:      c.TokenDuration,
// 		redisServer:        c.RedisServer,
// 		redisPassword:      c.RedisPassword,
// 		httpAllowedOrigins: c.HttpAllowedOrigins,
// 		loglevel:           c.loglevel,
// 		shutDownTime:       c.ShutDownTime,
// 		shutDowntype:       c.ShutDownType,
// 		maxConns:           c.MaxConns,
// 		minConns:           c.MinConns,
// 		maxConnLifetime:    c.MaxConnLifetime,
// 		maxConnIdleTime:    c.MaxConnIdleTime,
// 		healthcheckPeriod:  c.HealthCheckPeriod,
// 		healthcheckTimeout: c.HealthCheckTimeout,
// 	}
// }

// func Load(log *logger.Logger) Econfig {
// 	c := Econfig{}
// 	once.Do(func() {
// 		viper.SetConfigName("config")
// 		viper.AddConfigPath(".")
// 		viper.SetConfigType("yaml")

// 		// Attempt to read the configuration file
// 		if err := viper.ReadInConfig(); err != nil {
// 			log.Error("Failed to read configuration", err)
// 			os.Exit(1)
// 		}

// 		// Unmarshal the configuration into the struct
// 		if err := viper.Unmarshal(&c); err != nil {
// 			log.Error("Failed to unmarshal configuration", err)
// 			os.Exit(1)
// 		}
// 		//log.Debug("Config", c)
// 		//instance = NewConfig(c)

// 	})
// 	return c
// }

// func (c *Econfig) AppName() string {
// 	return c.appName

// }

// func (c *Econfig) AppEnv() string {
// 	return c.appEnv

// }
// func (c *Econfig) DBConnection() string {
// 	return c.dBConnection

// }

// func (c *Econfig) TokenSymmetricKey() string {
// 	return c.tokenSymmetricKey

// }

// func (c *Econfig) Dbhost() string {
// 	return c.dBHost

// }

// // HttpUrl returns the httpUrl field value.
// func (c *Econfig) HttpUrl() string {
// 	return c.httpUrl
// }

// // HttpPort returns the httpPort field value.
// func (c *Econfig) HttpPort() string {
// 	return c.httpPort
// }

// // DBHost returns the dBHost field value.
// func (c *Econfig) DBHost() string {
// 	return c.dBHost
// }

// // DBPort returns the dBPort field value.
// func (c *Econfig) DBPort() string {
// 	return c.dBPort
// }

// // DBDatabase returns the dBdatabase field value.
// func (c *Econfig) DBDatabase() string {
// 	return c.dBdatabase
// }

// // DBUsername returns the dBUsername field value.
// func (c *Econfig) DBUsername() string {
// 	return c.dBUsername
// }

// // DBPassword returns the dBPassword field value.
// func (c *Econfig) DBPassword() string {
// 	return c.dBPassword
// }

// // TokenDuration returns the tokenDuration field value.
// func (c *Econfig) TokenDuration() string {
// 	return c.tokenDuration
// }

// // RedisServer returns the redisServer field value.
// func (c *Econfig) RedisServer() string {
// 	return c.redisServer
// }

// // RedisPassword returns the redisPassword field value.
// func (c *Econfig) RedisPassword() string {
// 	return c.redisPassword
// }

// // HttpAllowedOrigins returns the httpAllowedOrigins field value.
// func (c *Econfig) HttpAllowedOrigins() string {
// 	return c.httpAllowedOrigins
// }

// // LogLevel returns the loglevel field value.
// func (c *Econfig) LogLevel() string {
// 	return c.loglevel
// }

// // ShutDownTime returns the shutDownTime field value.
// func (c *Econfig) ShutDownTime() string {
// 	return c.shutDownTime
// }

// // ShutDownType returns the shutDowntype field value.
// func (c *Econfig) ShutDownType() string {
// 	return c.shutDowntype
// }

// func (c *Econfig) MaxConns() int {
// 	return c.maxConns
// }

// func (c *Econfig) MinConns() int {
// 	return c.minConns
// }

// func (c *Econfig) MaxConnLifetime() int {
// 	return c.maxConnLifetime
// }

// func (c *Econfig) MaxConnIdleTime() int {
// 	return c.maxConnIdleTime
// }
// func (c *Econfig) HealthCheckPeriod() int {
// 	return c.healthcheckPeriod
// }
// func (c *Econfig) HealthCheckTimeout() int {
// 	return c.healthcheckTimeout
// }
