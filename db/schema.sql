/*
  Date: 14/10/2017 17:46:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for assigments
-- ----------------------------
DROP TABLE IF EXISTS `assigments`;
CREATE TABLE `assigments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `status` tinyint(3) unsigned NOT NULL,
  `due_date` datetime NOT NULL,
  `grade_parameter_id` int(11) unsigned NOT NULL,
  `description` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_assigments_gradeparameter1` (`grade_parameter_id`),
  CONSTRAINT `fk_assigments_gradeparameter1` FOREIGN KEY (`grade_parameter_id`) REFERENCES `grade_parameters` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for attendances
-- ----------------------------
DROP TABLE IF EXISTS `attendances`;
CREATE TABLE `attendances` (
  `meetings_number` tinyint(3) unsigned NOT NULL,
  `p_users_courses_users_id` int(10) unsigned NOT NULL,
  `p_users_courses_courses_id` int(10) unsigned NOT NULL,
  `status` tinyint(3) unsigned DEFAULT NULL,
  `meeting_date` date NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`meetings_number`,`p_users_courses_users_id`,`p_users_courses_courses_id`),
  KEY `fk_attendances_p_users_courses1` (`p_users_courses_users_id`,`p_users_courses_courses_id`),
  CONSTRAINT `fk_attendances_meetings` FOREIGN KEY (`meetings_number`) REFERENCES `meetings` (`number`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_attendances_p_users_courses1` FOREIGN KEY (`p_users_courses_users_id`, `p_users_courses_courses_id`) REFERENCES `p_users_courses` (`users_id`, `courses_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for courses
-- ----------------------------
DROP TABLE IF EXISTS `courses`;
CREATE TABLE `courses` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `description` text,
  `ucu` tinyint(4) unsigned NOT NULL,
  `semester` tinyint(1) unsigned NOT NULL,
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `start_time` smallint(5) unsigned NOT NULL,
  `end_time` smallint(5) unsigned NOT NULL,
  `classes` char(1) NOT NULL,
  `day` tinyint(1) unsigned NOT NULL,
  `places_id` varchar(30) NOT NULL,
  `created_by` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_courses_places` (`places_id`),
  KEY `fk_courses_users` (`created_by`),
  CONSTRAINT `fk_courses_places` FOREIGN KEY (`places_id`) REFERENCES `places` (`id`) ON UPDATE NO ACTION,
  CONSTRAINT `fk_courses_users` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=126 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of courses
-- ----------------------------
BEGIN;
INSERT INTO `courses` VALUES (1, 'Pemrograman Web', NULL, 3, 3, 1, 60, 800, 'C', 0, 'UXJT1234', 2, '2017-09-30 15:49:47', '2017-10-01 14:04:18');
INSERT INTO `courses` VALUES (2, 'Sistem Informasi', NULL, 3, 3, 1, 600, 720, 'C', 0, 'UDJT103', 2, '2017-09-30 15:51:14', '2017-10-01 13:36:55');
INSERT INTO `courses` VALUES (125, 'Pemrograman Web', NULL, 3, 3, 1, 600, 800, 'C', 0, 'UXJT1234', 2, '2017-10-01 13:34:58', '2017-10-01 13:36:55');
COMMIT;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `path` varchar(45) NOT NULL,
  `mime` varchar(35) NOT NULL,
  `extension` varchar(5) NOT NULL,
  `users_id` int(10) unsigned NOT NULL,
  `type` varchar(10) NOT NULL,
  `table_name` varchar(45) NOT NULL,
  `table_id` varchar(20) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_files_users` (`users_id`),
  CONSTRAINT `fk_files_users` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of files
-- ----------------------------
BEGIN;
INSERT INTO `files` VALUES (1, 'oil-108-self-digital-portrait-painting-18-24', 'private/profile/4.jpg', 'image/jpeg', 'jpg', 4, 'pl', 'users', '4', '2017-09-28 19:46:07', '2017-09-29 00:20:37');
INSERT INTO `files` VALUES (2, 'oil-108-self-digital-portrait-painting-18-24', 'private/profile/t_4.jpg', 'image/jpeg', 'jpg', 4, 'pl_t', 'users', '4', '2017-09-28 19:46:07', '2017-09-29 00:20:37');
COMMIT;

-- ----------------------------
-- Table structure for grade_parameters
-- ----------------------------
DROP TABLE IF EXISTS `grade_parameters`;
CREATE TABLE `grade_parameters` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(15) DEFAULT NULL,
  `percentage` float(5,2) unsigned DEFAULT NULL,
  `courses_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_grade_parameters_courses1` (`courses_id`),
  CONSTRAINT `fk_grade_parameters_courses1` FOREIGN KEY (`courses_id`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for informations
-- ----------------------------
DROP TABLE IF EXISTS `informations`;
CREATE TABLE `informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(45) DEFAULT NULL,
  `description` text,
  `courses_id` int(10) unsigned DEFAULT NULL,
  `type` tinyint(1) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_informations_courses1_idx` (`courses_id`) USING BTREE,
  CONSTRAINT `fk_informations_courses1` FOREIGN KEY (`courses_id`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of informations
-- ----------------------------
BEGIN;
INSERT INTO `informations` VALUES (1, 'Libur', NULL, NULL, 0, '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `informations` VALUES (2, 'Mantap jiwa', 'Joss', NULL, 0, '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `informations` VALUES (3, 'Jess tak', NULL, 1, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `informations` VALUES (4, 'tas duk', '123123123@#%^&^%$3', 1, 1, '2017-10-01 12:54:16', '2017-10-01 12:54:20');
INSERT INTO `informations` VALUES (5, 'asdfmmmgf', 'sdfjnasdkfjnasdfkj', 1, 1, '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `informations` VALUES (6, 'asdfknjkndf', 'njkfuweriuwre', NULL, 0, '2017-10-01 12:58:00', '2017-10-01 12:58:01');
COMMIT;

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
  PRIMARY KEY (`id`),
  KEY `fk_inventories_places` (`places_id`),
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
  PRIMARY KEY (`id`),
  KEY `fk_log_books_researches` (`researches_id`),
  CONSTRAINT `fk_log_books_researches` FOREIGN KEY (`researches_id`) REFERENCES `researches` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for meetings
-- ----------------------------
DROP TABLE IF EXISTS `meetings`;
CREATE TABLE `meetings` (
  `number` tinyint(3) unsigned NOT NULL,
  `courses_id` int(10) unsigned NOT NULL,
  `description` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`number`,`courses_id`),
  KEY `fk_meetings_courses` (`courses_id`),
  CONSTRAINT `fk_meetings_courses` FOREIGN KEY (`courses_id`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  PRIMARY KEY (`id`),
  KEY `fk_notifications_users1_idx` (`users_id`) USING BTREE,
  CONSTRAINT `fk_notifications_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for p_users_assignments
-- ----------------------------
DROP TABLE IF EXISTS `p_users_assignments`;
CREATE TABLE `p_users_assignments` (
  `assigments_id` int(10) unsigned NOT NULL,
  `users_id` int(10) unsigned NOT NULL,
  `courses_id` int(10) unsigned NOT NULL,
  `score` float(5,2) unsigned DEFAULT NULL,
  `description` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`assigments_id`,`users_id`,`courses_id`),
  KEY `fk_users_assignments_users_courses` (`users_id`,`courses_id`),
  CONSTRAINT `fk_users_assignments_assignments` FOREIGN KEY (`assigments_id`) REFERENCES `assigments` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_assignments_users_courses` FOREIGN KEY (`users_id`, `courses_id`) REFERENCES `p_users_courses` (`users_id`, `courses_id`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for p_users_courses
-- ----------------------------
DROP TABLE IF EXISTS `p_users_courses`;
CREATE TABLE `p_users_courses` (
  `users_id` int(10) unsigned NOT NULL,
  `courses_id` int(10) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`users_id`,`courses_id`),
  KEY `fk_users_has_lectures_lectures1` (`courses_id`),
  CONSTRAINT `fk_users_has_lectures_lectures1` FOREIGN KEY (`courses_id`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_lectures_users1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of p_users_courses
-- ----------------------------
BEGIN;
INSERT INTO `p_users_courses` VALUES (2, 1, 1, '2017-10-01 11:21:46', '2017-10-01 11:21:50');
INSERT INTO `p_users_courses` VALUES (4, 1, 0, '2017-10-01 12:02:09', '2017-10-01 12:02:10');
COMMIT;

-- ----------------------------
-- Table structure for p_users_researches
-- ----------------------------
DROP TABLE IF EXISTS `p_users_researches`;
CREATE TABLE `p_users_researches` (
  `users_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `researches_id` int(10) unsigned NOT NULL,
  `sort` tinyint(1) unsigned NOT NULL COMMENT 'sorting the researchers',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`users_id`,`researches_id`),
  KEY `fk_p_users_researches_researches` (`researches_id`),
  CONSTRAINT `fk_p_users_researches_researches` FOREIGN KEY (`researches_id`) REFERENCES `researches` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_p_users_researches_users` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for places
-- ----------------------------
DROP TABLE IF EXISTS `places`;
CREATE TABLE `places` (
  `id` varchar(30) NOT NULL,
  `description` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of places
-- ----------------------------
BEGIN;
INSERT INTO `places` VALUES ('UDJT102', 'Laboratorium AIO', '2017-09-30 15:47:11', '2017-09-30 15:47:15');
INSERT INTO `places` VALUES ('UDJT103', 'Laboratorium', '2017-09-30 15:47:40', '2017-09-30 15:47:44');
INSERT INTO `places` VALUES ('UXJT1234', NULL, '2017-10-01 13:34:58', '2017-10-01 13:34:58');
COMMIT;

-- ----------------------------
-- Table structure for research_categories
-- ----------------------------
DROP TABLE IF EXISTS `research_categories`;
CREATE TABLE `research_categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

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
  PRIMARY KEY (`id`),
  KEY `fk_researches_research_categories` (`research_categories_id`),
  KEY `fk_researches_places` (`places_id`),
  CONSTRAINT `fk_researches_places` FOREIGN KEY (`places_id`) REFERENCES `places` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_researches_research_categories` FOREIGN KEY (`research_categories_id`) REFERENCES `research_categories` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for rolegroups
-- ----------------------------
DROP TABLE IF EXISTS `rolegroups`;
CREATE TABLE `rolegroups` (
  `id` int(10) unsigned NOT NULL,
  `name` varchar(15) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of rolegroups
-- ----------------------------
BEGIN;
INSERT INTO `rolegroups` VALUES (1, 'Assistant', '2017-09-28 18:48:29', '2017-09-28 18:48:31');
INSERT INTO `rolegroups` VALUES (2, 'Lecturer', '2017-09-28 18:48:50', '2017-09-28 18:48:52');
COMMIT;

-- ----------------------------
-- Table structure for rolegroups_modules
-- ----------------------------
DROP TABLE IF EXISTS `rolegroups_modules`;
CREATE TABLE `rolegroups_modules` (
  `rolegroups_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `modules` enum('users','courses','attendances','roles') NOT NULL,
  `ability` enum('CREATE','READ','UPDATE','DELETE','XCREATE','XREAD','XUPDATE','XDELETE') NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`rolegroups_id`,`modules`,`ability`),
  KEY `fk_role_groups_has_modules_role_groups1_idx` (`rolegroups_id`) USING BTREE,
  CONSTRAINT `fk_role_groups_has_modules_role_groups1` FOREIGN KEY (`rolegroups_id`) REFERENCES `rolegroups` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of rolegroups_modules
-- ----------------------------
BEGIN;
INSERT INTO `rolegroups_modules` VALUES (2, 'users', 'UPDATE', '2017-09-28 18:50:09', '2017-09-28 18:50:11');
INSERT INTO `rolegroups_modules` VALUES (2, 'users', 'XCREATE', '2017-09-28 18:49:36', '2017-09-28 18:49:38');
INSERT INTO `rolegroups_modules` VALUES (2, 'users', 'XREAD', '2017-09-28 18:49:57', '2017-09-28 18:49:59');
INSERT INTO `rolegroups_modules` VALUES (2, 'users', 'XDELETE', '2017-09-28 18:50:29', '2017-09-28 18:50:37');
INSERT INTO `rolegroups_modules` VALUES (2, 'courses', 'UPDATE', '2017-10-01 13:52:34', '2017-10-01 13:52:35');
INSERT INTO `rolegroups_modules` VALUES (2, 'courses', 'XCREATE', '2017-10-01 07:20:20', '2017-10-01 07:20:22');
INSERT INTO `rolegroups_modules` VALUES (2, 'courses', 'XREAD', '2017-09-30 17:26:27', '2017-09-30 17:26:29');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `gender` tinyint(3) unsigned NOT NULL,
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
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_users_email` (`email`),
  UNIQUE KEY `unique_users_identity_code` (`identity_code`),
  KEY `fk_users_role_groups` (`rolegroups_id`),
  CONSTRAINT `fk_users_role_groups` FOREIGN KEY (`rolegroups_id`) REFERENCES `rolegroups` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (2, 'Risal Bro', 0, 'risal.falah@gmail.com', '2af9b1ba42dc5eb01743e6b3759b6e4b', '', 2, 2, NULL, NULL, '140810140090', 3923, '2017-10-09 20:07:27', NULL, '2017-09-28 18:15:28', '2017-10-09 19:37:27');
INSERT INTO `users` VALUES (4, 'Bro Risal', 1, 'risal@live.com', '2af9b1ba42dc5eb01743e6b3759b6e4b', 'Hello im risal falah', NULL, 2, '085860141146', 'risalf', '140810140016', 3513, '2017-09-28 20:10:04', 0, '2017-09-28 19:38:23', '2017-09-28 19:54:08');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
