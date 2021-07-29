package log4j

import "testing"

func Test_GetGoroutines(t *testing.T) {
	if goroutinesNo := GetGoroutines(); goroutinesNo == -1{
		t.Errorf("err No : %d", goroutinesNo)
	}
	
}