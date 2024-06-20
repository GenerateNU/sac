import json
from uuid import uuid4
import os
import argparse
import re
import datetime
import random

from progress.bar import IncrementalBar as Bar 
from openai import OpenAI
import psycopg2
import requests

# STEP 0: Initialize constants + utility/helper functions
parser = argparse.ArgumentParser(prog="Mock data generator",
                                 description="Generates mock club, category, and tag data for testing/development.")
parser.add_argument("-s", "--skip_generation", action='store_true', help="Skip mock data generation. Instead, will look for a MOCK_DATA_OUTPUT.json in the CWD and use that to write to the DB.")
parser.add_argument("-a", "--api_key", default=os.environ.get("OPENAI_API_KEY", ""), help="The OpenAI API key to use for generating descriptions and tags for clubs. (default: key in OPENAI_API_KEY or '' if that doesn't exist.)")
parser.add_argument("-t", "--top-n", default=1024, help="The max # of clubs to retrieve from NUEngage. (default: 1024)")
parser.add_argument("-p", "--parent_uuid", default="00000000-0000-0000-0000-000000000000", help="The UUID of the parent club for the generated clubs. (default: 00000000-0000-0000-0000-000000000000)")

args = parser.parse_args()

openai_client = OpenAI(api_key=args.api_key)
TOP_N = args.top_n
PARENT_UUID = args.parent_uuid

MOCK_CATEGORIES_AND_TAGS = {
    "Pre-Professional": [
        "Premed",
        "Prelaw",
    ],
    "Cultural And Identity": [
        "Judaism",
        "Christianity",
        "Hinduism",
        "Islam",
        "Latin America",
        "African American",
        "Asian American",
        "LGBTQ",
    ],
    "Arts And Creativity": [
        "Performing Arts",
        "Visual Arts",
        "Creative Writing",
        "Music",
    ],
    "Sports And Recreation": [
        "Soccer",
        "Hiking",
        "Climbing",
        "Lacrosse",
    ],
    "Science And Technology": [
        "Mathematics",
        "Physics",
        "Biology",
        "Chemistry",
        "Environmental Science",
        "Geology",
        "Neuroscience",
        "Psychology",
        "Software Engineering",
        "Artificial Intelligence",
        "Data Science",
        "Mechanical Engineering",
        "Electrical Engineering",
        "Industrial Engineering",
    ],
    "Community Service And Advocacy": [
        "Volunteerism",
        "Environmental Advocacy",
        "HumanRights",
        "Community Outreach",
    ],
    "Media And Communication": [
        "Journalism",
        "Broadcasting",
        "Film",
        "PublicRelations",
    ]
}

randomEventImageData = [
    'https://i.imgur.com/TM29405.png',
    'https://i.imgur.com/IfjHhrs.png',
    'https://i.imgur.com/TQQ7AEE.png',
    'https://i.imgur.com/Gp0Wwa8.png',
]

TAGS_STRING = ", ".join(f"\"{tag}\"" for tag in sum(MOCK_CATEGORIES_AND_TAGS.values(), [])) # for tags_prompt

def description_prompt(name: str, preview: str) -> str:
    return f"""
    I need your help writing club descriptions to create some mock data for an app I'm working on. The club's name is \"{name}\", and its short summary is \"{preview}\".
    Generate a one to two paragraph description of what the club is and what it does. Use a mix of friendly, welcoming, and informative language and content. Store in this a JSON field \"description\".
    Your output should just be the JSON with the generated description, no other text. Thank you!
    """

def tags_prompt(name: str, description: str) -> str:
    return f"""
    I need your help picking club tags to create some mock data for an app I'm working on. The club's name is \"{name}\" and its description is \"{description}\".
    Here are the tags you can choose from, in a JSON-formatted array: [{TAGS_STRING}].
    From this list, please return an array in JSON-format of 1-8 tags you think would be appropriate from this list. Store this array in a JSON field called \"tags\".
    Your output should just be the JSON with the array, no other text. Thank you!
    """

def events_prompt(name: str, description: str) -> str:
    return f"""
    I need your help writing event descriptions to create some mock data for an app I'm working on. The club's name is \"{name}\", and its description is \"{description}\".
    Generate a one paragraph description of an event this club cluld have. Use a mix of friendly, welcoming, and informative language and content. Store in this a JSON field \"event\".
    Your output should just be the JSON with the generated description, no other text. Thank you!
    """

def random_start_and_endtimes() -> tuple[datetime.datetime, datetime.datetime]:
    year = random.randint(2024,2026)
    month = random.randint(1,12)
    day = random.randint(1,28)
    hour = random.randint(12,23)
    minute = random.choice([0,15,30])

    extra_hours = time_to_add = datetime.timedelta(hours=random.randint(1,4))

    start_time = datetime.datetime(year, month, day, hour, minute)
    end_time = start_time + extra_hours

    return (start_time, end_time)

