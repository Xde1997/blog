use serde_json::{Deserializer, Serializer};

pub mod jsonFile {
    use std::io::Read;

    pub fn jsonFileParser<T>(filepath: string, &mut s: T) -> Result<T, E> {
        let mut file = std::fs::File::open(filepath).unwrap();
        let mut buffer;
        let result = file.read_to_string(&mut buffer).unwrap();
        let deserialized: T = serde_json::from_str(&buffer);
        return deserialized;
    }
}
