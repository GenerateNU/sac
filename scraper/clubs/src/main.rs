use std::process::exit;

use clap::Parser;
use sac_scraper::{cli::Args, domain::club::Club, dumper::sql::dump};

use reqwest::Error;
use sac_scraper::scraper::ClubsResponse;

#[tokio::main]
async fn main() -> Result<(), Error> {
    let args = Args::parse();

    let response = reqwest::get(format!("https://neu.campuslabs.com/engage/api/discovery/search/organizations?orderBy[0]=UpperName%20asc&top={}&skip=0", args.top_n)).await?;

    if !response.status().is_success() {
        println!("Request failed with status: {}", response.status());
        exit(1);
    }

    let body = response.text().await?;

    let response: ClubsResponse = serde_json::from_str(&body).expect("Failed to deserialize");

    dump(
        response.scraped_clubs.iter().map(Club::from).collect(),
        args.output,
        args.parent,
    )
    .expect("Failed to dump");

    Ok(())
}
