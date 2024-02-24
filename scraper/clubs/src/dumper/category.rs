use std::{error::Error, fs::File, io::Write};

use crate::domain::Category;

pub fn dump(categories: &Vec<Category>, file: &mut File) -> Result<(), Box<dyn Error>> {
    for category in categories {
        writeln!(
            file,
            r#"INSERT INTO "categories" ("id", "name") VALUES ('{}', '{}');"#,
            category.id, category.category
        )?;
    }
    Ok(())
}
