/*
 Navicat Premium Data Transfer

 Source Server         : Aliyun
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : 47.74.149.190:3306
 Source Schema         : meiko

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : 65001

 Date: 21/12/2017 07:55:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for assignments
-- ----------------------------
DROP TABLE IF EXISTS `assignments`;
CREATE TABLE `assignments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `description` text,
  `status` tinyint(3) unsigned NOT NULL,
  `due_date` datetime NOT NULL,
  `grade_parameters_id` int(11) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `max_size` int(11) DEFAULT NULL,
  `type` varchar(0) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_assigments_gradeparameter1` (`grade_parameters_id`) USING BTREE,
  CONSTRAINT `fk_assigments_gradeparameter1` FOREIGN KEY (`grade_parameters_id`) REFERENCES `grade_parameters` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=9999123 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for attendances
-- ----------------------------
DROP TABLE IF EXISTS `attendances`;
CREATE TABLE `attendances` (
  `meetings_id` int(10) unsigned NOT NULL,
  `users_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`meetings_id`,`users_id`) USING BTREE,
  KEY `fk_attendances_p_users_schedules` (`users_id`) USING BTREE,
  CONSTRAINT `fk_attendances_meetings` FOREIGN KEY (`meetings_id`) REFERENCES `meetings` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_attendances_users` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for bot_logs
-- ----------------------------
DROP TABLE IF EXISTS `bot_logs`;
CREATE TABLE `bot_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `message` text NOT NULL,
  `users_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `status` tinyint(1) unsigned NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2560 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for courses
-- ----------------------------
DROP TABLE IF EXISTS `courses`;
CREATE TABLE `courses` (
  `id` varchar(40) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text,
  `ucu` tinyint(2) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` varchar(30) NOT NULL,
  `name` varchar(45) NOT NULL,
  `mime` varchar(100) NOT NULL,
  `extension` varchar(5) NOT NULL,
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '0 = deleted, 1 = exist',
  `users_id` int(10) unsigned NOT NULL,
  `type` varchar(10) DEFAULT NULL,
  `table_name` varchar(45) DEFAULT NULL,
  `table_id` varchar(20) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_files_users` (`users_id`) USING BTREE,
  CONSTRAINT `fk_files_users` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for grade_parameters
-- ----------------------------
DROP TABLE IF EXISTS `grade_parameters`;
CREATE TABLE `grade_parameters` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` enum('QUIZ','ASSIGNMENT','FINAL','MID','ATTENDANCE') NOT NULL,
  `percentage` float(5,2) unsigned NOT NULL,
  `status_change` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `schedules_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uq_grade_parameter` (`type`,`schedules_id`) USING BTREE,
  KEY `fk_grade_parameters_courses1` (`schedules_id`) USING BTREE,
  CONSTRAINT `fk_grade_parameters_schedules` FOREIGN KEY (`schedules_id`) REFERENCES `schedules` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=1231232 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for informations
-- ----------------------------
DROP TABLE IF EXISTS `informations`;
CREATE TABLE `informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(45) DEFAULT NULL,
  `description` text,
  `schedules_id` int(10) unsigned DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_informations_courses1_idx` (`schedules_id`) USING BTREE,
  CONSTRAINT `fk_informations_courses1` FOREIGN KEY (`schedules_id`) REFERENCES `schedules` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for inventories
-- ----------------------------
DROP TABLE IF EXISTS `inventories`;
CREATE TABLE `inventories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `condition` tinyint(3) unsigned NOT NULL,
  `quantity` mediumint(8) unsigned DEFAULT '0',
  `note` text,
  `places_id` varchar(30) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_inventories_places` (`places_id`) USING BTREE,
  CONSTRAINT `fk_inventories_places` FOREIGN KEY (`places_id`) REFERENCES `places` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for log_books
-- ----------------------------
DROP TABLE IF EXISTS `log_books`;
CREATE TABLE `log_books` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `note` text NOT NULL,
  `researches_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_log_books_researches` (`researches_id`) USING BTREE,
  CONSTRAINT `fk_log_books_researches` FOREIGN KEY (`researches_id`) REFERENCES `researches` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for meetings
-- ----------------------------
DROP TABLE IF EXISTS `meetings`;
CREATE TABLE `meetings` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `number` tinyint(3) unsigned NOT NULL,
  `subject` varchar(255) NOT NULL,
  `description` text,
  `date` datetime NOT NULL,
  `schedules_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_meetings_schedules` (`schedules_id`) USING BTREE,
  KEY `uq_meetings` (`number`,`schedules_id`) USING BTREE,
  CONSTRAINT `fk_meetings_schedules` FOREIGN KEY (`schedules_id`) REFERENCES `schedules` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for notifications
-- ----------------------------
DROP TABLE IF EXISTS `notifications`;
CREATE TABLE `notifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `descriptions` text NOT NULL,
  `read_at` timestamp NULL DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `table_id` varchar(45) DEFAULT NULL,
  `table_name` varchar(45) DEFAULT NULL,
  `users_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_notifications_users1_idx` (`users_id`) USING BTREE,
  CONSTRAINT `fk_notifications_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for p_users_assignments
-- ----------------------------
DROP TABLE IF EXISTS `p_users_assignments`;
CREATE TABLE `p_users_assignments` (
  `assignments_id` int(10) unsigned NOT NULL,
  `users_id` int(10) unsigned NOT NULL,
  `score` float(5,2) unsigned DEFAULT NULL,
  `description` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`assignments_id`,`users_id`) USING BTREE,
  KEY `fk_users_assignments_users_courses` (`users_id`) USING BTREE,
  CONSTRAINT `fk_users_assignments_assignments` FOREIGN KEY (`assignments_id`) REFERENCES `assignments` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_p_users_assignments` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for p_users_researches
-- ----------------------------
DROP TABLE IF EXISTS `p_users_researches`;
CREATE TABLE `p_users_researches` (
  `users_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `researches_id` int(10) unsigned NOT NULL,
  `sort` tinyint(1) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`users_id`,`researches_id`) USING BTREE,
  KEY `fk_p_users_researches_researches` (`researches_id`) USING BTREE,
  CONSTRAINT `fk_p_users_researches_researches` FOREIGN KEY (`researches_id`) REFERENCES `researches` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_p_users_researches_users` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for p_users_schedules
-- ----------------------------
DROP TABLE IF EXISTS `p_users_schedules`;
CREATE TABLE `p_users_schedules` (
  `users_id` int(10) unsigned NOT NULL,
  `schedules_id` int(10) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`users_id`,`schedules_id`) USING BTREE,
  KEY `fk_users_has_lectures_lectures1` (`schedules_id`) USING BTREE,
  CONSTRAINT `fk_users_has_lectures_lectures1` FOREIGN KEY (`schedules_id`) REFERENCES `schedules` (`id`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_lectures_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for places
-- ----------------------------
DROP TABLE IF EXISTS `places`;
CREATE TABLE `places` (
  `id` varchar(30) NOT NULL,
  `description` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for research_categories
-- ----------------------------
DROP TABLE IF EXISTS `research_categories`;
CREATE TABLE `research_categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for researches
-- ----------------------------
DROP TABLE IF EXISTS `researches`;
CREATE TABLE `researches` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(150) NOT NULL,
  `description` text,
  `output` text,
  `places_id` varchar(30) NOT NULL,
  `research_categories_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_researches_research_categories` (`research_categories_id`) USING BTREE,
  KEY `fk_researches_places` (`places_id`) USING BTREE,
  CONSTRAINT `fk_researches_places` FOREIGN KEY (`places_id`) REFERENCES `places` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_researches_research_categories` FOREIGN KEY (`research_categories_id`) REFERENCES `research_categories` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rolegroups
-- ----------------------------
DROP TABLE IF EXISTS `rolegroups`;
CREATE TABLE `rolegroups` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(15) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for rolegroups_modules
-- ----------------------------
DROP TABLE IF EXISTS `rolegroups_modules`;
CREATE TABLE `rolegroups_modules` (
  `rolegroups_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `module` enum('users','courses','attendances','roles','schedules','assignments','informations','tutorials') NOT NULL,
  `ability` enum('CREATE','READ','UPDATE','DELETE','XCREATE','XREAD','XUPDATE','XDELETE') NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`rolegroups_id`,`module`,`ability`) USING BTREE,
  KEY `fk_role_groups_has_modules_role_groups1_idx` (`rolegroups_id`) USING BTREE,
  CONSTRAINT `fk_rolegroups_modules_rolegroups` FOREIGN KEY (`rolegroups_id`) REFERENCES `rolegroups` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for schedules
-- ----------------------------
DROP TABLE IF EXISTS `schedules`;
CREATE TABLE `schedules` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `start_time` smallint(5) unsigned NOT NULL,
  `end_time` smallint(5) unsigned NOT NULL,
  `day` tinyint(1) unsigned NOT NULL,
  `class` char(1) NOT NULL,
  `semester` tinyint(2) NOT NULL,
  `year` smallint(4) unsigned NOT NULL,
  `courses_id` varchar(40) NOT NULL,
  `places_id` varchar(30) NOT NULL,
  `created_by` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uq_schedules` (`semester`,`year`,`courses_id`,`class`) USING BTREE,
  KEY `fk_courses_places` (`places_id`) USING BTREE,
  KEY `fk_courses_users` (`created_by`) USING BTREE,
  KEY `fk_schedules_courses` (`courses_id`) USING BTREE,
  CONSTRAINT `fk_courses_places` FOREIGN KEY (`places_id`) REFERENCES `places` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_courses_users` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_schedules_courses` FOREIGN KEY (`courses_id`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=100193 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for tutorials
-- ----------------------------
DROP TABLE IF EXISTS `tutorials`;
CREATE TABLE `tutorials` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  `schedules_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_tutorials_schedules` (`schedules_id`) USING BTREE,
  CONSTRAINT `fk_tutorials_schedules` FOREIGN KEY (`schedules_id`) REFERENCES `schedules` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `gender` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `email` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `note` varchar(100) NOT NULL DEFAULT '',
  `rolegroups_id` int(11) unsigned DEFAULT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `phone` varchar(14) DEFAULT NULL,
  `line_id` varchar(45) DEFAULT NULL,
  `identity_code` varchar(18) NOT NULL,
  `email_verification_code` smallint(4) unsigned DEFAULT NULL,
  `email_verification_expire_date` datetime DEFAULT NULL,
  `email_verification_attempt` tinyint(1) unsigned DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `unique_users_email` (`email`) USING BTREE,
  UNIQUE KEY `unique_users_identity_code` (`identity_code`) USING BTREE,
  KEY `fk_users_role_groups` (`rolegroups_id`) USING BTREE,
  CONSTRAINT `fk_users_role_groups` FOREIGN KEY (`rolegroups_id`) REFERENCES `rolegroups` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=2000000005 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
