create table  if not exists `config_item`(
    `objid` int unsigned auto_increment,
    `type` int default 0 not null,
    `content` LONGTEXT default null,
    `label` varchar(255) not null,
    `servicename` varchar(255) not null,
    `module` varchar(255) not null,
    `function` varchar(255) not null,
    `properties` varchar(255) not null,
    `createdate` date,
    `creator` varchar(255) not null,
    `lastmodifier`  varchar(255) not null,
    `lastmodifydate` date,
    PRIMARY KEY (objid)
)engine=innodb default charset=utf8;
CREATE INDEX searchConfig ON config_item (label,servicename,module,type)

-- /blog/api/configd/{label:dev/release}/{servicename}/{module}/{function}/{properties}