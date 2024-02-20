use once_cell::sync::Lazy;
use strum::IntoEnumIterator;
use strum_macros::Display;
use strum_macros::EnumIter;

#[derive(Debug)]
pub struct Tag {
    pub category: Category,
    pub name: String,
}

#[derive(Debug)]
pub enum Category {
    PreProfessional,
    CulturalAndIdentity,
    ArtsAndCreativity,
    SportsAndRecreation,
    ScienceAndTechnology,
    CommunityServiceAndAdvocacy,
    MediaAndCommunication,
}

fn tags_for_category<T: IntoEnumIterator + std::fmt::Display + TagCategory>() -> Vec<Tag> {
    T::iter()
        .map(|item| Tag {
            category: T::category(),
            name: item.to_string(),
        })
        .collect()
}

pub static TAGS: Lazy<Vec<Tag>> = Lazy::new(|| {
    let mut tags = Vec::new();

    tags.extend(tags_for_category::<PreProfessional>());
    tags.extend(tags_for_category::<CulturalAndIdentity>());
    tags.extend(tags_for_category::<ArtsAndCreativity>());
    tags.extend(tags_for_category::<SportsAndRecreation>());
    tags.extend(tags_for_category::<ScienceAndTechnology>());
    tags.extend(tags_for_category::<CommunityServiceAndAdvocacy>());
    tags.extend(tags_for_category::<MediaAndCommunication>());

    tags
});

trait TagCategory {
    fn category() -> Category;
}

#[derive(EnumIter, Display)]
pub enum PreProfessional {
    Premed,
    Prelaw,
    Other,
}

impl TagCategory for PreProfessional {
    fn category() -> Category {
        Category::PreProfessional
    }
}

#[derive(EnumIter, Display)]
pub enum CulturalAndIdentity {
    Judaism,
    Christianity,
    Hinduism,
    Islam,
    LatinAmerica,
    AfricanAmerican,
    AsianAmerican,
    LGBTQ,
    Other,
}

impl TagCategory for CulturalAndIdentity {
    fn category() -> Category {
        Category::CulturalAndIdentity
    }
}

#[derive(EnumIter, Display)]
pub enum ArtsAndCreativity {
    PerformingArts,
    VisualArts,
    CreativeWriting,
    Music,
    Other,
}

impl TagCategory for ArtsAndCreativity {
    fn category() -> Category {
        Category::ArtsAndCreativity
    }
}

#[derive(EnumIter, Display)]
pub enum SportsAndRecreation {
    Soccer,
    Hiking,
    Climbing,
    Lacrosse,
    Other,
}

impl TagCategory for SportsAndRecreation {
    fn category() -> Category {
        Category::SportsAndRecreation
    }
}

#[derive(EnumIter, Display)]
pub enum ScienceAndTechnology {
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
    Other,
}

impl TagCategory for ScienceAndTechnology {
    fn category() -> Category {
        Category::ScienceAndTechnology
    }
}

#[derive(EnumIter, Display)]
pub enum CommunityServiceAndAdvocacy {
    Volunteerism,
    EnvironmentalAdvocacy,
    HumanRights,
    CommunityOutreach,
    Other,
}

impl TagCategory for CommunityServiceAndAdvocacy {
    fn category() -> Category {
        Category::CommunityServiceAndAdvocacy
    }
}

#[derive(EnumIter, Display)]
pub enum MediaAndCommunication {
    Journalism,
    Broadcasting,
    Film,
    PublicRelations,
    Other,
}

impl TagCategory for MediaAndCommunication {
    fn category() -> Category {
        Category::MediaAndCommunication
    }
}
