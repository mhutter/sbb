# sbb

CLI Timetable queries to the [opendata Transport API](https://transport.opendata.ch/)

## Installation

    go install github.com/mhutter/sbb/...

## Usage

    $ sbb from Zurich to Brig
    ---------------------------------------------------------
    Zürich HB                      dep: 06:02     dur: 02:09
    Brig                           arr: 08:11
    ---------------------------------------------------------
    Zürich HB                      dep: 06:32     dur: 02:08
    Brig                           arr: 08:40
    ---------------------------------------------------------
    Zürich HB                      dep: 07:02     dur: 02:09
    Brig                           arr: 09:11
    ---------------------------------------------------------
    Zürich HB                      dep: 08:02     dur: 02:09
    Brig                           arr: 10:11
    ---------------------------------------------------------

### Options
| Option | Description |
| ------ | ----------- |
| from   | Specifies the departure location of the connection |
| to     | Specifies the arrival location of the connection |
| date   | Date of the connection |
| time   | Time of the connection |


[![CircleCI](https://circleci.com/gh/mhutter/sbb.svg?style=svg)](https://circleci.com/gh/mhutter/sbb)
[![codecov](https://codecov.io/gh/mhutter/sbb/branch/master/graph/badge.svg)](https://codecov.io/gh/mhutter/sbb)
[![Go Report Card](https://goreportcard.com/badge/github.com/mhutter/sbb)](https://goreportcard.com/report/github.com/mhutter/sbb)

---
> by [@mhutter](https://github.com/mhutter)
> λ [hutter.io](https://hutter.io/)
> λ twitter: [@Dratir](https://twitter.com/Dratir)
