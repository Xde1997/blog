use chrono::prelude::*;
use serde::de;
use serde::{Deserialize,Serialize};

#[derive(Deserialize,Serialize,Debug)]
pub struct config_item {
    objid: u32,
    r#type: String,
    content: String,
    label: String,
    service_name: String,
    module: String,
    function: String,
    properties: String,
    create_date:DateTime<Utc>,
    creator: String,
    last_modifier:String,
    last_modify_date:DateTime<Utc>
}
impl config_item{
    pub fn Query(){}
    pub fn Save(){}
    pub fn Delete(){}
}

pub struct configmgr{

}