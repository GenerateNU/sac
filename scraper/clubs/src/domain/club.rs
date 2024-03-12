use rand::thread_rng;
use strum::IntoEnumIterator;
use strum_macros::{Display, EnumIter};

use rand::{seq::IteratorRandom, Rng};

use voca_rs::Voca;

use crate::scraper::ScrapedClub;

use crate::domain::tag::{Tag, TAGS};

#[derive(Debug, PartialEq, EnumIter, Display)]
pub enum RecruitmentCycle {
    #[strum(serialize = "fall")]
    Fall,
    #[strum(serialize = "spring")]
    Spring,
    #[strum(serialize = "fallSpring")]
    FallSpring,
    #[strum(serialize = "always")]
    Always,
}

#[derive(Debug, PartialEq, EnumIter, Display)]
pub enum RecruitmentType {
    #[strum(serialize = "unrestricted")]
    Unrestricted,
    #[strum(serialize = "application")]
    Tryout,
    #[strum(serialize = "application")]
    Application,
}

#[derive(Debug)]
pub struct Club {
    pub id: uuid::Uuid,
    pub name: String,
    pub preview: String,
    pub description: String,
    pub num_members: usize,
    pub is_recruiting: bool,
    pub recruitment_cycle: RecruitmentCycle,
    pub recruitment_type: RecruitmentType,
    pub tags: Vec<&'static &'static Tag>,
}

impl<'a> From<&'a ScrapedClub> for Club {
    fn from(scraped: &'a ScrapedClub) -> Self {
        let mut rng = thread_rng();

        let num_tags = rng.gen_range(1..8);

        Club {
            id: uuid::Uuid::new_v4(),
            name: scraped.name.clone().replace("(Tentative) ", ""),
            preview: scraped.preview.clone(),
            description: scraped.description._strip_tags().replace("&nbsp;", " "),
            num_members: rng.gen_range(1..1024),
            is_recruiting: rng.gen_bool(0.5),
            recruitment_cycle: RecruitmentCycle::iter().choose(&mut rng).unwrap(),
            recruitment_type: RecruitmentType::iter().choose(&mut rng).unwrap(),
            tags: TAGS.iter().choose_multiple(&mut rng, num_tags),
        }
    }
}
