use mysql;

enum DbType{
    Mysql,
    PostgreSQL,
    Oracle,
    Sql_Server
}

pub struct  DbMgr{
    dbtype :DbType,

}

// pub trait Factory{
//     fn connect(&self,db_type:DbType) ->Box<dyn DbMgr>{
//         match db_type{
//             DbType::Mysql => {
//                 mysq
//             }
//             DbType::PostgreSQL => {}
//             DbType::Oracle => {}
//             DbType::Sql_Server => {}
//         }
//     }
// }