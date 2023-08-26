DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '评论表id',
    `video_id` int(11) NOT NULL COMMENT '视频id',
    `user_id` int(11) NOT NULL COMMENT '评论用户的id',
    `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '逻辑删除',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY (`video_id`) USING BTREE COMMENT '为被评论视频设置索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci；

