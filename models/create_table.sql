CREATE TABLE `user` (--``包裹列名或表名以使用特殊字符、区分大小写
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) not null,
    `username` varchar(64) collate utf8mb4_general_ci not null,
    `password` varchar(64) collate utf8mb4_general_ci not null,
    `email` varchar(64) collate utf8mb4_general_ci,
    `gender` tinyint(4) not null default '0',
    `create_time` timestamp null default current_timestamp,
    `update_time` timestamp null default current_timestamp on update current_timestamp,
    primary key(`id`),
    unique key `idx_username` (`username`) using btree,--给属性设置一个用于作为索引的别名
    unique key `idx_user_id` (`user_id`) using btree
)engine=InnoDB default charset=utf8mb4 collate=utf8mb4_general_ci;

Drop table if exists `community`;
create table `community`
(
    `id`             int(11) not null auto_increment,
    `community_id`   int(10) unsigned not null,
    `community_name` varchar(128) collate utf8mb4_general_ci not null,
    `introduction`   varchar(256) collate utf8mb4_general_ci not null,
    `create_time`    timestamp                               not null default current_timestamp,
    `update_time`    timestamp                               not null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
)engine=InnoDB default charset=utf8mb4 collate=utf8mb4_general_ci;

Create table `post`(
    `id` bigint(20) not null auto_increment,
    `post_id` bigint(20) not null comment '帖子ID',
    `title` varchar(128) collate utf8mb4_general_ci not null comment '标题' ,
    `content` varchar(8192) collate utf8mb4_general_ci not null comment '内容',
    `author_id` bigint(20) not null comment '发布者的用户ID',
    `community_id` bigint(20) not null comment '所属社区',
    `status` tinyint(4) not null default '1' comment '帖子状态',
    `create_time` timestamp null default current_timestamp comment '创建时间',
    `update_time` timestamp null default current_timestamp on update current_timestamp comment '更新时间',
    primary key (`id`),
    unique key `id_post_id` (`post_id`),
    key `idx_author_id` (`author_id`)
    key `idx_community_id` (`community_id`)
)engine=InnoDB default charset=utf8mb4 collate=utf8mb4_general_ci;