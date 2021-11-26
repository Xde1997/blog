use serde_json::{Deserializer, Serializer};
use serde::Deserialize;

pub mod JsonFile {
    use std::io::Read;

    pub fn JsonFileParser<T>(filepath: &str, s: &mut T) -> T
    where T: serde::Deserialize{
        let mut file = std::fs::File::open(filepath).unwrap();
        let mut buffer;
        let result = file.read_to_string(&mut buffer)?;
        let deserialized:T = serde_json::from_str(&buffer);
        return deserialized;
    }
}
