package log4j

type Level uint8

const (
	DEBUG Level = iota + 1
	INFO
	WARN
	ERROR
	FATAL
)

var levelMap = [...]string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
}

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "[ DEBUG ]"
	case INFO:
		return "[ INFO  ]"
	case WARN:
		return "[ WARIN ]"
	case ERROR:
		return "[ ERROR ]"
	case FATAL:
		return "[ FATAL ]"
	default:
		return "[ UNKNOW ]"
	}
}
