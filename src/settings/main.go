package settings

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Settings holds the configuration values from the .env file
type Settings struct {
	Port     int    `env:"PORT,default=8080"`
	LogLevel string `env:"LOG_LEVEL,default=info"`

	UploadPath string `env:"UPLOAD_PATH,default=./videos/uploaded"`
	HlsPath    string `env:"HLS_PATH,default=./videos/hls_converted"`
}

var (
	SettingsIns *Settings
	err         error
)

func init() {
	SettingsIns, err = newSettings()
	if err != nil {
		// Handle error appropriately
		panic(err) // Replace with a suitable error handling mechanism
	}
}

func newSettings() (*Settings, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	// Create new Settings struct
	settings := &Settings{}

	// Get and parse environment variables
	v := reflect.ValueOf(settings).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("env")
		if tag == "" {
			continue
		}
		parts := strings.Split(tag, ",")
		envKey := parts[0]
		defaultValue := ""
		if len(parts) > 1 {
			defaultValue = parts[1]
		}
		if err := parseEnv(envKey, defaultValue, v.Field(i)); err != nil {
			return nil, err
		}
	}

	return settings, nil
}

// parseEnv handles parsing environment variables
func parseEnv(key string, defaultValue string, field reflect.Value) error {
	envValue, exists := os.LookupEnv(key)
	if !exists {
		envValue = defaultValue
	}

	// Parse the value based on the type of the field
	switch field.Kind() {
	case reflect.Int:
		intValue, err := strconv.Atoi(envValue)
		if err != nil {
			return fmt.Errorf("error parsing environment variable %s: %w", key, err)
		}
		field.SetInt(int64(intValue))
	case reflect.String:
		field.SetString(envValue)
	default:
		return fmt.Errorf("unsupported type for environment variable %s: %s", key, field.Kind())
	}

	return nil
}
