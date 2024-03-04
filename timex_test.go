package timex_test

import (
	"github.com/innotechdevops/timex"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	// Given
	currentTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// When
	actual := currentTime.Add(time.Duration(0) * time.Minute)

	// Then
	if actual.Minute() != 0 {
		t.Error("Error", actual.Minute())
	}
}

func TestUtcToGmt7(t *testing.T) {
	// Given
	datetime, _ := timex.ParseBy("2023-12-03T13:51:06.474Z", timex.DateTimeFormatISO)

	// When
	date := timex.UtcToGmt7(datetime)
	format := date.Format(timex.TimeFormatDash1)

	// Then
	if format != "2023-12-03 20:51:06" {
		t.Error("Error parse", datetime, "to", format)
	}
}

func TestParseByGMT7(t *testing.T) {
	// Given
	datetime := "2023-02-14 07:00:00"

	// When
	date, _ := timex.ParseByGMT7(datetime, timex.TimeFormatDash1)
	format := date.Format(timex.TimeFormatDash1)

	// Then
	if format != "2023-02-14 00:00:00" {
		t.Error("Error", format)
	}
}

func TestIsWorkdayBy(t *testing.T) {
	// Given
	date := "2023-02-14"

	// When
	w, _ := timex.IsWorkdayBy(date, timex.DateFormatDash)

	// Then
	if !w {
		t.Error("It's not a workday.")
	}
}

func TestFormatSlash(t *testing.T) {
	// Given
	d := time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC)

	// When
	f := d.Format(timex.DateFormatSlash2)

	// Then
	if f != "02/01/2022" {
		t.Error("Format fail!")
	}
}

func TestDateFormat(t *testing.T) {
	// Given
	d := time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC)

	// When
	f := d.Format(timex.DateFormat)

	// Then
	if f != "20220102" {
		t.Error("Format fail!")
	}
}

func TestTimeFormatDash2(t *testing.T) {
	// Given
	d := time.Date(2022, 1, 2, 10, 11, 12, 0, time.UTC)

	// When
	f := d.Format(timex.TimeFormatDash2)

	// Then
	if f != "2022-01-02 10:11" {
		t.Error("Format fail!")
	}
}

func TestDateFormatDash(t *testing.T) {
	// Given
	d := time.Date(2022, 1, 2, 10, 11, 12, 0, time.UTC)

	// When
	f := d.Format(timex.DateFormatDash)

	// Then
	if f != "2022-01-02" {
		t.Error("Format fail!")
	}
}

func TestParse(t *testing.T) {
	// Given
	dt := "3333-3-09T00:00:00.000Z"

	// When
	actual, err := timex.Parse(dt)

	// Then
	if err == nil && !actual.IsZero() {
		t.Error("Parse is not error, actual is", actual)
	}
}

func TestParseDdMmYyyyIsValid1(t *testing.T) {
	// Given
	d := "1/2/2022"

	// When
	actual, err := timex.ParseDdMmYyyy(d)

	// Then
	if err != nil || actual != "01/02/2022" {
		t.Error(err, actual)
	}
}

func TestParseDdMmYyyyIsValid2(t *testing.T) {
	// Given
	d := "10/2/2022"

	// When
	actual, err := timex.ParseDdMmYyyy(d)

	// Then
	if err != nil || actual != "10/02/2022" {
		t.Error(err, actual)
	}
}

func TestParseDdMmYyyyIsValid3(t *testing.T) {
	// Given
	d := "1/12/2022"

	// When
	actual, err := timex.ParseDdMmYyyy(d)

	// Then
	if err != nil || actual != "01/12/2022" {
		t.Error(err, actual)
	}
}

func TestParseDdMmYyyyIsDayInvalid1(t *testing.T) {
	// Given
	d := "0/2/2022"

	// When
	actual, err := timex.ParseDdMmYyyy(d)

	// Then
	if err == nil || actual != "" {
		t.Error("Success", actual)
	}
}

func TestParseDdMmYyyyIsDayInvalid2(t *testing.T) {
	// Given
	d := "/2/2022"

	// When
	actual, err := timex.ParseDdMmYyyy(d)

	// Then
	if err == nil || actual != "" {
		t.Error("Success", actual)
	}
}

func TestParseDdMmYyyyIsDayInvalid3(t *testing.T) {
	// Given
	d := "99/99/9999"

	// When
	actual, err := timex.ParseDdMmYyyy(d)

	// Then
	if err == nil || actual != "" {
		t.Error("Success", actual)
	}
}

func TestParseDdMmYyyyIsMonthInvalid1(t *testing.T) {
	// Given
	d := "10//2022"

	// When
	actual, err := timex.ParseDdMmYyyy(d)

	// Then
	if err == nil || actual != "" {
		t.Error("Success", actual)
	}
}

func TestParseDdMmYyyyIsMonthInvalid2(t *testing.T) {
	// Given
	d := "10/0/2022"

	// When
	actual, err := timex.ParseDdMmYyyy(d)

	// Then
	if err == nil || actual != "" {
		t.Error("Success", actual)
	}
}

func TestParseDdMmYyyyIsYearInvalid1(t *testing.T) {
	// Given
	d := "10/1/022"

	// When
	actual, err := timex.ParseDdMmYyyy(d)

	// Then
	if err == nil || actual != "" {
		t.Error("Success", actual)
	}
}

func TestConvertDdMmYyyyByYyyyMmDd(t *testing.T) {
	// Given
	d := "10/12/2022"

	// When
	actual := timex.ConvertDdMmYyyyBy(d, timex.DateFormatSlash1)

	// Then
	if actual != "2022/12/10" {
		t.Error("Error", actual)
	}
}

func TestEndOfMonth(t *testing.T) {
	currentDate, _ := timex.ParseBy("2023-05-01", timex.DateFormatDash)

	actual := timex.EndOfMonth(currentDate)

	if actual != 31 {
		t.Error("Error", actual)
	}
}
