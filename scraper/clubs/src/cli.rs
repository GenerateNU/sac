use std::path::PathBuf;

use clap::Parser;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
pub struct Args {
    /// Top N results to parse
    #[clap(short, long, default_value = "2")]
    pub top_n: usize,
    /// Output file
    #[clap(short, long, default_value = "mock.sql")]
    pub output: PathBuf,
    /// Parent club UUID
    #[clap(short, long, default_value = "00000000-0000-0000-0000-000000000000")]
    pub parent: uuid::Uuid,
}
