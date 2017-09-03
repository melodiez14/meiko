/*
 Navicat Premium Data Transfer

 Source Server         : meiko
 Source Server Type    : MySQL
 Source Server Version : 50635
 Source Host           : localhost:3306
 Source Schema         : meiko

 Target Server Type    : MySQL
 Target Server Version : 50635
 File Encoding         : 65001

 Date: 03/09/2017 13:17:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for assigments
-- ----------------------------
DROP TABLE IF EXISTS `assigments`;
CREATE TABLE `assigments` (
  `id` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL,
  `upload_date` datetime DEFAULT NULL,
  `due_date` datetime NOT NULL,
  `grade_parameter_id` int(11) NOT NULL,
  `description` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_assigments_gradeparameter1_idx` (`grade_parameter_id`),
  CONSTRAINT `fk_assigments_gradeparameter1` FOREIGN KEY (`grade_parameter_id`) REFERENCES `grade_parameters` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for attendances
-- ----------------------------
DROP TABLE IF EXISTS `attendances`;
CREATE TABLE `attendances` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `meeting_number` tinyint(2) DEFAULT NULL,
  `p_users_courses_users_id` varchar(12) NOT NULL,
  `p_users_courses_courses_id` varchar(10) NOT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `meeting_date` date NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_attendances_p_users_courses1_idx` (`p_users_courses_users_id`,`p_users_courses_courses_id`),
  CONSTRAINT `fk_attendances_p_users_courses1` FOREIGN KEY (`p_users_courses_users_id`, `p_users_courses_courses_id`) REFERENCES `p_users_courses` (`users_id`, `courses_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for courses
-- ----------------------------
DROP TABLE IF EXISTS `courses`;
CREATE TABLE `courses` (
  `id` varchar(10) NOT NULL,
  `name` varchar(45) NOT NULL,
  `ucu` tinyint(4) NOT NULL,
  `semester` tinyint(1) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `end_time` time NOT NULL,
  `start_time` time NOT NULL,
  `classes` char(1) NOT NULL,
  `day` enum('Sun','mon','tue','wed','thu','fri','sat') NOT NULL,
  `places_id` varchar(10) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`,`places_id`),
  KEY `fk_courses_places1_idx` (`places_id`),
  CONSTRAINT `fk_courses_places1` FOREIGN KEY (`places_id`) REFERENCES `places` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `path` varchar(45) NOT NULL,
  `mime` varchar(255) DEFAULT NULL,
  `extension` varchar(5) DEFAULT NULL,
  `size` int(4) DEFAULT NULL,
  `upload_by` varchar(45) DEFAULT NULL,
  `users_id` varchar(12) NOT NULL,
  `table_name` varchar(45) DEFAULT NULL,
  `table_id` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_files_users1_idx` (`users_id`),
  CONSTRAINT `fk_files_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for grade_parameters
-- ----------------------------
DROP TABLE IF EXISTS `grade_parameters`;
CREATE TABLE `grade_parameters` (
  `id` int(11) NOT NULL,
  `type` varchar(15) DEFAULT NULL,
  `percentage` int(2) DEFAULT NULL,
  `courses_id` varchar(10) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_gradeparameter_courses1_idx` (`courses_id`),
  CONSTRAINT `fk_gradeparameter_courses1` FOREIGN KEY (`courses_id`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for informations
-- ----------------------------
DROP TABLE IF EXISTS `informations`;
CREATE TABLE `informations` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(45) DEFAULT NULL,
  `description` text,
  `courses_id` varchar(10) NOT NULL,
  `type` enum('General','Material') DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`,`courses_id`),
  KEY `fk_informations_courses1_idx` (`courses_id`),
  CONSTRAINT `fk_informations_courses1` FOREIGN KEY (`courses_id`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for modules
-- ----------------------------
DROP TABLE IF EXISTS `modules`;
CREATE TABLE `modules` (
  `id` int(11) NOT NULL,
  `name` varchar(45) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for notifications
-- ----------------------------
DROP TABLE IF EXISTS `notifications`;
CREATE TABLE `notifications` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `descriptions` text NOT NULL,
  `read_at` timestamp NULL DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `table_id` varchar(45) DEFAULT NULL,
  `table_name` varchar(45) DEFAULT NULL,
  `users_id` varchar(12) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`,`users_id`),
  KEY `fk_notifications_users1_idx` (`users_id`),
  CONSTRAINT `fk_notifications_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for p_rolegroups_modules
-- ----------------------------
DROP TABLE IF EXISTS `p_rolegroups_modules`;
CREATE TABLE `p_rolegroups_modules` (
  `rolegroups_id` int(11) NOT NULL,
  `modules_id` int(11) NOT NULL,
  `ability` enum('create','read','update','delete','xcreate','xread','xupdate','xdelete') NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`rolegroups_id`,`modules_id`,`ability`),
  KEY `fk_role_groups_has_modules_modules1_idx` (`modules_id`),
  KEY `fk_role_groups_has_modules_role_groups1_idx` (`rolegroups_id`),
  CONSTRAINT `fk_role_groups_has_modules_modules1` FOREIGN KEY (`modules_id`) REFERENCES `modules` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_role_groups_has_modules_role_groups1` FOREIGN KEY (`rolegroups_id`) REFERENCES `rolegroups` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for p_users_assignments
-- ----------------------------
DROP TABLE IF EXISTS `p_users_assignments`;
CREATE TABLE `p_users_assignments` (
  `assigments_id` int(11) NOT NULL,
  `users_id` varchar(12) NOT NULL,
  `score` float DEFAULT NULL,
  `description` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`assigments_id`,`users_id`),
  KEY `fk_assigments_has_users_users1_idx` (`users_id`),
  KEY `fk_assigments_has_users_assigments1_idx` (`assigments_id`),
  CONSTRAINT `fk_assigments_has_users_assigments1` FOREIGN KEY (`assigments_id`) REFERENCES `assigments` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_assigments_has_users_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for p_users_courses
-- ----------------------------
DROP TABLE IF EXISTS `p_users_courses`;
CREATE TABLE `p_users_courses` (
  `users_id` varchar(12) NOT NULL,
  `courses_id` varchar(10) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`users_id`,`courses_id`),
  KEY `fk_users_has_lectures_lectures1_idx` (`courses_id`),
  KEY `fk_users_has_lectures_users1_idx` (`users_id`),
  CONSTRAINT `fk_users_has_lectures_lectures1` FOREIGN KEY (`courses_id`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_lectures_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for places
-- ----------------------------
DROP TABLE IF EXISTS `places`;
CREATE TABLE `places` (
  `id` varchar(10) NOT NULL,
  `name` varchar(45) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rolegroups
-- ----------------------------
DROP TABLE IF EXISTS `rolegroups`;
CREATE TABLE `rolegroups` (
  `id` int(11) NOT NULL,
  `name` varchar(15) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` varchar(12) NOT NULL,
  `name` varchar(50) NOT NULL,
  `gender` enum('L','P') NOT NULL,
  `email` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `college` varchar(100) NOT NULL DEFAULT '',
  `note` varchar(100) NOT NULL DEFAULT '',
  `rolegroups_id` int(11) DEFAULT NULL,
  `status` int(1) NOT NULL DEFAULT '0',
  `phone` varchar(14) DEFAULT NULL,
  `line_id` varchar(45) DEFAULT NULL,
  `email_verification_code` int(4) DEFAULT NULL,
  `email_verification_expire_date` datetime DEFAULT NULL,
  `email_verification_attempt` int(1) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email_UNIQUE` (`email`),
  UNIQUE KEY `phone_UNIQUE` (`phone`),
  UNIQUE KEY `line_id_UNIQUE` (`line_id`),
  KEY `fk_users_role_groups_idx` (`rolegroups_id`),
  CONSTRAINT `fk_users_role_groups` FOREIGN KEY (`rolegroups_id`) REFERENCES `rolegroups` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES ('140810140016', 'Risal Falah', 'L', 'risal@live.com', 'c5003e1b1145975bca5630c2ac826173', 'Unpad', '', NULL, 0, NULL, NULL, 7414, '2017-09-03 11:44:07', 0, '2017-09-02 22:32:09', '2017-09-02 22:32:14');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
