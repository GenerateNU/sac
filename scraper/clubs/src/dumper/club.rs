use std::{error::Error, fs::File, io::Write};

use crate::domain::club::Club;

pub fn dump(clubs: Vec<Club>, file: &mut File, parent: uuid::Uuid) -> Result<(), Box<dyn Error>> {
    for club in clubs {
        writeln!(
            file,
            r#"INSERT INTO "clubs" ("name", "preview", "description", "num_members", "is_recruiting", "recruitment_cycle", "recruitment_type", "parent") VALUES ('{}', '{}', '{}', {}, {}, '{}', '{}', '{}');"#,
            club.name.replace('\'', "''"),
            club.preview.replace('\'', "''"),
            club.description.replace('\'', "''"),
            club.num_members,
            club.is_recruiting,
            club.recruitment_cycle,
            club.recruitment_type,
            parent,
        )?;
    }

    Ok(())
}