# STEP 1: Generate mock data.
if not args.skip_generation:
    print("Skip generation is false, making mock data...")
    # Get the clubs from nuengage, do some cleaning, and generate description (if needed) and tags (always needed)
    print("Making NUEngage request...")
    clubs_response = requests.get(f"https://neu.campuslabs.com/engage/api/discovery/search/organizations?orderBy[0]=UpperName%20asc&top={TOP_N}&skip=0")
    print("NUEngage request done.")

    clubs_json = json.loads(clubs_response.text)['value']
    print(clubs_json)
    clubs = []
    events = []

    bar = Bar('Making club data', max=len(clubs_json), suffix="%(remaining)d Remaining, ETA: %(eta_td)s")
    bad_club_descriptions = 0
    for club in clubs_json:
        name = club['Name'].replace("(Tentative) ", "").replace("( Tentative) ", "")
        logo = "https://se-images.campuslabs.com/clink/images/{}".format(club['ProfilePicture'])
        preview = club['Summary'][:250]
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

            print('Created description for club:', name)

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

        logoOptions = [

        ]
        
        clubData = {
            "id": str(uuid4()),
            "name": name,
            "preview": preview,
            "description": description,
            "num_members": random.randint(1,1024),
            "is_recruiting": random.choice(["TRUE", "FALSE"]),
            "recruitment_cycle": random.choice(["fall", "spring", "fallSpring", "always"]),
            "recruitment_type": random.choice(["unrestricted", "application"]),
            "tags": tags,
            "logo": logo
        }

        clubs.append(clubData)

        iterations = random.randint(1,3)
        for i in range(0,iterations):
            event_response = openai_client.chat.completions.create(
                model="gpt-3.5-turbo-0125",
                response_format={ "type": "json_object" },
                messages=[
                    {"role": "system", "content": "You are a helpful assistant designed to output JSON."},
                    {"role": "user", "content": events_prompt(name, description)}
                ]
            )

            # Some descriptions will not be valid JSON
            event_description = json.loads(event_response.choices[0].message.content)['event']
            if not isinstance(event_description, str):
                event_description = "Event description"

            times = random_start_and_endtimes()
            
            event_type = random.choice(["hybrid", "in_person", "virtual"])
            link = ""
            if event_type == "hybrid" or event_type == "virtual":
                link = "https://northeastern.zoom.us/j/6610103275"
            
            location = random.choice(["West Village H", "Curry Student Center", "Snell Library", "Krentzman Quad", "Marino", "ISEC", "ISEC 2"])
            if event_type == "online":
                location = ""
            
            image = random.choice(randomEventImageData)

            events.append({
                "id": str(uuid4()),
                "club_id": clubData["id"],
                "name": f"Event for {clubData["name"]}",
                "image": image,
                "preview": "This club is holding an event.",
                "description": event_description,
                "event_type": event_type,
                "start_time": times[0].strftime("%Y-%m-%d %H:%M:%S"),
                "end_time": times[1].strftime("%Y-%m-%d %H:%M:%S"),
                "link": link, 
                "location": location,
                "tags": random.sample(tags, random.randint(1,len(tags)))              
            })
        bar.next()

#         name: z.string().max(255),
#   preview: z.string().max(255),
#   description: z.string(),
#   start_time: z.string(),
#   end_time: z.string(),
#   location: z.string().max(255),
#   link: z.string().max(255).optional(),
#   event_type: eventTypeEnum,
#   is_recurring: z.boolean().optional(),
#   is_public: z.boolean(),
#   is_draft: z.boolean(),
#   is_archived: z.boolean(),
#   host: z.string().uuid(),
    
    bar.finish()
    print(f"Bad club descriptions (JSON not created properly): {bad_club_descriptions}")

    # Generate the UUIDS for Categories, Tags, and a table for tags to categories.
    category_uuids = {}
    for category in MOCK_CATEGORIES_AND_TAGS:
        category_uuids[category] = str(uuid4()) 

    tag_uuids = {}
    for category, tags in MOCK_CATEGORIES_AND_TAGS.items():
        for tag in tags:
            tag_uuids[tag] = str(uuid4())

    tags_to_categories = {}
    for category, tags in MOCK_CATEGORIES_AND_TAGS.items():
        for tag in tags:
            tags_to_categories[tag] = category

    
    with open("MOCK_DATA_OUTPUT" + ".json", 'w', encoding='utf-8') as file:
        print("Writing JSON file of mock data (needed in case of debugging, free to delete otherwise)")
        file.write(json.dumps({
            "category_uuids": category_uuids,
            "tag_uuids": tag_uuids,
            "tags_to_categories": tags_to_categories,
            "clubs": clubs,
            "events": events
        }
    ))

if not os.path.exists("MOCK_DATA_OUTPUT.json"):
    print("Error: no MOCK_DATA_OUTPUT.json found. Did you try running this tool with '--skip-generation=False'?")
    exit(0)

# STEP 3: Connect to DB and write entries.
category_uuids = {}
tag_uuids = {}
tags_to_categories = {}
clubs = []
events = []

with open('MOCK_DATA_OUTPUT.json', 'r', encoding='utf-8') as file:
    data = json.loads(file.read())
    clubs = data['clubs']
    events = data['events']
    tags_to_categories = data['tags_to_categories']
    tag_uuids = data['tag_uuids']
    category_uuids = data['category_uuids']
    


