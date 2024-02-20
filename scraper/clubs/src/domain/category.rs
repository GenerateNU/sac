use once_cell::sync::Lazy;
use strum::IntoEnumIterator;
use strum_macros::{Display, EnumIter};

use crate::domain::tag::{
    ArtsAndCreativityTag, CommunityServiceAndAdvocacyTag, CulturalAndIdentityTag,
    MediaAndCommunicationTag, PreProfessionalTag, ScienceAndTechnologyTag, SportsAndRecreationTag,
    Tag,
};

#[derive(Debug, EnumIter, Display)]
pub enum CategoryExample {
    PreProfessional(PreProfessionalTag),
    CulturalAndIdentity(CulturalAndIdentityTag),
    ArtsAndCreativity(ArtsAndCreativityTag),
    SportsAndRecreation(SportsAndRecreationTag),
    ScienceAndTechnology(ScienceAndTechnologyTag),
    CommunityServiceAndAdvocacy(CommunityServiceAndAdvocacyTag),
    MediaAndCommunication(MediaAndCommunicationTag),
}

#[derive(Debug)]
pub struct Category {
    pub id: uuid::Uuid,
    pub category: CategoryExample,
    pub tags: Vec<Tag>,
}

macro_rules! match_category {
    ($expr:expr, $id:expr, $( $variant:ident => $type:ty ),*) => {{
        fn create_tags_for_category<T: IntoEnumIterator + std::fmt::Display>(
            category_id: uuid::Uuid,
        ) -> Vec<Tag> {
            T::iter()
                .map(|tag| Tag::new(&tag.to_string(), category_id))
                .collect()
        }

        match $expr {
            $(
                CategoryExample::$variant(_) => create_tags_for_category::<$type>($id),
            )*
        }
    }};
}

pub static CATEGORIES: Lazy<Vec<Category>> = Lazy::new(|| {
    CategoryExample::iter()
        .map(|category_example| {
            let id = uuid::Uuid::new_v4();

            let tags = match_category! { category_example, id,
                PreProfessional => PreProfessionalTag,
                CulturalAndIdentity => CulturalAndIdentityTag,
                ArtsAndCreativity => ArtsAndCreativityTag,
                SportsAndRecreation => SportsAndRecreationTag,
                ScienceAndTechnology => ScienceAndTechnologyTag,
                CommunityServiceAndAdvocacy => CommunityServiceAndAdvocacyTag,
                MediaAndCommunication => MediaAndCommunicationTag
            };

            Category {
                id,
                category: category_example,
                tags,
            }
        })
        .collect()
});
