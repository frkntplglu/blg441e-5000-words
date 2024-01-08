package config

import "github.com/frkntplglu/emir-backend/pkg/config"

var APP_NAME = config.GetEnv("APP_NAME", "5000_WORDS")
var APP_VERSION = config.GetEnv("APP_VERSION", "0.0.1")

var DB_HOST = config.GetEnv("DB_HOST", "mahmud.db.elephantsql.com")
var DB_PORT = config.GetEnvInt("DB_PORT", 5432)
var DB_NAME = config.GetEnv("DB_NAME", "qqcpfrps")
var DB_USER = config.GetEnv("DB_USER", "qqcpfrps")
var DB_PASSWORD = config.GetEnv("DB_PASSWORD", "nz47OOhXDvOTy2By1vNrCfWMqDgz-CRF")

var ACCESS_TOKEN = config.GetEnv("ACCESS_TOKEN", "MYACCESSKEY")
var REFRESH_TOKEN = config.GetEnv("REFRESH_TOKEN", "MYREFRESHKEY")
