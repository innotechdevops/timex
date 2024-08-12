package timex

import (
	"fmt"
	"github.com/jinzhu/now"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
	_ "time/tzdata"
)

const (
	Day                 = "day"
	Week                = "week"
	Month               = "month"
	Year                = "year"
	Custom              = "custom"
	PeriodPrevious      = "previous"
	PeriodCurrent       = "current"
	TimeZoneAsiaBangkok = "Asia/Bangkok"
	TimeFormatDash1     = "2006-01-02 15:04:05"
	TimeFormatSlash     = "2006/01/02 15:04:05"
	TimeFormatDash2     = "2006-01-02 15:04"
	DateFormatSlash1    = "2006/01/02"
	DateFormatSlash2    = "02/01/2006"
	DateFormatSlash3    = "2/1/2006"
	DateFormatSlash4    = "01/02/2006 15:04:05"
	DateFormatDash      = "2006-01-02"
	DateFormat          = "20060102"
	DateFormatTime      = "15:04:05"
	DateFormatMin       = "15:04"
	DateFormatTime2     = "200601021504"
	DateTimeFormatISO   = "2006-01-02T15:04:05.000Z"
	DateTimeFormatUTC   = "2006-01-02 15:04:05 +0000 UTC"
	DateTimeFormatUTC2  = "2006-01-02 15:04:05 -0700 -07"
)

// MonthsToDays Function to convert months to days considering leap years
func MonthsToDays(months int) int {
	days := 0
	for m := 1; m <= months; m++ {
		var daysInMonth int
		switch time.Month(m) {
		case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
			daysInMonth = 31
		case time.April, time.June, time.September, time.November:
			daysInMonth = 30
		case time.February:
			year := time.Now().Year() + months/12
			if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
				daysInMonth = 29 // Leap year
			} else {
				daysInMonth = 28
			}
		}
		days += daysInMonth
	}
	return days
}

func ParseByYyyyMm(month string, format string) time.Time {
	m := fmt.Sprintf("%s-01", month)
	t, _ := ParseBy(m, format)
	return t
}

func IsYyyyMm(month string, format string) bool {
	m := fmt.Sprintf("%s-01", month)
	_, e := ParseBy(m, format)
	return e == nil
}

func IsWorkdayBy(date string, layout string) (bool, error) {
	t, err := ParseBy(date, layout)
	if err != nil {
		return false, err
	}

	return IsWorkday(t), nil
}

func IsWorkday(date time.Time) bool {
	weekday := date.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}
	return true
}

func GetWeekday(date time.Time) string {
	return date.Weekday().String()
}

// ParseStartOfDay date string 2021-09-09 to time.Time
func ParseStartOfDay(date string) (time.Time, error) {
	startTime := fmt.Sprintf("%sT00:00:00.000Z", date)
	return Parse(startTime)
}

// ParseEndOfDay date string 2021-09-09 to time.Time
func ParseEndOfDay(date string) (time.Time, error) {
	endTime := fmt.Sprintf("%sT23:59:59.999Z", date)
	return Parse(endTime)
}

// Parse date string 2021-09-09T00:00:00.000Z to time.Time
func Parse(date string) (time.Time, error) {
	return time.Parse(DateTimeFormatISO, date)
}

func ParseBy(date string, layout string) (time.Time, error) {
	return time.Parse(layout, date)
}

func ParseByYmDash(ym string) (time.Time, error) {
	date := fmt.Sprintf("%s-01", ym)
	return ParseBy(date, DateFormatDash)
}

func ParseByGMT7(date string, layout string, utc ...bool) (time.Time, error) {
	return ParseByLocation(date, layout, TimeZoneAsiaBangkok, utc...)
}

func ParseByLocation(date string, layout string, tz string, utc ...bool) (time.Time, error) {
	t, err := time.ParseInLocation(layout, date, GetTimeZone(tz))
	if err == nil {
		if len(utc) > 0 {
			if !utc[0] {
				return t, err
			}
		}
		return t.UTC(), err
	}
	return t, err
}

func ParseDdMmYyyy(date string) (string, error) {
	if len(date) == 0 {
		return "", nil
	}
	err := errors.New("The date is in an invalid format")

	ddMmYyyy := strings.Split(date, "/")
	if len(ddMmYyyy) == 3 {
		dd, erd := strconv.Atoi(ddMmYyyy[0])
		mm, erm := strconv.Atoi(ddMmYyyy[1])
		yyyy, ery := strconv.Atoi(ddMmYyyy[2])

		if erd != nil || dd == 0 {
			return "", err
		}
		if erm != nil || mm == 0 {
			return "", err
		}
		if ery != nil || yyyy == 0 {
			return "", err
		}

		fillZero := func(n int, digit int) string {
			out := ""
			if n < 10 {
				out = fmt.Sprintf("0%d", n)
			} else {
				out = fmt.Sprint(n)
			}

			if len(out) == digit {
				return out
			}
			return ""
		}

		yyyyFilled := fillZero(yyyy, 4)
		if len(yyyyFilled) < 4 {
			return "", err
		}
		out := fmt.Sprintf("%s/%s/%s", fillZero(dd, 2), fillZero(mm, 2), yyyyFilled)

		// Validate date format
		t, tErr := ParseBy(out, DateFormatSlash2)
		if tErr != nil || t.IsZero() {
			return "", err
		}
		return out, nil
	}
	return "", err
}

