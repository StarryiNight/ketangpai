/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : ketangpai

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 05/01/2022 15:30:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `v0` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `v1` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `v2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `v3` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `v4` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `v5` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  INDEX `idx_casbin_rule`(`p_type`, `v0`, `v1`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/api/v1/student*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/api/v1/community/*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/api/v1/community/*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/api/v1/student*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/lesson*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/lesson*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/class*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/class*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/lesson*', 'PATCH', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/community/*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/test/*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/test/*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/password/*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/api/v1/password/*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/api/v1/file/*', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'student', '/api/v1/file/*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/file/*', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'teacher', '/api/v1/file/*', 'GET', '', '', '');

-- ----------------------------
-- Table structure for choicequestion
-- ----------------------------
DROP TABLE IF EXISTS `choicequestion`;
CREATE TABLE `choicequestion`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `question_id` bigint NOT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `options` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `answer` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`, `question_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of choicequestion
-- ----------------------------
INSERT INTO `choicequestion` VALUES (1, '思政', 372172907624792065, '把马克思列宁主义原理同中国革命实际相结合，实现第一次历史性飞跃的代表人物是（）。', 'A.毛泽东 B.邓小平 C.周恩来 D.朱德', 'A');
INSERT INTO `choicequestion` VALUES (2, '思政', 372189741614891009, '把毛泽东思想确定为党的指导思想的会议是（）。', 'A.遵义会议 B.中共七大 C.中共六届七中全会 D.中共六届六中全会', 'B');
INSERT INTO `choicequestion` VALUES (3, '思政', 372189807515795457, '马克思主义中国化的第一个理论成果是（）。', 'A.邓小平理论 B.毛泽东思想 C.三个代表重要思想 D.中国特色社会主义理论体系', 'B');
INSERT INTO `choicequestion` VALUES (4, '思政', 372189871336325121, '作为党的思想路线“实事求是”中的“是”是指（）。', 'A.客观存在着的一切事物 B.主观对客观的能动反映 C.客观事物的内部联系，即规律性 D.客观事物的运动、变化和发展，即辩证法', 'C');
INSERT INTO `choicequestion` VALUES (5, '思政', 372189871336325123, '在马克思主义中国化的过程中,产生了毛泽东思想和中国特色社会主义理论体系,这两大理论体系一脉相承主要体现在:二者具有共同的（）。', 'A.马克思主义的理论基础\r\n\r\nB.革命和建设的根本任务\r\n\r\nC.实事求是的理论基础\r\n\r\nD.和平与发展的时代背景', 'AC');
INSERT INTO `choicequestion` VALUES (6, '思政', 372189871336325122, '新民主主义革命时期压在中国人民头上的“三座大山”是指（）', 'A.帝国主义 B.封建主义 C.官僚资本主义 D.资本主义', 'ABC');
INSERT INTO `choicequestion` VALUES (7, '思政', 372189871336325124, '近代中国社会最主要的矛盾是（）', 'A.农民阶级和地主阶级的矛盾\r\n\r\nB.帝国主义和中华民族的矛盾\r\n\r\nC.封建主义和资本主义的矛盾\r\n\r\nD.民族资本主义和帝国主义的矛盾', 'B');
INSERT INTO `choicequestion` VALUES (8, '思政', 372189871336325126, '开辟世界无产阶级社会主义革命新纪元的是（）', 'A.五四运动\r\n\r\nB.俄国十月革命\r\n\r\nC.第一次世界大战\r\n\r\nD.辛亥革命', 'B');
INSERT INTO `choicequestion` VALUES (9, '思政', 372189871336325127, '新民主主义革命阶段的开端是（）', 'A.十月革命\r\n\r\nB.辛亥革命\r\n\r\nC.五四运动\r\n\r\nD.中国共产党的成立', 'C');
INSERT INTO `choicequestion` VALUES (10, '思政', 372189871336325128, '近代中国贫穷落后和一切灾祸的总根源是（）', 'A.封建主义\r\n\r\nB.帝国主义\r\n\r\nC.民族资本主义\r\n\r\nD.官僚资本主义', 'B');

-- ----------------------------
-- Table structure for class
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class`  (
  `class_id` int NOT NULL AUTO_INCREMENT,
  `class_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`class_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of class
-- ----------------------------
INSERT INTO `class` VALUES (1, '数学');
INSERT INTO `class` VALUES (2, '语文');
INSERT INTO `class` VALUES (3, '物理');
INSERT INTO `class` VALUES (4, '计算机导论');
INSERT INTO `class` VALUES (5, '数据库');
INSERT INTO `class` VALUES (6, '计算机网络');
INSERT INTO `class` VALUES (7, '数据结构');
INSERT INTO `class` VALUES (8, '算法设计');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comment_id` bigint UNSIGNED NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `post_id` bigint NOT NULL,
  `author_id` bigint NOT NULL,
  `parent_id` bigint NOT NULL DEFAULT 0,
  `status` tinyint UNSIGNED NOT NULL DEFAULT 1,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_comment_id`(`comment_id`) USING BTREE,
  INDEX `idx_author_Id`(`author_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, 371014145262223361, '已完成！', 1, 370984071263682561, 370985769638035457, 1, '2021-09-03 20:14:23', '2021-09-03 20:15:03');
INSERT INTO `comment` VALUES (6, 371016305446223873, '已完成2！', 2, 370984071263682561, 370985769638035457, 1, '2021-09-03 20:35:50', '2021-09-03 20:35:50');
INSERT INTO `comment` VALUES (7, 371017319796703233, '已完成2！', 3, 370984071263682561, 370985769638035457, 1, '2021-09-03 20:45:55', '2021-09-03 20:45:55');
INSERT INTO `comment` VALUES (8, 371965700048158721, '已完成2！', 3, 371032570705477633, 370985769638035457, 1, '2021-09-10 09:47:14', NULL);
INSERT INTO `comment` VALUES (9, 371987318531162113, '已完成2！', 4, 370984071263682561, 370985769638035457, 1, '2021-09-10 13:21:59', NULL);

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `community_id` int UNSIGNED NOT NULL,
  `community_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_community_id`(`community_id`) USING BTREE,
  UNIQUE INDEX `idx_community_name`(`community_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of community
-- ----------------------------
INSERT INTO `community` VALUES (1, 1, '公告', '公告情况', '2020-11-01 08:10:10', '2021-09-03 15:12:06');
INSERT INTO `community` VALUES (2, 2, '学习讨论', '*******', '2020-01-01 08:00:00', '2021-09-03 15:12:19');
INSERT INTO `community` VALUES (3, 3, '资料下载', '*******', '2021-08-07 08:30:00', '2021-09-03 15:12:20');
INSERT INTO `community` VALUES (4, 4, '作业展示', '*******', '2021-01-01 08:00:00', '2021-09-03 15:32:40');

-- ----------------------------
-- Table structure for gapfilling
-- ----------------------------
DROP TABLE IF EXISTS `gapfilling`;
CREATE TABLE `gapfilling`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `question_id` bigint NOT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `answer` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gapfilling
-- ----------------------------
INSERT INTO `gapfilling` VALUES (2, '思政', 372187810808987649, '把马克思列宁主义原理同中国革命实际相结合，实现第一次历史性飞跃的代表人物是（）。', '毛泽东');
INSERT INTO `gapfilling` VALUES (3, '思政', 372190115579035649, '思想政治教育学是（）为研究客体的（）应用科学。', '思想政治教育 综合性');
INSERT INTO `gapfilling` VALUES (4, '思政', 372189871336325127, '中国特色社会主义最本质的特征是（）', '中国共产党的领导');
INSERT INTO `gapfilling` VALUES (5, '思政', 372189871336325129, '十八大进一步把（）纳入到现代化建设布局里，形成了“五位一体”的建设布局', '生态文明');
INSERT INTO `gapfilling` VALUES (6, '思政', 372189871336325429, '(）是最大的民生', '就业');
INSERT INTO `gapfilling` VALUES (7, '思政', 372189871336326429, '()是当代中国发展进步的根本制度保障', '中国特色社会主义制度');
INSERT INTO `gapfilling` VALUES (8, '思政', 372189871336366429, '提高就业质量和人民收入水平，要坚持就业优先战略和积极就业政策，实现（）就业', '更高质量和更充分');
INSERT INTO `gapfilling` VALUES (9, '思政', 372189871356366429, '中国特色社会主义事业总体布局是（）', '五位一体');
INSERT INTO `gapfilling` VALUES (11, '思政', 372189871756366429, '我国社会生产力落后，经济基础薄弱，所以要优先重点发展（）', '国防工业');
INSERT INTO `gapfilling` VALUES (12, '思政', 372189871056366429, '社会主义社会的基本矛盾是（）', '非对抗性');

-- ----------------------------
-- Table structure for lesson
-- ----------------------------
DROP TABLE IF EXISTS `lesson`;
CREATE TABLE `lesson`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `lesson_id` bigint NOT NULL,
  `class_id` bigint NOT NULL,
  `teacher_id` bigint NOT NULL,
  `homework` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `start_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `end_time` timestamp NULL DEFAULT '2030-01-01 00:00:00',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lesson
-- ----------------------------
INSERT INTO `lesson` VALUES (1, 371380494127857665, 2, 370984071263682561, '布置作业1111', '2021-09-06 08:53:44', '2021-09-06 09:10:15');
INSERT INTO `lesson` VALUES (4, 371380519864107009, 2, 370984071263682561, NULL, '2021-09-06 08:53:59', NULL);
INSERT INTO `lesson` VALUES (5, 371431001986957313, 2, 370984071263682561, NULL, '2021-09-06 17:15:29', '2021-09-10 09:59:12');
INSERT INTO `lesson` VALUES (6, 371567347703480321, 2, 371032570705477633, NULL, '2021-09-07 15:49:57', '2021-09-06 17:16:17');
INSERT INTO `lesson` VALUES (7, 371966893411532801, 2, 371032570705477633, NULL, '2021-09-10 09:59:05', '2030-01-01 00:00:00');

-- ----------------------------
-- Table structure for lessoninfo
-- ----------------------------
DROP TABLE IF EXISTS `lessoninfo`;
CREATE TABLE `lessoninfo`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `lesson_id` bigint NOT NULL,
  `student_id` bigint NOT NULL,
  `submit_status` tinyint NULL DEFAULT 0,
  `submit_time` datetime NULL DEFAULT '0001-01-01 00:00:00',
  `signin_status` tinyint NOT NULL DEFAULT 0,
  `signin_time` datetime NOT NULL DEFAULT '1900-01-01 00:00:00',
  PRIMARY KEY (`id`, `signin_status`, `signin_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 57 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lessoninfo
-- ----------------------------
INSERT INTO `lessoninfo` VALUES (57, 371567347703480321, 370984071263682561, 0, '0001-01-01 00:00:00', -1, '2021-09-10 13:25:09');

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint NOT NULL COMMENT '帖子id',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` varchar(8192) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `author_id` bigint NOT NULL COMMENT '作者的用户id',
  `community_id` bigint NOT NULL COMMENT '所属社区',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '帖子状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_post_id`(`post_id`) USING BTREE,
  INDEX `idx_author_id`(`author_id`) USING BTREE,
  INDEX `idx_community_id`(`community_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post
-- ----------------------------
INSERT INTO `post` VALUES (1, 370985769638035457, '9.1作业', '书法临摹10篇', 370984071263682561, 1, 1, '2021-09-03 15:32:30', '2021-09-03 15:33:29');
INSERT INTO `post` VALUES (3, 370985915062943745, '9.2作业', '巩固复习，明天考试', 370984071263682561, 1, 1, '2021-09-03 15:33:56', '2021-09-03 15:33:56');
INSERT INTO `post` VALUES (4, 370985943852646401, '9.3作业', '英语听力p96 1、2、3', 370984071263682561, 1, 1, '2021-09-03 15:34:14', '2021-09-03 15:34:14');
INSERT INTO `post` VALUES (11, 371965671409451009, '9.10作业', 'c语言上机测试', 371032570705477633, 1, 1, '2021-09-10 09:46:57', '2021-09-15 22:58:58');
INSERT INTO `post` VALUES (12, 371987092357513217, '9.10作业', '英语听力p110 ', 370984071263682561, 1, 1, '2021-09-10 13:19:45', '2021-09-10 13:19:45');

-- ----------------------------
-- Table structure for score
-- ----------------------------
DROP TABLE IF EXISTS `score`;
CREATE TABLE `score`  (
  `class_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `score` tinyint NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of score
-- ----------------------------
INSERT INTO `score` VALUES (23, 371133821019488257, 88);

-- ----------------------------
-- Table structure for test
-- ----------------------------
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `test_id` bigint NOT NULL,
  `type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `publisher` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `creat_time` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of test
-- ----------------------------
INSERT INTO `test` VALUES (1, 372274927627141121, '思政', 'admin', '2021-09-12 12:59:08');
INSERT INTO `test` VALUES (2, 372895191821975553, '思政', 'admin', '2021-09-16 19:40:54');

-- ----------------------------
-- Table structure for test_score
-- ----------------------------
DROP TABLE IF EXISTS `test_score`;
CREATE TABLE `test_score`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `test_id` bigint NOT NULL,
  `student_id` bigint NOT NULL,
  `student_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `submit_time` datetime NOT NULL,
  `score` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of test_score
-- ----------------------------
INSERT INTO `test_score` VALUES (1, 372270825094512641, 371133821019488257, '小明', '2021-09-12 14:18:52', '100');
INSERT INTO `test_score` VALUES (2, 372270825094512641, 88638472083996673, '小王', '2021-09-12 14:30:51', '90');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `gender` tinyint NOT NULL DEFAULT 0,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `position` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE,
  UNIQUE INDEX `idx_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 370984071263682561, 'admin', '31323334353678f5ef94a851cd855e9e4eabd80d5956', '984434045@qq.com', 0, '2021-09-03 15:15:37', '2021-09-10 20:16:27', 'admin');
INSERT INTO `user` VALUES (2, 88638472083996673, '小王', '31323334353678f5ef94a851cd855e9e4eabd80d5956', NULL, 0, '2021-09-03 19:44:06', '2021-09-12 14:29:24', 'student');
INSERT INTO `user` VALUES (3, 371032570705477633, 'abc', '31323334353678f5ef94a851cd855e9e4eabd80d5956', '', 0, '2021-09-03 23:17:25', '2021-09-10 20:05:58', 'teacher');
INSERT INTO `user` VALUES (4, 371133821019488257, '小明', '31323334353678f5ef94a851cd855e9e4eabd80d5956', NULL, 0, '2021-09-04 16:03:15', '2021-09-16 18:55:33', 'teacher');
INSERT INTO `user` VALUES (10, 372888777120546817, '小红', '31323334353678f5ef94a851cd855e9e4eabd80d5956', '12345678@qq.com', 0, '2021-09-16 18:37:11', '2021-09-16 18:37:11', 'student');
INSERT INTO `user` VALUES (11, 372894636378685441, 'abcdef', '31323334353678f5ef94a851cd855e9e4eabd80d5956', '12345678@qq.com', 0, '2021-09-16 19:35:23', '2021-09-16 19:35:23', 'student');

SET FOREIGN_KEY_CHECKS = 1;
