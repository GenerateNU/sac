# SAC Club Scraper

A Rust CLI that scrapes Clubs from to be used as mock data for natural language search.

> [!NOTE]
> It is assumed you have Rust installed on your machine. If not, you can install it [here](https://www.rust-lang.org/tools/install).

## Usage

```console
Usage: sac_club_scraper [OPTIONS]

Options:
  -t, --top-n <TOP_N>    Top N results to parse [default: 1024]
  -o, --output <OUTPUT>  Output file [default: mock.sql]
  -p, --parent <PARENT>  Parent club UUID [default: 00000000-0000-0000-0000-000000000000]
  -h, --help             Print help
  -V, --version          Print version
```
