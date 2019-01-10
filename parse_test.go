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

        duration, phraseType, err := FormatToDuration(phrases[0])
        if err != nil {
            t.Fatalf("Could not parse (1): [%s] [%s] [%s]", durationUnitAbbreviation, phrases[0], err)
        } else if phraseType != PhraseTypeInterval {
            t.Fatalf("Phrase-type not interval-type as expected.")
        }

        if duration != durationUnit*1 {
            t.Fatalf("phrase does not parse (1): [%s]", durationUnitAbbreviation)
        }

        duration, phraseType, err = FormatToDuration(phrases[1])
        if err != nil {
            t.Fatalf("Could not parse (2): [%s] [%s] [%s]", durationUnitAbbreviation, phrases[0], err)
        } else if phraseType != PhraseTypeInterval {
            t.Fatalf("Phrase-type not interval-type as expected.")
        }

        if duration != durationUnit*11 {
            t.Fatalf("phrase does not parse (2): [%s]", durationUnitAbbreviation)
        }

        _, _, err = FormatToDuration(phrases[2])
        if err == nil || log.Is(err, ErrInvalidFormat) == false {
            t.Fatalf("phrase should not parse but does: [%s]", durationUnitAbbreviation)
        }
    }
}

func TestFormatToDuration_Polarities(t *testing.T) {
    duration, phraseType, err := FormatToDuration("-33us")
    if err != nil {
        t.Fatalf("Could not parse: [%s]", err)
    } else if phraseType != PhraseTypeInterval {
        t.Fatalf("Phrase-type not interval-type as expected.")
    }

    expected := time.Microsecond * 33 * -1
    if duration != expected {
        t.Fatalf("Parsed negative value not correct: (%d) != (%s)", duration, expected)
    }

    duration, phraseType, err = FormatToDuration("+33us")
    if err != nil {
        t.Fatalf("Could not parse: [%s]", err)
    } else if phraseType != PhraseTypeInterval {
        t.Fatalf("Phrase-type not interval-type as expected.")
    }

    expected = time.Microsecond * 33
    if duration != expected {
        t.Fatalf("Parsed positive value not correct: (%d) != (%s)", duration, expected)
    }
}

func TestHumanToDuration(t *testing.T) {
    tests := []struct {
        phrase           string
        phraseType       PhraseType
        expectedDuration time.Duration
    }{
        {"1 second ago", PhraseTypeTime, time.Second * 1 * -1},
        {"6 days ago", PhraseTypeTime, oneDay * 6 * -1},
        {"1 week ago", PhraseTypeTime, oneWeek * 1 * -1},
        {"2 weeks ago", PhraseTypeTime, oneWeek * 2 * -1},
        {"1 second from now", PhraseTypeTime, time.Second * 1},
        {"12 seconds from now", PhraseTypeTime, time.Second * 12},
        {"45 seconds from now", PhraseTypeTime, time.Second * 45},
        {"15 minutes from now", PhraseTypeTime, time.Minute * 15},
        {"2 hours from now", PhraseTypeTime, time.Hour * 2},
        {"21 hours from now", PhraseTypeTime, time.Hour * 21},
        {"1 day from now", PhraseTypeTime, oneDay * 1},
        {"2 days from now", PhraseTypeTime, oneDay * 2},
        {"3 days from now", PhraseTypeTime, oneDay * 3},
        {"2 weeks from now", PhraseTypeTime, oneWeek * 2},
        {"1 month from now", PhraseTypeTime, oneMonth * 1},
        {"1 year from now", PhraseTypeTime, oneYear * 1},
        {"2 years from now", PhraseTypeTime, oneYear * 2},
        {"30 seconds from now", PhraseTypeTime, time.Second * 30},
        {"120 minutes from now", PhraseTypeTime, time.Minute * 120},
        {"1260 minutes from now", PhraseTypeTime, time.Minute * 1260},
        {"1 week from now", PhraseTypeTime, oneWeek},
        {"4 weeks from now", PhraseTypeTime, oneWeek * 4},
        {"25 weeks from now", PhraseTypeTime, oneWeek * 25},
        {"12 months from now", PhraseTypeTime, oneMonth * 12},
        {"24 months from now", PhraseTypeTime, oneMonth * 24},
        {"now", PhraseTypeTime, 0},
        {"every 15 minutes", PhraseTypeInterval, time.Minute * 15},
    }

    for i, parameters := range tests {
        actualDuration, phraseType, err := HumanToDuration(parameters.phrase)
        if err != nil {
            t.Fatalf("Could not parse human phrase: (%d) [%s]", i, parameters.phrase)
        } else if phraseType != parameters.phraseType {
            t.Fatalf("Phrase-type not interval-type as expected: (%d) [%s]", i, parameters.phrase)
        }

        if actualDuration != parameters.expectedDuration {
            t.Fatalf("Actual duration does not match expected duration: (%d) != (%d)", actualDuration, parameters.expectedDuration)
        }
    }
}

func TestParseDuration(t *testing.T) {
    tests := []struct {
        phrase           string
        phraseType       PhraseType
        expectedDuration time.Duration
    }{
        {"1 second ago", PhraseTypeTime, time.Second * 1 * -1},
        {"6 days ago", PhraseTypeTime, oneDay * 6 * -1},
        {"now", PhraseTypeTime, 0},
        {"1 week ago", PhraseTypeTime, oneWeek * 1 * -1},
        {"11h", PhraseTypeInterval, time.Hour * 11},
        {"36W", PhraseTypeInterval, time.Hour * 24 * 7 * 36},
        {"1Y", PhraseTypeInterval, time.Hour * 24 * 365},
    }

    for i, parameters := range tests {
        actualDuration, phraseType, err := ParseDuration(parameters.phrase)
        if err != nil {
            t.Fatalf("Could not parse phrase: (%d) [%s]", i, parameters.phrase)
        } else if phraseType != parameters.phraseType {
            t.Fatalf("Phrase-type not interval-type as expected: (%d) [%s]", i, parameters.phrase)
        }

        if actualDuration != parameters.expectedDuration {
            t.Fatalf("Actual duration does not match expected duration: (%d) != (%d)", actualDuration, parameters.expectedDuration)
        }
    }
}

func ExampleParseDuration() {
    actualDuration, phraseType, err := ParseDuration("24 days from now")
    log.PanicIf(err)

    fmt.Printf("%d [%s]\n", actualDuration/time.Hour/24, phraseType)

    actualDuration, phraseType, err = ParseDuration("now")
    log.PanicIf(err)

    fmt.Printf("%d [%s]\n", actualDuration, phraseType)

    actualDuration, phraseType, err = ParseDuration("12m")
    log.PanicIf(err)

    fmt.Printf("%d [%s]\n", actualDuration/time.Minute, phraseType)

    actualDuration, phraseType, err = ParseDuration("every 6 hours")
    log.PanicIf(err)

    fmt.Printf("%d [%s]\n", actualDuration/time.Hour, phraseType)

    // Output:
    // 24 [time]
    // 0 [time]
    // 12 [interval]
    // 6 [interval]
}
