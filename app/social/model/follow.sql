DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`  (
  `id` int(11) NOT NULL COMMENT '关注表的主键id',
  `user_id` int(11) NOT NULL COMMENT '关注人的id',
  `to_user_id` int(11) NOT NULL COMMENT '被关注人id',
  `status` int(1) NOT NULL COMMENT '是否为互相关注 1为相互关注 0不是', 
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '逻辑删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_id`(`user_id`) USING BTREE COMMENT '关注列表使用此索引',
  UNIQUE KEY `to_user_id`(`to_user_id`) USING BTREE COMMENT '粉丝列表使用此索引'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