# FIXME: should probably be in args or something but ughhhhh 
db_conn = psycopg2.connect(
    dbname="sac",
    user="postgres",
    password="password",
    host="127.0.0.1",
    port="5432"
)

cur = db_conn.cursor()

# STEP 3: Write to SQL file.
def transaction(func):
    def wrapper(*args, **kwargs):
        try:
            func(*args, **kwargs)
        except psycopg2.OperationalError as e:
            db_conn.rollback()
            print(e)
            exit(0)
        except psycopg2.ProgrammingError as e:
            db_conn.rollback()
            print(e)
            exit(0)
        finally:
            db_conn.commit() 
    return wrapper

@transaction
def delete_existing():
    print("Deleting from existing tables...")
    cur.execute("DELETE FROM clubs")
    cur.execute("DELETE FROM categories")
    cur.execute("DELETE FROM tags")
    cur.execute("DELETE FROM club_tags")

@transaction
def fill_clubs():
    print("Creating club data...")
    cmd = """INSERT INTO "clubs" ("id", "name", "preview", "description", "logo", "num_members", "parent") VALUES (%s, %s, %s, %s, %s, %s, %s)"""

    for club in clubs:
        data = (
            club['id'],
            club['name'].replace("\'", "''"),
            club['preview'].replace("\'", "''"),
            club['description'].replace('\'', "''"),
            club['logo'],
            str(club['num_members']),
            PARENT_UUID
        )

        cur.execute(cmd, data)

@transaction
def fill_events():
    print("Creating event data...")
    cmd = """INSERT INTO "events" ("id", "host", "name", "preview", "description", "event_type", "start_time", "end_time", "is_public", "is_draft", "is_archived") VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"""

    for event in events:
        data = (
            event['id'],
            event['club_id'],
            event['name'].replace("\'", "''"),
            event['preview'].replace("\'", "''"),
            event['description'].replace('\'', "''"),
            event['event_type'],
            event['start_time'],
            event['end_time'],
            True,
            False,
            False
        )

        cur.execute(cmd, data)

@transaction
def fill_tags_and_categories():
    print("Creating category and tag data...")
    cmd = """INSERT INTO "categories" ("id", "name") VALUES (%s, %s);"""
    for category, uuid in category_uuids.items():
        cur.execute(cmd, (uuid, category))
    
    cmd = """INSERT INTO "tags" ("id", "name", "category_id") VALUES (%s, %s, %s)"""
    for tag, uuid in tag_uuids.items():
        cur.execute(cmd, (uuid, tag, category_uuids[tags_to_categories[tag]]))

@transaction
def connect_clubs_to_tags():
    print("Adding tags to clubs...")
    cmd = """INSERT INTO "club_tags" ("tag_id", "club_id") VALUES (%s, %s);"""
   
    # FIXME (maybe)?: The ChatGPT tags prompt will return tags that aren't in the list, hence this if statement being here.
    # Don't think its a huge issue because ~half of tags are good. 
    # FIXME (maybe)?: The ChatGPT tags prompt will return duplicate tags. oml
    bad_tags = 0
    total_tags = 0 
    for club in clubs:
        clubTagsWithDuplicatesRemoved = list(set(club['tags']))
        total_tags += len(clubTagsWithDuplicatesRemoved)
        for tag in clubTagsWithDuplicatesRemoved:
            if tag in tag_uuids:
                cur.execute(cmd, (tag_uuids[tag], club['id']))
            else:
                bad_tags += 1

    print(f"{bad_tags}/{total_tags} marked for not existing in tag_uuids.")

@transaction 
def connect_events_to_tags():
    print("Adding tags to events...")
    cmd = """INSERT INTO "event_tags" ("tag_id", "event_id") VALUES (%s, %s);"""
   
    # FIXME (maybe)?: The ChatGPT tags prompt will return tags that aren't in the list, hence this if statement being here.
    # Don't think its a huge issue because ~half of tags are good. 
    # FIXME (maybe)?: The ChatGPT tags prompt will return duplicate tags. oml
    bad_tags = 0
    total_tags = 0 
    for event in events:
        eventTagsWithDuplicatesRemoved = list(set(event['tags']))
        total_tags += len(eventTagsWithDuplicatesRemoved)
        for tag in eventTagsWithDuplicatesRemoved:
            if tag in tag_uuids:
                cur.execute(cmd, (tag_uuids[tag], event['id']))
            else:
                bad_tags += 1

    print(f"{bad_tags}/{total_tags} marked for not existing in tag_uuids.")

@transaction
def connect_clubs_to_events():
    print("Connecting events to clubs...")
    cmd = """INSERT INTO "club_events" ("club_id", "event_id") VALUES (%s, %s);"""

    for event in events:
        data = (event["club_id"], event["id"])

        cur.execute(cmd, data)

delete_existing()
fill_clubs()
fill_events()
fill_tags_and_categories()
connect_clubs_to_tags()
connect_events_to_tags()
connect_clubs_to_events()
