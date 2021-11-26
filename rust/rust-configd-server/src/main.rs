mod demo;
mod file;
mod init;
mod config;

use actix_web::{web, App, HttpServer};
use demo::common_demo;
use config::config as config_item;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let mut instancecfg: init::init::Cfg::serverInstanceCfg;
    let result = file::file::jsonFile::jsonFileParser(
        "../../cfg/serverInstanceCfg.json",
        &mut instancecfg,
    )
    .unwrap();
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