func CalcEndOfMonth(date time.Time) time.Time {
	lastDate := now.With(date)
	return lastDate.EndOfMonth()
}

func ConvertDdMmYyyyBy(date string, layout string) string {
	if date == "" {
		return ""
	}
	t, err := ParseBy(date, DateFormatSlash2)
	if err != nil {
		return ""
	}
	return t.Format(layout)
}

func GetTimeZone(zone string) *time.Location {
	tz, err := time.LoadLocation(zone)
	timeZone := time.Local
	if err == nil {
		timeZone = tz
	}
	return timeZone
}

func TimeNowFormat(zone string, format string) string {
	timeZone := GetTimeZone(zone)
	return time.Now().In(timeZone).Format(format)
}

func UtcToGmt7(utcTime time.Time) time.Time {
	timeZone := GetTimeZone(TimeZoneAsiaBangkok)
	return utcTime.In(timeZone)
}

func PrevDay(num int) time.Time {
	return Now().AddDate(0, 0, -num)
}

func PrevMonth(num int) time.Time {
	t := Now()
	return SubMonth(t, num)
}

func SubMonth(t time.Time, num int) time.Time {
	sub := time.Date(t.Year(), (t.Month()+1)-time.Month(num), 0, 0, 00, 00, 00, t.Location())
	return sub
}

func PrevYear(num int) time.Time {
	return Now().AddDate(-num, 0, 0)
}

func GetFebruaryLastOfMonth() {
	month := 2
	feb := time.Date(2016, time.Month(month+1), 0, 0, 0, 0, 0, time.Local)
	fmt.Println(feb.Day()) // 29 days
}

func Now() time.Time {
	timeZone := GetTimeZone(TimeZoneAsiaBangkok)
	return time.Now().In(timeZone)
}

func NowWithoutTime() time.Time {
	n := Now()
	timeZone := GetTimeZone(TimeZoneAsiaBangkok)
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, timeZone)
}

func Date() (int, time.Month, int) {
	return Now().Date()
}

func TimeNow() *now.Now {
	return now.With(Now())
}

func FromTimestampGMT7(timestamp int64) time.Time {
	return FromTimestamp(timestamp, GetTimeZone(TimeZoneAsiaBangkok))
}

func FromTimestamp(timestamp int64, tz *time.Location) time.Time {
	return time.Unix(timestamp, 0).In(tz)
}

func Format(year int, m time.Month, d int) string {
	month := fmt.Sprintf("%d", m)
	day := fmt.Sprintf("%d", d)
	if m < 9 {
		month = fmt.Sprintf("0%d", m)
	}
	if d < 9 {
		day = fmt.Sprintf("0%d", d)
	}
	return fmt.Sprintf("%d-%s-%s", year, month, day)
}

func EndOfMonth(date time.Time) int {
	n := now.With(date)
	return n.EndOfMonth().Day()
}

func NextDay(day int) time.Time {
	return Now().Add((24 * time.Hour) * time.Duration(day))
}

func SubDay(currentDate time.Time, day int) time.Time {
	return time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.Location()).AddDate(0, 0, -day)
}

func Validate(date string, layout string) bool {
	_, err := time.Parse(layout, date)
	return err == nil
}

func ValidateDateFormatSlash3(dateTime string) bool {
	ot := strings.Split(dateTime, "-")
	for _, date := range ot {
		if len(date) == 0 {
			continue
		}
		if !Validate(date, DateFormatSlash3) {
			return false
		}
	}
	return true
}

func ValidateDateFormatDash(dateTime string) bool {
	if !Validate(dateTime, DateFormatDash) {
		return false
	}
	return true
}

func IsIntervalElapsed(location string, timestamp time.Time, updatedAt time.Time, interval int64) (time.Time, bool) {
	loc, _ := time.LoadLocation(location)
	notifyTimestamp := timestamp.In(loc)
	updateAtAdd := updatedAt.Add(time.Duration(interval) * time.Minute)
	return notifyTimestamp, notifyTimestamp.Unix() >= updateAtAdd.Unix()
}
