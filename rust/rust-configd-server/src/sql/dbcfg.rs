use serde::{Deserialize,Serialize};
use crate::file::filehelper;

#[derive(Debug,Deserialize,Serialize)]
pub struct DbCfg {
    sql_type:String,
    db_cfg: String,
    user:String,
    password:String
}


// pub fn Init(filepath:String){
//     filehelper::JsonFile::JsonFileParser(filepath.as_str(),)
// }