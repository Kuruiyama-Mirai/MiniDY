DROP TABLE IF EXISTS `favourite`;
CREATE TABLE `favorite`  (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '点赞表的主键',
    `user_id` int(11) NOT NULL COMMENT '点赞人的id',
    `video_id` int(11) NOT NULL COMMENT '视频id',
    `status` int(1) NOT NULL COMMENT '点赞状态1为点赞2为取消',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '逻辑删除',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `user_id`(`user_id`) USING BTREE COMMENT '针对用户设置索引',
    UNIQUE KEY `user_id_2`(`user_id`, `video_id`) USING BTREE COMMENT '针对userid和videoid设置联合索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
