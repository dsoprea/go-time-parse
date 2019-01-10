package timeparse

import (
    "errors"
    "regexp"
    "strconv"
    "strings"
    "time"

    "github.com/dsoprea/go-logging"
)

const (
    oneDay   = time.Hour * 24
    oneWeek  = oneDay * 7
    oneMonth = oneDay * 31
    oneYear  = oneDay * 365
)

var (
    formatRe       = regexp.MustCompile(`^([+-])?([0-9]+)([a-zA-Z]+)$`)
    humanPast1Re   = regexp.MustCompile(`^([0-9]+) +([a-z]+) +ago$`)
    humanPast2Re   = regexp.MustCompile(`^a[n]? +([a-z]+) +ago$`)
    humanFuture1Re = regexp.MustCompile(`^([0-9]+) +([a-z]+) +from +now$`)
    humanFuture2Re = regexp.MustCompile(`^a[n]? +([a-z]+) +from +now$`)

    DurationMap = map[string]time.Duration{
        "nanosecond":   time.Nanosecond,
        "nanoseconds":  time.Nanosecond,
        "microsecond":  time.Microsecond,
        "microseconds": time.Microsecond,
        "millisecond":  time.Millisecond,
        "milliseconds": time.Millisecond,
        "second":       time.Second,
        "seconds":      time.Second,
        "minute":       time.Minute,
        "minutes":      time.Minute,
        "hour":         time.Hour,
        "hours":        time.Hour,
        "day":          oneDay,
        "days":         oneDay,
        "week":         oneWeek,
        "weeks":        oneWeek,
        "month":        oneMonth,
        "months":       oneMonth,
        "year":         oneYear,
        "years":        oneYear,
    }
)

var (
    ErrInvalidFormat = errors.New("invalid format")
)

func FormatToDuration(phrase string) (duration time.Duration, err error) {
    defer func() {
        if state := recover(); state != nil {
            err = log.Wrap(state.(error))
        }
    }()

    parts := formatRe.FindStringSubmatch(phrase)
    if parts == nil {
        log.Panic(ErrInvalidFormat)
    }

    polarityPhrase := parts[1]
    nPhrase := parts[2]
    durationUnitAbbreviation := parts[3]

    var durationUnit time.Duration
    switch durationUnitAbbreviation {
    case "ns":
        durationUnit = time.Nanosecond
    case "us":
        durationUnit = time.Microsecond
    case "ms":
        durationUnit = time.Millisecond
    case "s":
        durationUnit = time.Second
    case "m":
        durationUnit = time.Minute
    case "h":
        durationUnit = time.Hour
    case "D":
        durationUnit = oneDay
    case "W":
        durationUnit = oneWeek
    case "M":
        durationUnit = oneMonth
    case "Y":
        durationUnit = oneYear
    default:
        log.Panic(ErrInvalidFormat)
    }

    n, err := strconv.Atoi(nPhrase)
    if err != nil {
        log.Panic(ErrInvalidFormat)
    }

    if polarityPhrase == "-" {
        durationUnit *= -1
    }

    duration = durationUnit * time.Duration(n)
    return duration, nil
}

func HumanToDuration(phrase string) (duration time.Duration, err error) {
    defer func() {
        if state := recover(); state != nil {
            err = log.Wrap(state.(error))
        }
    }()

    phrase = strings.TrimSpace(phrase)
    phrase = strings.ToLower(phrase)

    parts := humanPast1Re.FindStringSubmatch(phrase)
    if parts != nil {
        nPhrase := parts[1]
        unitPhrase := parts[2]

        n, err := strconv.Atoi(nPhrase)
        if err != nil {
            log.Panic(ErrInvalidFormat)
        }

        durationUnit, found := DurationMap[unitPhrase]
        if found == false {
            log.Panic(ErrInvalidFormat)
        }

        duration = durationUnit * time.Duration(n) * -1
        return duration, nil
    }

    parts = humanPast2Re.FindStringSubmatch(phrase)
    if parts != nil {
        unitPhrase := parts[1]

        durationUnit, found := DurationMap[unitPhrase]
        if found == false {
            log.Panic(ErrInvalidFormat)
        }

        duration = durationUnit * 1 * -1
        return duration, nil
    }

    parts = humanFuture1Re.FindStringSubmatch(phrase)
    if parts != nil {
        nPhrase := parts[1]
        unitPhrase := parts[2]

        n, err := strconv.Atoi(nPhrase)
        if err != nil {
            log.Panic(ErrInvalidFormat)
        }

        durationUnit, found := DurationMap[unitPhrase]
        if found == false {
            log.Panic(ErrInvalidFormat)
        }

        duration = durationUnit * time.Duration(n)
        return duration, nil
    }

    parts = humanFuture2Re.FindStringSubmatch(phrase)
    if parts != nil {
        unitPhrase := parts[1]

        durationUnit, found := DurationMap[unitPhrase]
        if found == false {
            log.Panic(ErrInvalidFormat)
        }

        duration = durationUnit * 1
        return duration, nil
    }

    if phrase == "now" {
        duration = time.Duration(0)
        return duration, nil
    }

    log.Panic(ErrInvalidFormat)
    return duration, nil
}

func ParseDuration(phrase string) (duration time.Duration, err error) {
    duration, err = FormatToDuration(phrase)
    if err == nil {
        return duration, nil
    } else if log.Is(err, ErrInvalidFormat) == false {
        log.PanicIf(err)
    }

    duration, err = HumanToDuration(phrase)
    if err == nil {
        return duration, nil
    } else if log.Is(err, ErrInvalidFormat) == false {
        log.PanicIf(err)
    }

    log.PanicIf(err)
    return duration, nil
}
