use serde_json::{Deserializer, Serializer};
use serde::de;

pub mod JsonFile {
    use std::io::Read;

    pub fn JsonFileParser<'a, T,E>(filepath: &str, s: &mut T) -> Result<T, E>
    where T: serde::de::Deserialize<'a>{
        let mut file = std::fs::File::open(filepath).unwrap();
        let mut buffer;
        let result = file.read_to_string(&mut buffer)?;
        serde_json::from_str(&buffer)?
    }
}
