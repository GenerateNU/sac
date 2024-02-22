import requests
import random
import uuid
from argparse import ArgumentParser


user_agents = [
    'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36',
    'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36',
    'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36',
    'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36',
    'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36',
    'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
    'Mozilla/5.0 (Macintosh; Intel Mac OS X 13_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15',
]


def fetch_all_majors(url):
    """Fetches all majors by iterating through paginated results."""
    all_majors = []
    page = 1
    while True:
        paginated_url = f"{url}&page={page}"
        headers = {'User-Agent': random.choice(user_agents)}
        response = requests.get(paginated_url, headers=headers)

        if response.status_code == 200:
            data = response.json()
            if not data:
                break
            all_majors.extend(data)
            page += 1
        else:
            print(f"Error fetching page {page}: {response.status_code}")
            break

    return all_majors


def create_major_insertion_query(major_name):
    """Creates an SQL INSERT query for a given major."""
    major_id = str(uuid.uuid4())
    query = f"INSERT INTO majors (id, name) VALUES ('{major_id}', '{major_name}');"
    return query


def export_majors(majors, export_file):
    """Exports majors to a specified file. Creates file if it doesn't exist."""
    try:
        with open(export_file, "w") as file:
            file.write("BEGIN;\n")
            for major in majors:
                query = create_major_insertion_query(major['acf']['major'])
                file.write(query + '\n')
            file.write("COMMIT;\n")
    except FileNotFoundError:
        print(f"File not found: {export_file}")
    except Exception as e:
        print(f"Error exporting majors: {e}")


def main():
    parser = ArgumentParser(
        description="Scrape Northeastern University majors and export to SQL file.")
    parser.add_argument(
        "-e", "--export", help="Path to the export file", default="majors.sql", type=str)

    args = parser.parse_args()

    url = "https://admissions.northeastern.edu/wp-json/wp/v2/major?per_page=100&order=asc&orderby=title&_fields=id%2Clink%2Ctype%2Cacf.major%2Cacf.catalog_link%2Cacf.colleges"
    all_majors = fetch_all_majors(url)

    export_majors(all_majors, args.export)
    print(f"Exported {len(all_majors)} majors to {args.export}")


if __name__ == "__main__":
    main()
