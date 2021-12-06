create database blog;
-- ----------------------------
-- Table structure for cate
-- ----------------------------
DROP TABLE IF EXISTS `cate`;
CREATE TABLE `cate` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '' COMMENT '分类名',
  `intro` varchar(255) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cate_id` int(11) DEFAULT NULL COMMENT '分类Id',
  `kind` int(11) DEFAULT NULL COMMENT '类型1-文章，2-页面',
  `status` int(11) DEFAULT NULL COMMENT '状态1-草稿，2-已发布',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `path` varchar(255) DEFAULT NULL COMMENT '访问路径',
  `summary` text DEFAULT NULL COMMENT '摘要',
  `markdown` mediumtext DEFAULT NULL COMMENT 'markdown内容',
  `richtext` mediumtext DEFAULT NULL COMMENT '富文本内容',
  `allow` tinyint(4) DEFAULT 1 COMMENT '允许评论',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '修改时间',
  `creator` id DEFAULT NULL COMMENT '创建人',
  `updater` id DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_post_path` (`path`),
  KEY `create_time` (`created`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) DEFAULT NULL COMMENT '姓名',
  `num` varchar(255) DEFAULT NULL COMMENT '账号',
  `passwd` varchar(255) DEFAULT NULL COMMENT '密码',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(255) DEFAULT NULL COMMENT '电话',
  `ecount` int(11) DEFAULT 0 COMMENT '错误次数',
  `ltime` datetime DEFAULT NULL COMMENT '上次登录时间',
  `ctime` datetime DEFAULT NULL COMMENT '创建时间',
  `openid_qq` varchar(64) DEFAULT NULL COMMENT 'qq_openid',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UQE_sys_user_num` (`num`),
  UNIQUE KEY `UQE_sys_user_openid_qq` (`openid_qq`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

insert into `sys_user` (name,num,passwd) values('system','system','system')

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '标签名',
  `intro` varchar(255) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4;
