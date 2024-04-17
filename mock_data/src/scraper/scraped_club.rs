use std::fmt;

use rand::{thread_rng, Rng};
use serde::{
    de::{self, Visitor},
    Deserializer,
};

use lipsum::lipsum;

#[derive(Debug, serde::Deserialize)]
pub struct ClubsResponse {
    #[serde(rename = "value")]
    pub scraped_clubs: Vec<ScrapedClub>,
}

#[derive(Debug, serde::Deserialize)]
pub struct ScrapedClub {
    #[serde(rename = "Name", deserialize_with = "deserialize_with_lipsum")]
    pub name: String,
    #[serde(rename = "Summary", deserialize_with = "deserialize_with_lipsum")]
    pub preview: String,
    #[serde(rename = "Description", deserialize_with = "deserialize_with_lipsum")]
    pub description: String,
}

fn deserialize_with_lipsum<'de, D>(deserializer: D) -> Result<String, D::Error>
where
    D: Deserializer<'de>,
{
    struct StringOrLipsum;

    impl<'de> Visitor<'de> for StringOrLipsum {
        type Value = String;

        fn expecting(&self, formatter: &mut fmt::Formatter) -> fmt::Result {
            formatter.write_str("a string or null for a lipsum replacement")
        }

        fn visit_str<E>(self, value: &str) -> Result<Self::Value, E>
        where
            E: de::Error,
        {
            Ok(value.to_owned())
        }

        fn visit_unit<E>(self) -> Result<Self::Value, E>
        where
            E: de::Error,
        {
            let mut rng = thread_rng();

            Ok(lipsum(rng.gen_range(16..128)))
        }
    }

    deserializer.deserialize_any(StringOrLipsum)
}
