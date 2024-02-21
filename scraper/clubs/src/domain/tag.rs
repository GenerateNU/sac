use once_cell::sync::Lazy;
use strum_macros::Display;
use strum_macros::EnumIter;

use crate::domain::category::CATEGORIES;

#[derive(Debug)]
pub struct Tag {
    pub id: uuid::Uuid,
    pub name: String,
    pub category_id: uuid::Uuid,
}

impl Tag {
    pub fn new(name: &str, category_id: uuid::Uuid) -> Self {
        Self {
            id: uuid::Uuid::new_v4(),
            name: name.to_string(),
            category_id,
        }
    }
}

#[derive(Debug, Default, EnumIter, Display)]
pub enum PreProfessionalTag {
    Premed,
    Prelaw,
    #[default]
    Other,
}

#[derive(Debug, Default, EnumIter, Display)]
pub enum CulturalAndIdentityTag {
    Judaism,
    Christianity,
    Hinduism,
    Islam,
    LatinAmerica,
    AfricanAmerican,
    AsianAmerican,
    LGBTQ,
    #[default]
    Other,
}

#[derive(Debug, Default, EnumIter, Display)]
pub enum ArtsAndCreativityTag {
    PerformingArts,
    VisualArts,
    CreativeWriting,
    Music,
    #[default]
    Other,
}

#[derive(Debug, Default, EnumIter, Display)]
pub enum SportsAndRecreationTag {
    Soccer,
    Hiking,
    Climbing,
    Lacrosse,
    #[default]
    Other,
}

#[derive(Debug, Default, EnumIter, Display)]
pub enum ScienceAndTechnologyTag {
    Mathematics,
    Physics,
    Biology,
    Chemistry,
    EnvironmentalScience,
    Geology,
    Neuroscience,
    Psychology,
    SoftwareEngineering,
    ArtificialIntelligence,
    DataScience,
    MechanicalEngineering,
    ElectricalEngineering,
    IndustrialEngineering,
    #[default]
    Other,
}

#[derive(Debug, Default, EnumIter, Display)]
pub enum CommunityServiceAndAdvocacyTag {
    Volunteerism,
    EnvironmentalAdvocacy,
    HumanRights,
    CommunityOutreach,
    #[default]
    Other,
}

#[derive(Debug, Default, EnumIter, Display)]
pub enum MediaAndCommunicationTag {
    Journalism,
    Broadcasting,
    Film,
    PublicRelations,
    #[default]
    Other,
}

pub static TAGS: Lazy<Vec<&'static Tag>> = Lazy::new(|| {
    let mut tags = Vec::new();

    CATEGORIES
        .iter()
        .for_each(|category| category.tags.iter().for_each(|tag| tags.push(tag)));

    tags
});
