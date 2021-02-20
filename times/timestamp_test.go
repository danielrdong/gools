package times

import (
	"fmt"
	"testing"
	"time"
)

func TestTs(t *testing.T) {
	fmt.Println("===== test: init daily ts ================")
	fmt.Println(InitDailyTs(time.Now()))
	fmt.Println("===== test: init month ts ================")
	fmt.Println(InitMonthTs(time.Now()))
	fmt.Println("===== test: init custom ts ================")
	fmt.Println(InitCustomTs(1611072000, 1611218024))
	fmt.Println("===== test: init custom format ts ================")
	fmt.Println(InitCustomFormatTs(1611072000, 1611218024))
	fmt.Println("===== test: timestamp conv to day & month ================")
	fmt.Println(Ts2Day(1611218024), Ts2Month(1611218024))
}
