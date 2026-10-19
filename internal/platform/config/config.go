package config

import (
	"os"
	"strings"
	"time"
	_ "time/tzdata"
)

var (
	EnvLocale       string = "America/Sao_Paulo"
	EnvServerAddr   string = ":8080"
	EnvReportFolder string = "/tmp"
	EnvDatabaseUrl  string = "postgresql://postgres:postgres:localhost@5432/db"
)

func getEnv(key string) (string, bool) {
	if env, ok := os.LookupEnv(key); ok && strings.TrimSpace(env) != "" {
		return strings.TrimSpace(env), ok
	}

	return "", false
}

func Bootstrap() (err error) {
	if locale, ok := getEnv("LOCALE"); ok {
		EnvLocale = locale
	}

	if serverAddr, ok := getEnv("SERVER_ADDR"); ok {
		EnvServerAddr = serverAddr
	}

	if reportFolder, ok := getEnv("REPORT_FOLDER"); ok {
		EnvReportFolder = reportFolder
	}

	if databaseUrl, ok := getEnv("DATABASE_URL"); ok {
		EnvDatabaseUrl = databaseUrl
	}

	time.Local, err = time.LoadLocation(EnvLocale)

	return err
}
