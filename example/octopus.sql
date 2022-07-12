/*
 Navicat Premium Data Transfer

 Source Server         : 0.LOCALHOST
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 127.0.0.1:3306
 Source Schema         : octopus

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 19/02/2022 18:18:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for claw_config
-- ----------------------------
DROP TABLE IF EXISTS `claw_config`;
CREATE TABLE `claw_config` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `source` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '数据来源',
  `target_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'ELASTIC_SEARCH' COMMENT 'MYSQL/ELASTIC_SEARCH',
  `target_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '目标数据源',
  `target_set` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '数据同步目标地',
  `comments` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='同步配置';

-- ----------------------------
-- Records of claw_config
-- ----------------------------
BEGIN;
INSERT INTO `claw_config` (`id`, `name`, `source`, `target_type`, `target_source`, `target_set`, `comments`, `deleted`, `created_time`, `updated_time`) VALUES (1, 'example', '{\"sourceType\":\"MYSQL\",\"dataSource\":\"octopus_a\",\"dataSet\":\"a\",\"priority\":0,\"subscribed\":true,\"idColumn\":\"id\",\"timeColumn\":\"updated_time\",\"deleteColumn\":\"deleted\",\"fields\":[{\"column\":\"id\",\"target\":\"id\"},{\"column\":\"name\",\"target\":\"name\"}],\"relations\":[{\"fields\":[{\"column\":\"id\",\"target\":\"a_id\"}],\"sourceTree\":{\"sourceType\":\"MYSQL\",\"dataSource\":\"octopus_a\",\"dataSet\":\"b\",\"priority\":0,\"subscribed\":true,\"idColumn\":\"id\",\"timeColumn\":\"updated_time\",\"deleteColumn\":\"deleted\",\"fields\":[{\"column\":\"id\",\"target\":\"b_id\"},{\"column\":\"name\",\"target\":\"b_name\"}],\"relations\":[{\"fields\":[{\"column\":\"id\",\"target\":\"b_id\"}],\"sourceTree\":{\"sourceType\":\"MYSQL\",\"dataSource\":\"octopus_a\",\"dataSet\":\"c\",\"priority\":0,\"subscribed\":true,\"idColumn\":\"id\",\"timeColumn\":\"updated_time\",\"deleteColumn\":\"deleted\",\"fields\":[{\"column\":\"id\",\"target\":\"c_id\"},{\"column\":\"name\",\"target\":\"c_name\"}]}},{\"fields\":[{\"column\":\"id\",\"target\":\"b_id\"}],\"sourceTree\":{\"sourceType\":\"MYSQL\",\"dataSource\":\"octopus_a\",\"dataSet\":\"d\",\"priority\":0,\"subscribed\":true,\"idColumn\":\"id\",\"timeColumn\":\"updated_time\",\"deleteColumn\":\"deleted\",\"fields\":[{\"column\":\"id\",\"target\":\"d_id\"},{\"column\":\"name\",\"target\":\"d_name\"}]}}]}},{\"fields\":[{\"column\":\"x_id\",\"target\":\"id\"}],\"sourceTree\":{\"sourceType\":\"MYSQL\",\"dataSource\":\"octopus_b\",\"dataSet\":\"x\",\"priority\":0,\"subscribed\":true,\"idColumn\":\"id\",\"timeColumn\":\"updated_time\",\"deleteColumn\":\"deleted\",\"fields\":[{\"column\":\"id\",\"target\":\"x_id\"},{\"column\":\"name\",\"target\":\"x_name\"}],\"relations\":[{\"fields\":[{\"column\":\"y_id\",\"target\":\"id\"}],\"sourceTree\":{\"sourceType\":\"MYSQL\",\"dataSource\":\"octopus_b\",\"dataSet\":\"y\",\"priority\":0,\"subscribed\":true,\"idColumn\":\"id\",\"timeColumn\":\"updated_time\",\"deleteColumn\":\"deleted\",\"fields\":[{\"column\":\"id\",\"target\":\"y_id\"},{\"column\":\"name\",\"target\":\"y_name\"}]}}]}}]}', 'ELASTIC_SEARCH', 'local_es', 'demo_index_v1', '', 0, '2021-12-05 03:21:08', '2022-03-11 16:05:26');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
