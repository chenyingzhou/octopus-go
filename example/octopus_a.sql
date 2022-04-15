/*
 Navicat Premium Data Transfer

 Source Server         : 0.LOCALHOST
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 127.0.0.1:3306
 Source Schema         : octopus_a

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 19/02/2022 18:19:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for a
-- ----------------------------
DROP TABLE IF EXISTS `a`;
CREATE TABLE `a` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `x_id` int NOT NULL DEFAULT '0',
  `deleted` tinyint NOT NULL DEFAULT '0',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of a
-- ----------------------------
BEGIN;
INSERT INTO `a` VALUES (1, 'a1', 1, 0, '2022-02-19 09:24:06');
INSERT INTO `a` VALUES (2, 'a2', 1, 0, '2022-02-19 09:24:29');
INSERT INTO `a` VALUES (3, 'a3', 2, 0, '2022-02-19 09:24:26');
COMMIT;

-- ----------------------------
-- Table structure for b
-- ----------------------------
DROP TABLE IF EXISTS `b`;
CREATE TABLE `b` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `a_id` int NOT NULL,
  `deleted` tinyint NOT NULL DEFAULT '0',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of b
-- ----------------------------
BEGIN;
INSERT INTO `b` VALUES (1, 'b1', 1, 0, '2022-02-19 09:24:48');
INSERT INTO `b` VALUES (2, 'b2', 1, 0, '2022-02-19 09:24:55');
INSERT INTO `b` VALUES (3, 'b3', 2, 0, '2022-02-19 09:25:02');
INSERT INTO `b` VALUES (4, 'b4', 2, 0, '2022-02-19 09:25:09');
INSERT INTO `b` VALUES (5, 'b5', 3, 0, '2022-02-19 09:25:18');
COMMIT;

-- ----------------------------
-- Table structure for c
-- ----------------------------
DROP TABLE IF EXISTS `c`;
CREATE TABLE `c` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `b_id` int NOT NULL,
  `deleted` tinyint NOT NULL DEFAULT '0',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of c
-- ----------------------------
BEGIN;
INSERT INTO `c` VALUES (1, 'c1', 1, 0, '2022-02-19 09:25:36');
INSERT INTO `c` VALUES (2, 'c2', 1, 0, '2022-02-19 09:25:43');
INSERT INTO `c` VALUES (3, 'c3', 2, 0, '2022-02-19 09:25:49');
INSERT INTO `c` VALUES (4, 'c4', 3, 0, '2022-02-19 09:26:01');
INSERT INTO `c` VALUES (5, 'c5', 4, 0, '2022-02-19 09:26:09');
INSERT INTO `c` VALUES (6, 'c6', 5, 0, '2022-02-19 09:26:15');
COMMIT;

-- ----------------------------
-- Table structure for d
-- ----------------------------
DROP TABLE IF EXISTS `d`;
CREATE TABLE `d` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `b_id` int NOT NULL,
  `deleted` tinyint NOT NULL DEFAULT '0',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of d
-- ----------------------------
BEGIN;
INSERT INTO `d` VALUES (1, 'd1', 1, 0, '2022-02-19 09:26:29');
INSERT INTO `d` VALUES (2, 'd2', 2, 0, '2022-02-19 09:26:35');
INSERT INTO `d` VALUES (3, 'd3', 3, 0, '2022-02-19 09:26:42');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
