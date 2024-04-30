import json
from uuid import uuid4
import os
import argparse
import re
import datetime
import random

from progress.bar import IncrementalBar as Bar 
from openai import OpenAI
import requests

# STEP 0: Initialize constants + utility/helper functions
parser = argparse.ArgumentParser(prog="Mock data generator",
                                 description="Generates mock club, category, and tag data for testing/development.")
parser.add_argument("output_file", help="The SQL file to output the data to.")
parser.add_argument("-a", "--api-key", default=os.environ.get("OPENAI_API_KEY", ""), help="The OpenAI API key to use for generating descriptions and tags for clubs. (default: key in OPENAI_API_KEY or '' if that doesn't exist.)")
parser.add_argument("-t", "--top-n", default=1024, help="The max # of clubs to retrieve from NUEngage. (default: 1024)")
parser.add_argument("-p", "--parent_uuid", default="00000000-0000-0000-0000-000000000000", help="The UUID of the parent club for the generated clubs. (default: 00000000-0000-0000-0000-000000000000)")
args = parser.parse_args()

openai_client = OpenAI(api_key=args.api_key)
OUTPUT_FILE = args.output_file
TOP_N = args.top_n
PARENT_UUID = args.parent_uuid

MOCK_CATEGORIES_AND_TAGS = {
    "PreProfessional": [
        "Premed",
        "Prelaw",
    ],
    "CulturalAndIdentity": [
        "Judaism",
        "Christianity",
        "Hinduism",
        "Islam",
        "LatinAmerica",
        "AfricanAmerican",
        "AsianAmerican",
        "LGBTQ",
    ],
    "ArtsAndCreativity": [
        "PerformingArts",
        "VisualArts",
        "CreativeWriting",
        "Music",
    ],
    "SportsAndRecreation": [
        "Soccer",
        "Hiking",
        "Climbing",
        "Lacrosse",
    ],
    "ScienceAndTechnology": [
        "Mathematics",
        "Physics",
        "Biology",
        "Chemistry",
        "EnvironmentalScience",
        "Geology",
        "Neuroscience",
        "Psychology",
        "SoftwareEngineering",
        "ArtificialIntelligence",
        "DataScience",
        "MechanicalEngineering",
        "ElectricalEngineering",
        "IndustrialEngineering",
    ],
    "CommunityServiceAndAdvocacy": [
        "Volunteerism",
        "EnvironmentalAdvocacy",
        "HumanRights",
        "CommunityOutreach",
    ],
    "MediaAndCommunication": [
        "Journalism",
        "Broadcasting",
        "Film",
        "PublicRelations",
    ]
}

TAGS_STRING = "" # for tags_prompt
for key, value in MOCK_CATEGORIES_AND_TAGS.items():
    for tag in value:
        TAGS_STRING += f"\"{tag}\", "
TAGS_STRING = TAGS_STRING[:-2] # remove the last ", "

def description_prompt(name, preview):
    return f"""
    I need your help writing club descriptions to create some mock data for an app I'm working on. The club's name is \"{name}\", and its short summary is \"{preview}\".
    Generate a one to two paragraph description of what the club is and what it does. Use a mix of friendly, welcoming, and informative language and content. Store in this a JSON field \"description\".
    Your output should just be the JSON with the generated description, no other text. Thank you!
    """

def tags_prompt(name, description):
    return f"""
    I need your help picking club tags to create some mock data for an app I'm working on. The club's name is \"{name}\" and its description is \"{description}\".
    Here are the tags you can choose from, in a JSON-formatted array: [{TAGS_STRING}].
    From this list, please return an array in JSON-format of 1-8 tags you think would be appropriate from this list. Store this array in a JSON field called \"tags\".
    Your output should just be the JSON with the array, no other text. Thank you!
    """

# STEP 1: Generate all the data to write to mock.sql

# Get the clubs from nuengage, do some cleaning, and generate description (if needed) and tags (always needed)
clubs_response = requests.get(f"https://neu.campuslabs.com/engage/api/discovery/search/organizations?orderBy[0]=UpperName%20asc&top={TOP_N}&skip=0")
clubs_json = json.loads(clubs_response.text)['value']
clubs = []

