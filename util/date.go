package util

import (
	"fmt"
	"sftp/constant"
	"time"
)

// CreateProcessingDate 起動引数が存在しない場合、現在時刻の日にちを返す
func CreateProcessingDate(reqProcessingDate string) (d string, e error) {
	var t time.Time
	if reqProcessingDate == "" {
		// 現在時刻の日にちを返す
		utc, _ := time.LoadLocation(constant.LOCAL)
		now := time.Now().In(utc)
		t = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, utc)
	} else {
		parsedDate, err := time.Parse(constant.ReqProcessingDateFormat, reqProcessingDate)
		if err != nil {
			//　引数が存在しない日付だったエラーを返却する
			e = fmt.Errorf("'%s'は日付ではありません", reqProcessingDate)
			return
		}
		// reqProcessingDateをフォーマットして返す
		t = parsedDate
	}
	d = t.Format(constant.ProcessingDateFormat)
	return
}
