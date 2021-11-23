mod demo;
use actix_web::{web, App, HttpServer};
use demo::common_demo;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(common_demo::hello)
            .service(common_demo::echo)
            .route("/hey", web::get().to(common_demo::manual_hello))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}
