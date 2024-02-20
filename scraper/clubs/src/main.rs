use std::error::Error;
use std::ops::Deref;
use std::{fs::OpenOptions, process::exit};

use clap::Parser;
use sac_club_scraper::domain::category::CATEGORIES;
use sac_club_scraper::domain::tag::TAGS;
use sac_club_scraper::dumper::dump::dump_all;
use sac_club_scraper::{cli::Args, domain::club::Club};

use sac_club_scraper::scraper::ClubsResponse;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let args = Args::parse();

    let response = reqwest::get(format!("https://neu.campuslabs.com/engage/api/discovery/search/organizations?orderBy[0]=UpperName%20asc&top={}&skip=0", args.top_n)).await?;

    if !response.status().is_success() {
        println!("Request failed with status: {}", response.status());
        exit(1);
    }

    let body = response.text().await?;

    let response: ClubsResponse = serde_json::from_str(&body).expect("Failed to deserialize");

    let mut file = OpenOptions::new()
        .create(true)
        .append(true)
        .open(args.output)?;

    dump_all(
        CATEGORIES.deref(),
        TAGS.deref(),
        response.scraped_clubs.iter().map(Club::from).collect(),
        args.parent,
        &mut file,
    )?;

    Ok(())
}
