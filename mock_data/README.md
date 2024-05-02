# SAC Club Scraper

A Python script that creates mock club, tag, and category data for development and testing.

## Installation and Usage
This script was last developed using Python 3.12.2, and assumes you have Python installed.

```console
pip install -r requirements.txt
python main.py mock.sql
```

```console
usage: Mock data generator [-h] [-a API_KEY] [-t TOP_N] [-p PARENT_UUID] output_file

Generates mock club, category, and tag data for testing/development.

positional arguments:
  output_file           The SQL file to output the data to.

options:
  -h, --help            show this help message and exit
  -a API_KEY, --api-key API_KEY
                        The OpenAI API key to use for generating descriptions and tags for clubs. (default: key in OPENAI_API_KEY or '' if that doesn't exist.)
  -t TOP_N, --top-n TOP_N
                        The max # of clubs to retrieve from NUEngage. (default: 1024)
  -p PARENT_UUID, --parent_uuid PARENT_UUID
                        The UUID of the parent club for the generated clubs. (default: 00000000-0000-0000-0000-000000000000)
```
