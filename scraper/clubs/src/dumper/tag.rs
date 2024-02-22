use std::{error::Error, fs::File, io::Write};

use crate::domain::tag::Tag;

pub fn dump(tags: &Vec<&'static Tag>, file: &mut File) -> Result<(), Box<dyn Error>> {
    for tag in tags {
        writeln!(
            file,
            r#"INSERT INTO "tags" ("id", "name", "category_id") VALUES ('{}', '{}', '{}');"#,
            tag.id, tag.name, tag.category_id
        )?;
    }

    Ok(())
}
