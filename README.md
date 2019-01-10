[![Build Status](https://travis-ci.org/dsoprea/go-time-parse.svg?branch=master)](https://travis-ci.org/dsoprea/go-time-parse)
[![Coverage Status](https://coveralls.io/repos/github/dsoprea/go-time-parse/badge.svg?branch=master)](https://coveralls.io/github/dsoprea/go-time-parse?branch=master)
[![GoDoc](https://godoc.org/github.com/dsoprea/go-time-parse?status.svg)](https://godoc.org/github.com/dsoprea/go-time-parse)


# Overview

This project converts from human-friendly duration expressions to `time.Duration` quantities. 

It maintains reasonable compatibility with the time expressions produced by [go-humanize](https://github.com/dustin/go-humanize).


# Usage

There is one general function that'll try to decode a duration expression any way we can:

- `ParseDuration(phrase string)`

There are also two specific functions that can be called directly if you know in advance what format to expect:

- `FormatToDuration(phrase string)`
- `HumanToDuration(phrase string)`

These functions return both a `time.Duration` and a `PhraseType` indicating whether the phrase pointed to a specific time or just described an interval.


# Examples

See the [ParseDuration](https://godoc.org/github.com/dsoprea/go-time-parse#example-ParseDuration) testable example.


## Human-Style

"before" expressions:

- "1 nanosecond ago"
- "1 microsecond ago"
- "1 millisecond ago"
- "1 second ago"
- "1 minute ago"
- "1 month ago"
- "6 days ago"
- "1 week ago"
- "2 weeks ago"
- "2 months ago"
- "2 years ago"
- "a minute ago"

"after" expressions:

- "1 second from now"
- "45 seconds from now"
- "21 hours from now"
- "1 day from now"
- "2 weeks from now"
- "1 month from now"
- "2 years from now"
- "a minute from now"

All of the above are interpreted as time-type phrases.

In addition, the phrase "now" will be parsed as a zero duration with a "time" phrase-type.

"every" expressions can also be provided and will always be interpreted as interval-type phrases:

- "every 6 hours"


## Format-Style

More concise expressions can be used to describe quantities, and may have an optional polarity to [explicitly] express positive/negative durations:

- "12ns"
- "-23us"
- "+34ms"
- "45s"
- "56m"
- "67h"
- "78D"
- "89W"
- "90M"
- "1Y"

These expressions are always interpreted as interval-type phrases.