bar = Bar('Making club data', max=len(clubs_json), suffix="%(remaining)d Remaining, ETA: %(eta_td)s")
bad_club_descriptions = 0
for club in clubs_json:
    name = club['Name'].replace("(Tentative) ", "")
    preview = club['Summary']
    description = club['Description']

    # Generate description for club if needed.
    if not isinstance(description, str):
        description_response = openai_client.chat.completions.create(
            model="gpt-3.5-turbo-0125",
            response_format={ "type": "json_object" },
            messages=[
                {"role": "system", "content": "You are a helpful assistant designed to output JSON."},
                {"role": "user", "content": description_prompt(name, preview)}
            ]
        )

        description = json.loads(description_response.choices[0].message.content)['description']
        # Some description responses will not contain valid JSON, hence this if check.
        if not isinstance(description, str):
            description = ""
            bad_club_descriptions += 1     
    else:
        description =  re.sub('<[^<]+?>', '', club['Description']).replace("&nbsp;", " ")

    tags_response = openai_client.chat.completions.create(
        model="gpt-3.5-turbo-0125",
        response_format={ "type": "json_object" },
        messages=[
            {"role": "system", "content": "You are a helpful assistant designed to output JSON."},
            {"role": "user", "content": tags_prompt(name, description)}
        ]
    )

    tags = json.loads(tags_response.choices[0].message.content)['tags']
    
    clubs.append({
        "id": str(uuid4()),
        "name": name,
        "preview": preview,
        "description": description,
        "num_members": random.randint(1,1024),
        "is_recruiting": random.choice(["TRUE", "FALSE"]),
        "recruitment_cycle": random.choice(["fall", "spring", "fallSpring", "always"]),
        "recruitment_type": random.choice(["unrestricted", "application"]),
        "tags": tags
    })

    bar.next()
bar.finish()
print(f"Bad club descriptions (JSON not created properly): {bad_club_descriptions}")

# Generate the UUIDS for Categories, Tags, and a table for tags to categories.
category_uuids = {}
for category in MOCK_CATEGORIES_AND_TAGS:
    category_uuids[category] = str(uuid4())

tag_uuids = {}
tags_to_categories = {}
for category, tags in MOCK_CATEGORIES_AND_TAGS.items():
    for tag in tags:
        tag_uuids[tag] = str(uuid4())
        tags_to_categories[tag] = category 

# STEP 2: Write to SQL file.
with open(OUTPUT_FILE, 'w', encoding='utf-8') as file:
    def w(x):
        file.write(f"{x}\n")

    # Header
    w("-- AUTO GENERATED MOCK DATA. DO NOT MODIFY")
    w(f"-- GENERATED AT {datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")}")

    # Categories
    w("BEGIN;")
    for category, uuid in category_uuids.items():
        w(f"""INSERT INTO "categories" ("id", "name") VALUES ('{uuid}', '{category}');""")
    
    # Tags
    for tag, uuid in tag_uuids.items():
        w(f"""INSERT INTO "tags" ("id", "name", "category_id") VALUES ('{uuid}', '{tag}', '{category_uuids[tags_to_categories[tag]]}');""")

    # Clubs
    for club in clubs:
        w("""INSERT INTO "clubs" ("id", "name", "preview", "description", "num_members", "is_recruiting", "recruitment_cycle", "recruitment_type", "parent") VALUES ('{}', '{}', '{}', '{}', '{}', '{}', '{}', '{}', '{}');""".format(
            club['id'],
            club['name'].replace("\'", "''"),
            club['preview'].replace("\'", "''"),
            club['description'].replace('\'', "''"),
            str(club['num_members']),
            club['is_recruiting'],
            club['recruitment_cycle'],
            club['recruitment_type'],
            PARENT_UUID
        ))
    
    # Tags <-> Clubs
    # FIXME (maybe)?: The ChatGPT tags prompt will return tags that aren't in the list, hence this if statement being here.
    # Don't think its a huge issue because ~half of tags are good. 
    bad_tags = 0
    total_tags = 0 
    for club in clubs:
        total_tags += len(club['tags'])
        for tag in club['tags']:
            if tag in tag_uuids:
                w(f"""INSERT INTO "club_tags" ("tag_id", "club_id") VALUES ('{tag_uuids[tag]}', '{club['id']}')""")
            else:
                bad_tags += 1

    print(f"{bad_tags}/{total_tags} marked for not existing in tag_uuids.")

    w("COMMIT;")

    w("-- END GENERATED MOCK DATA. Made with <3 by Garrett and Michael")