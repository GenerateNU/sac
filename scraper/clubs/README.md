

<h1 align="center">SAC Club Scraper</h1>

<br />

<div align="center">
  <!-- Github Actions -->
  <a href="https://github.com/GenerateNU/sac/actions/workflows/club_scraper.yml">
    <img src="https://github.com/GenerateNU/sac/actions/workflows/club_scraper.yml/badge.svg"
      alt="Club Scraper Workflow Status" />
  </a>
</div>

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
