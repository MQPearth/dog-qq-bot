SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for story
-- ----------------------------
DROP TABLE IF EXISTS `story`;
CREATE TABLE `story`
(
    `id`          int(0) NOT NULL AUTO_INCREMENT,
    `content`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
    `type`        tinyint(0) NOT NULL COMMENT '1-鬼故事 2-毒鸡汤 3-阴阳',
    `insert_date` datetime(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP (0),
    `status`      tinyint(0) NOT NULL DEFAULT 1 COMMENT '1-正常 2-删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

SET
FOREIGN_KEY_CHECKS = 1;
