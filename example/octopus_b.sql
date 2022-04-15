/*
 Navicat Premium Data Transfer

 Source Server         : 0.LOCALHOST
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 127.0.0.1:3306
 Source Schema         : octopus_b

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 19/02/2022 18:19:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for x
-- ----------------------------
DROP TABLE IF EXISTS `x`;
CREATE TABLE `x` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `y_id` int NOT NULL,
  `deleted` tinyint NOT NULL DEFAULT '0',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of x
-- ----------------------------
BEGIN;
INSERT INTO `x` VALUES (1, 'x1', 1, 0, '2022-02-19 09:23:29');
INSERT INTO `x` VALUES (2, 'x2', 1, 0, '2022-02-19 09:23:35');
COMMIT;

-- ----------------------------
-- Table structure for y
-- ----------------------------
DROP TABLE IF EXISTS `y`;
CREATE TABLE `y` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `deleted` tinyint NOT NULL DEFAULT '0',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of y
-- ----------------------------
BEGIN;
INSERT INTO `y` VALUES (1, 'y1', 0, '2022-02-19 09:23:15');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
