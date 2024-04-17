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