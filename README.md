# sbb

CLI Timetable queries to the [opendata Transport API](https://transport.opendata.ch/)

## Installation

### Homebrew

    brew install mhutter/tap/sbb

### `go install`

    go install github.com/mhutter/sbb/...

### Manual installation

You can always download the latest binaries from the [releases page](https://github.com/mhutter/sbb/releases/latest).

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

### Keywords

| Option | Description |
| ------ | ----------- |
| from   | Specifies the departure location of the connection |
| to     | Specifies the arrival location of the connection |
| date   | Date of the connection |
| time   | Time of the connection |


---
> by [@mhutter](https://github.com/mhutter)
> λ [hutter.io](https://hutter.io/)
> λ twitter: [@Dratir](https://twitter.com/Dratir)
