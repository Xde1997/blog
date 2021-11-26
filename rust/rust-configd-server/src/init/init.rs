pub mod Cfg {
    pub struct serverInstanceCfg {
        kind: String,
        instance_name: String,
        sd_host: String,
        host: String,
        port: u32,
        support_cors: bool,
        local_enable: bool,
        support_sr: bool,
    }
}
