package log4j

import (
	"testing"
	"time"
)


func Test_output(t *testing.T) {
	Debug("%s", "hellom world")
	Info("%s", "hellom world")
	Warin("%s", "hellom world")
	Error("%s", "hellom world")
	go func ()  {
		Fatal("%s", "hellom world")
	}()
	time.Sleep(time.Second)
}