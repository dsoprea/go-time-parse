# Overview

This project converts from human-friendly duration expressions to `time.Duration` quantities. 

It maintains reasonable compatibility with the time expressions produced by [go-humanize](https://github.com/dustin/go-humanize).


# Usage

There is one general function that'll try to decode a duration expression any way we can:

- `ParseDuration(phrase string)`

There are also two specific functions that can be called directly if you know in advance what format to expect:

- `timeparse.FormatToDuration(phrase string)`
- `timeparse.HumanToDuration(phrase string)`


# Examples

## Human-Style

"before" expressions:

- "1 nanosecond ago": time.Nanosecond * 1 * -1
- "1 microsecond ago": time.Microsecond * 1 * -1
- "1 millisecond ago": time.Millisecond * 1 * -1
- "1 second ago": time.Second * 1 * -1
- "1 minute ago": time.Minute * 1 * -1
- "1 month ago": time.Hour * 24 * 31 * 1 * -1
- "6 days ago": time.Hour * 24 * 6 * -1
- "1 week ago": time.Hour * 24 * 7 * 1 * -1
- "2 weeks ago": time.Hour * 24 * 7 * 2 * -1
- "2 months ago": time.Hour * 24 * 31 * 2 * -1
- "2 years ago": time.Hour * 24 * 365 * 2 * -1
- "a minute ago": time.Minute * 1 * -1

"after" expressions:

- "1 second from now": time.Second * 1
- "45 seconds from now": time.Second * 45
- "21 hours from now": time.Hour * 21
- "1 day from now": time.Hour * 24 * 1
- "2 weeks from now": time.Hour * 24 * 7 * 2
- "1 month from now": time.Hour * 24 * 31 * 1
- "2 years from now": time.Hour * 24 * 365 * 2
- "a minute from now": time.Minute * 1

Misc:

- "now": 0


## Format-Style

More concise expressions can be used to describe quantities, and may have a polarity to express negative durations:

- "12ns": time.Nanosecond * 12
- "-23us": time.Microsecond * 23 * -1
- "+34ms": time.Millisecond * 34
- "45s": time.Second * 45
- "56m": time.Minute * 56
- "67h": time.Hour * 67
- "78D": time.Hour * 24 * 78
- "89W": time.Hour * 24 * 7 * 89
- "90M": time.Hour * 24 * 31 * 90
- "1Y": time.Hour * 365
