package times

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Ts struct {
	BegTs    int
	EndTs    int
	BegDay   int
	EndDay   int
	NowDay   int
	SpanDay  int
	BegMonth int
	EndMonth int
}

// Ts2Day 时间戳转日期(年月日->20060102)
func Ts2Day(timestamp int64) int {
	day, _ := strconv.Atoi(time.Unix(timestamp, 0).Format("20060102"))
	return day
}

// Ts2Month 时间戳转日期(年月->200601)
func Ts2Month(timestamp int64) int {
	month, _ := strconv.Atoi(time.Unix(timestamp, 0).Format("200601"))
	return month
}

// InitDailyTs 获取传入参数(时间)前一天的 Ts 信息
func InitDailyTs(t time.Time) *Ts {
	begTime := time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, -1, 0, t.Location())
	return genTs(begTime, endTime, t)
}

// InitMonthTs 获取传入参数(时间)前一月的 Ts 信息
func InitMonthTs(t time.Time) *Ts {
	begTime := time.Date(t.Year(), t.Month()-1, 1, 0, 0, 0, 0, t.Location())
	endTime := time.Date(t.Year(), t.Month(), 1, 0, 0, -1, 0, t.Location())
	return genTs(begTime, endTime, t)
}

// InitCustomFormatTs 获取传入参数(时间戳)之间 "标准化" 的 Ts 信息
func InitCustomFormatTs(startTs, stopTs int64) *Ts {
	beg, end := time.Unix(startTs, 0), time.Unix(stopTs, 0)

	begTime := time.Date(beg.Year(), beg.Month(), beg.Day(), 0, 0, 0, 0, beg.Location())
	endTime := time.Date(end.Year(), end.Month(), end.Day()+1, 0, 0, -1, 0, beg.Location())
	nowTime := time.Now()
	return genTs(begTime, endTime, nowTime)
}

// InitCustomTs 获取传入参数(时间戳)之间的 Ts 信息
func InitCustomTs(startTs, stopTs int64) *Ts {
	begTime := time.Unix(startTs, 0)
	endTime := time.Unix(stopTs, 0)
	nowTime := time.Now()
	return genTs(begTime, endTime, nowTime)
}

func genTs(begTime, endTime, nowTime time.Time) *Ts {
	begDay, _ := strconv.Atoi(begTime.Format("20060102"))
	endDay, _ := strconv.Atoi(endTime.Format("20060102"))
	nowDay, _ := strconv.Atoi(nowTime.Format("20060102"))

	begMonth, _ := strconv.Atoi(begTime.Format("200601"))
	endMonth, _ := strconv.Atoi(endTime.Format("200601"))

	begTs := int(begTime.Unix())
	endTs := int(endTime.Unix())

	spanDay := (endTs + 1 - begTs) / Day

	return &Ts{
		BegTs:    begTs,
		EndTs:    endTs,
		BegMonth: begMonth,
		EndMonth: endMonth,
		BegDay:   begDay,
		EndDay:   endDay,
		NowDay:   nowDay,
		SpanDay:  spanDay,
	}
}

func (ts *Ts) String() string {
	b, err := json.Marshal(ts)
	if err != nil {
		return fmt.Sprintf("%+v", *ts)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *ts)
	}
	return out.String()
}
