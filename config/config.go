package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production
	// user ...
	UserServiceHost string
	UserServicePort int

	// teacher ...
	TeacherServiceHost string
	TeacherServicePort int

	// student ...
	StudentServiceHost string
	StudentServicePort int

	// system ...
	SystemServiceHost string
	SystemServicePort int

	// finance ...
	FinanceServiceHost string
	FinanceServicePort int

	// course ...
	CourseServiceHost string
	CourseServicePort int

	// messenger ...
	MessengerServiceHost string
	MessengerServicePort int

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	// api-gateway idfs
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	// system-service idfs
	c.SystemServiceHost = cast.ToString(getOrReturnDefault("SYSTEM_SERVICE_HOST", "46.101.101.23"))
	c.SystemServicePort = cast.ToInt(getOrReturnDefault("SYSTEM_SERVICE_PORT", 50050))

	// user-service idfs
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "46.101.101.23"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 50051))

	// teacher-service idfs
	c.TeacherServiceHost = cast.ToString(getOrReturnDefault("TEACHER_SERVICE_HOST", "46.101.101.23"))
	c.TeacherServicePort = cast.ToInt(getOrReturnDefault("TEACHER_SERVICE_PORT", 50052))

	// student-service idfs
	c.StudentServiceHost = cast.ToString(getOrReturnDefault("STUDENT_SERVICE_HOST", "46.101.101.23"))
	c.StudentServicePort = cast.ToInt(getOrReturnDefault("STUDENT_SERVICE_PORT", 50053))

	// user-service idfs
	c.CourseServiceHost = cast.ToString(getOrReturnDefault("COURSE_SERVICE_HOST", "46.101.101.23"))
	c.CourseServicePort = cast.ToInt(getOrReturnDefault("COURSE_SERVICE_PORT", 50005))

	// finance-service idfs
	c.FinanceServiceHost = cast.ToString(getOrReturnDefault("FINANCE_SERVICE_HOST", "46.101.101.23"))
	c.FinanceServicePort = cast.ToInt(getOrReturnDefault("FINANCE_SERVICE_PORT", 50055))

	// finance-service idfs
	c.MessengerServiceHost = cast.ToString(getOrReturnDefault("MESSENGER_SERVICE_HOST", "46.101.101.23"))
	c.MessengerServicePort = cast.ToInt(getOrReturnDefault("MESSENGER_SERVICE_PORT", 50054))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
