use mysql::*;
use mysql::prelude::*;
use chrono::prelude::*;
use lazy_static::lazy_static;

// rust bu zhichi zhishengming
//static mut db :mysql::Conn;

// lazy_static!{
//     static mut db :mysql::Conn;
// }

pub struct Mysqlmgr {
    db_conn:mysql::Conn,
}

impl Mysqlmgr{
    pub fn new()->Self{
        return Self
    }
}