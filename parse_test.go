package timeparse

import (
    "fmt"
    "testing"
    "time"

    "github.com/dsoprea/go-logging"
)

func TestFormatToDuration(t *testing.T) {
    parameters := map[string][3]string{
        "nanosecond":  {"1ns", "11ns", "ns"},
        "microsecond": {"1us", "11us", "us"},
        "millisecond": {"1ms", "11ms", "ms"},
        "second":      {"1s", "11s", "s"},
        "minute":      {"1m", "11m", "m"},
        "hour":        {"1h", "11h", "h"},
        "day":         {"1D", "11D", "D"},
        "week":        {"1W", "11W", "W"},
        "month":       {"1M", "11M", "M"},
        "year":        {"1Y", "11Y", "Y"},
    }

    for durationUnitAbbreviation, phrases := range parameters {
        durationUnit := DurationMap[durationUnitAbbreviation]

        duration, err := FormatToDuration(phrases[0])
        if err != nil {
            t.Fatalf("Could not parse (1): [%s] [%s] [%s]", durationUnitAbbreviation, phrases[0], err)
        }

        if duration != durationUnit*1 {
            t.Fatalf("phrase does not parse (1): [%s]", durationUnitAbbreviation)
        }

        duration, err = FormatToDuration(phrases[1])
        if err != nil {
            t.Fatalf("Could not parse (2): [%s] [%s] [%s]", durationUnitAbbreviation, phrases[0], err)
        }

        if duration != durationUnit*11 {
            t.Fatalf("phrase does not parse (2): [%s]", durationUnitAbbreviation)
        }

        _, err = FormatToDuration(phrases[2])
        if err == nil || log.Is(err, ErrInvalidFormat) == false {
            t.Fatalf("phrase should not parse but does: [%s]", durationUnitAbbreviation)
        }
    }
}

func TestFormatToDuration_Polarities(t *testing.T) {
    duration, err := FormatToDuration("-33us")
    if err != nil {
        t.Fatalf("Could not parse: [%s]", err)
    }

    expected := time.Microsecond * 33 * -1
    if duration != expected {
        t.Fatalf("Parsed negative value not correct: (%d) != (%s)", duration, expected)
    }

    duration, err = FormatToDuration("+33us")
    if err != nil {
        t.Fatalf("Could not parse: [%s]", err)
    }

    expected = time.Microsecond * 33
    if duration != expected {
        t.Fatalf("Parsed positive value not correct: (%d) != (%s)", duration, expected)
    }
}

func TestHumanToDuration(t *testing.T) {
    tests := []struct {
        phrase           string
        expectedDuration time.Duration
    }{
        {"1 second ago", time.Second * 1 * -1},
        {"6 days ago", oneDay * 6 * -1},
        {"now", 0},
        {"1 week ago", oneWeek * 1 * -1},
        {"2 weeks ago", oneWeek * 2 * -1},
        {"1 second from now", time.Second * 1},
        {"12 seconds from now", time.Second * 12},
        {"45 seconds from now", time.Second * 45},
        {"15 minutes from now", time.Minute * 15},
        {"2 hours from now", time.Hour * 2},
        {"21 hours from now", time.Hour * 21},
        {"1 day from now", oneDay * 1},
        {"2 days from now", oneDay * 2},
        {"3 days from now", oneDay * 3},
        {"2 weeks from now", oneWeek * 2},
        {"1 month from now", oneMonth * 1},
        {"1 year from now", oneYear * 1},
        {"2 years from now", oneYear * 2},
        {"30 seconds from now", time.Second * 30},
        {"120 minutes from now", time.Minute * 120},
        {"1260 minutes from now", time.Minute * 1260},
        {"1 week from now", oneWeek},
        {"4 weeks from now", oneWeek * 4},
        {"25 weeks from now", oneWeek * 25},
        {"12 months from now", oneMonth * 12},
        {"24 months from now", oneMonth * 24},
    }

    for i, parameters := range tests {
        actualDuration, err := HumanToDuration(parameters.phrase)
        if err != nil {
            t.Fatalf("Could not parse human phrase: (%d) [%s]", i, parameters.phrase)
        }

        if actualDuration != parameters.expectedDuration {
            t.Fatalf("Actual duration does not match expected duration: (%d) != (%d)", actualDuration, parameters.expectedDuration)
        }
    }
}

func TestParseDuration(t *testing.T) {
    tests := []struct {
        phrase           string
        expectedDuration time.Duration
    }{
        {"1 second ago", time.Second * 1 * -1},
        {"6 days ago", oneDay * 6 * -1},
        {"now", 0},
        {"1 week ago", oneWeek * 1 * -1},
        {"11h", time.Hour * 11},
        {"36W", time.Hour * 24 * 7 * 36},
    }

    for i, parameters := range tests {
        actualDuration, err := ParseDuration(parameters.phrase)
        if err != nil {
            t.Fatalf("Could not parse phrase: (%d) [%s]", i, parameters.phrase)
        }

        if actualDuration != parameters.expectedDuration {
            t.Fatalf("Actual duration does not match expected duration: (%d) != (%d)", actualDuration, parameters.expectedDuration)
        }
    }
}

func ExampleParseDuration() {
    actualDuration, err := ParseDuration("24 days from now")
    log.PanicIf(err)

    fmt.Printf("%d\n", actualDuration/time.Hour/24)

    actualDuration, err = ParseDuration("now")
    log.PanicIf(err)

    fmt.Printf("%d\n", actualDuration)

    actualDuration, err = ParseDuration("12m")
    log.PanicIf(err)

    fmt.Printf("%d\n", actualDuration/time.Minute)

    // Output:
    // 24
    // 0
    // 12
}
