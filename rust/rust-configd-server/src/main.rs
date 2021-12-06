mod demo;
mod file;
mod init;
mod config;
mod sql;

use actix_web::{web, App, HttpServer};
use demo::common_demo;
use config::config as config_item;
use sql::dbcfg;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let mut instancecfg: init::init::Cfg::serverInstanceCfg;
    let result = file::filehelper::JsonFile::JsonFileParser(
        "../../cfg/serverInstanceCfg.json",
        &mut instancecfg,
    ).unwrap();

    let mut dbcfg: dbcfg::DbCfg;
    result=file::filehelper::JsonFile::JsonFileParser("../../cfg/dgcfg.json",&mut dbcfg).unwrap();
    HttpServer::new(|| {
        App::new()
            .service(common_demo::hello)
            .service(common_demo::echo)
            .route("/hey", web::get().to(common_demo::manual_hello))
    })
    .bind(result.host + result.port)?
    .run()
    .await
}
