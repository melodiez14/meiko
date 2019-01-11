/*
 Navicat Premium Data Transfer

 Source Server         : main-local-conection
 Source Server Type    : MySQL
 Source Server Version : 50718
 Source Host           : localhost:3306
 Source Schema         : meiko

 Target Server Type    : MySQL
 Target Server Version : 50718
 File Encoding         : 65001

 Date: 12/01/2019 01:17:32
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
  `max_size` tinyint(3) unsigned DEFAULT NULL,
  `max_file` tinyint(3) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_assigments_gradeparameter1` (`grade_parameters_id`) USING BTREE,
  CONSTRAINT `fk_assigments_gradeparameter1` FOREIGN KEY (`grade_parameters_id`) REFERENCES `grade_parameters` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of assignments
-- ----------------------------
BEGIN;
INSERT INTO `assignments` VALUES (1, 'Tugas 1 - Matematika Dasar', 'Kerjakan soal exercise 1 pada buku matematika dasar 1, kemudian kirim tulis di word dan upload filenya', 1, '2019-01-03 11:54:28', 4, '2018-01-18 08:45:32', '2018-01-18 08:45:36', 1, 1);
INSERT INTO `assignments` VALUES (2, 'Tugas 1 - Bahasa Indonesia', 'Buat puisi yang bertemakan nasional kemudian upload', 1, '2018-02-02 08:47:08', 4, '2018-01-18 08:47:17', '2018-01-18 08:47:18', 1, 1);
INSERT INTO `assignments` VALUES (3, 'Tugas 1 - Fisika', 'Kerjakan pilihan ganda exercise 1 kemudian upload', 1, '2019-01-10 00:54:15', 1, '2018-01-18 08:54:06', '2018-01-18 08:54:08', 1, 1);
INSERT INTO `assignments` VALUES (5, 'Tugas 2 - Matematika ', 'Kerjakan soal execise 2 pada buku matematika dasar kemudian upload', 1, '2019-01-15 00:55:20', 4, '2018-02-10 11:08:39', '2018-02-10 11:08:39', 0, 1);
INSERT INTO `assignments` VALUES (6, '', 'Membuat makalah sederhana', 1, '2018-12-18 12:00:00', 4, '2018-02-10 11:08:59', '2018-02-10 11:08:59', 0, 1);
INSERT INTO `assignments` VALUES (7, 'Midtest 1 Take Home', NULL, 0, '2018-02-28 12:12:00', 5, '2018-02-10 11:43:05', '2018-02-10 11:43:05', NULL, NULL);
INSERT INTO `assignments` VALUES (9, 'Ujian Akhir', NULL, 0, '2018-02-20 12:00:00', 6, '2018-02-10 14:47:55', '2018-02-10 14:47:55', NULL, NULL);
COMMIT;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of attendances
-- ----------------------------
BEGIN;
INSERT INTO `attendances` VALUES (1, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (1, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (1, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (1, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (1, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (2, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (2, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (2, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (2, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (2, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (3, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (3, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (3, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (3, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (3, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (4, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (4, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (4, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (4, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (4, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (6, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (6, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (6, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (6, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (6, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (7, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (7, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (7, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (7, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (7, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (8, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (8, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (8, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (8, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (8, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (9, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (9, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (9, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (9, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (9, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (10, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (10, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (10, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (10, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (10, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (11, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (11, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (11, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (11, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (11, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (12, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (12, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (12, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (12, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (12, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (13, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (13, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (13, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (13, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (13, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (14, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (14, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (14, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (14, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (14, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (15, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (15, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (15, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (15, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (15, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (16, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (16, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (16, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (16, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (16, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (17, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (17, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (17, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (17, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (17, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (18, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (18, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (18, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (18, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (18, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (19, 1, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (19, 2, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (19, 3, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (19, 4, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
INSERT INTO `attendances` VALUES (19, 5, '2019-01-12 01:12:16', '2019-01-12 01:12:16');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=306 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of bot_logs
-- ----------------------------
BEGIN;
INSERT INTO `bot_logs` VALUES (83, 'jadwal hari ini', 4, '2018-01-18 04:38:16', 0);
INSERT INTO `bot_logs` VALUES (84, '{\"entity\":[{\"course_name\":\"Struktur Data\",\"day\":\"Thursday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\"}', 4, '2018-01-18 04:38:16', 1);
INSERT INTO `bot_logs` VALUES (85, 'asisten yang ngajar hari ini siapa sih?', 4, '2018-01-18 04:38:29', 0);
INSERT INTO `bot_logs` VALUES (86, '{\"entity\":[{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199666274331738.434554.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199848912061665.429961.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\"}', 4, '2018-01-18 04:38:29', 1);
INSERT INTO `bot_logs` VALUES (87, 'informasi hari ini', 4, '2018-01-18 06:13:12', 0);
INSERT INTO `bot_logs` VALUES (88, '{\"entity\":[],\"intent\":\"information\"}', 4, '2018-01-18 06:13:12', 1);
INSERT INTO `bot_logs` VALUES (89, 'informasi', 4, '2018-01-18 06:13:15', 0);
INSERT INTO `bot_logs` VALUES (90, '{\"entity\":[{\"description\":\"Ini adalah desc informasi Informasi 17 Januari 2018\",\"id\":2,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197473,\"title\":\"Informasi Informasi 18 Januari 2018\"},{\"description\":\"Ini adalah deskripsi informasi tanggal Informasi 17 Januari 2018\",\"id\":1,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197440,\"title\":\"Informasi 17 Januari 2018\"}],\"intent\":\"information\"}', 4, '2018-01-18 06:13:15', 1);
INSERT INTO `bot_logs` VALUES (91, 'informasi kemaring', 4, '2018-01-18 06:13:21', 0);
INSERT INTO `bot_logs` VALUES (92, '{\"entity\":[{\"description\":\"Ini adalah desc informasi Informasi 17 Januari 2018\",\"id\":2,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197473,\"title\":\"Informasi Informasi 18 Januari 2018\"},{\"description\":\"Ini adalah deskripsi informasi tanggal Informasi 17 Januari 2018\",\"id\":1,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197440,\"title\":\"Informasi 17 Januari 2018\"}],\"intent\":\"information\"}', 4, '2018-01-18 06:13:21', 1);
INSERT INTO `bot_logs` VALUES (93, 'asisten yang ngajar hari ini', 4, '2018-01-18 06:57:39', 0);
INSERT INTO `bot_logs` VALUES (94, '{\"entity\":[{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199666274331738.434554.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199848912061665.429961.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\"}', 4, '2018-01-18 06:57:39', 1);
INSERT INTO `bot_logs` VALUES (95, 'bro kasih tau dong asisten yang ngajar saya', 4, '2018-01-18 06:58:14', 0);
INSERT INTO `bot_logs` VALUES (96, '{\"entity\":[{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199848912061665.429961.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199751435760215.217800.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199927657053139.403577.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199124102442230.279765.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199666274331738.434554.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199556090164418.385200.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"}],\"intent\":\"assistant\"}', 4, '2018-01-18 06:58:14', 1);
INSERT INTO `bot_logs` VALUES (97, 'bro kasih dong informasi hari ini', 4, '2018-01-18 07:02:40', 0);
INSERT INTO `bot_logs` VALUES (98, '{\"entity\":[],\"intent\":\"information\"}', 4, '2018-01-18 07:02:40', 1);
INSERT INTO `bot_logs` VALUES (99, 'informasi kemarin', 4, '2018-01-18 07:02:48', 0);
INSERT INTO `bot_logs` VALUES (100, '{\"entity\":[{\"description\":\"Ini adalah desc informasi Informasi 17 Januari 2018\",\"id\":2,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197473,\"title\":\"Informasi Informasi 18 Januari 2018\"},{\"description\":\"Ini adalah deskripsi informasi tanggal Informasi 17 Januari 2018\",\"id\":1,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197440,\"title\":\"Informasi 17 Januari 2018\"}],\"intent\":\"information\"}', 4, '2018-01-18 07:02:48', 1);
INSERT INTO `bot_logs` VALUES (101, 'jadwal saya', 4, '2018-01-18 07:03:03', 0);
INSERT INTO `bot_logs` VALUES (102, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"day\":\"Wednesday\",\"place\":\"UDJT-0206\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Metode Numerik\",\"day\":\"Monday\",\"place\":\"UDJT-0207\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Struktur Data\",\"day\":\"Thursday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\"}', 4, '2018-01-18 07:03:03', 1);
INSERT INTO `bot_logs` VALUES (103, 'jadwal kemarin', 4, '2018-01-18 07:03:28', 0);
INSERT INTO `bot_logs` VALUES (104, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"day\":\"Wednesday\",\"place\":\"UDJT-0206\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\"}', 4, '2018-01-18 07:03:28', 1);
INSERT INTO `bot_logs` VALUES (105, 'tugas yang belum dikerjain', 4, '2018-01-18 07:03:55', 0);
INSERT INTO `bot_logs` VALUES (106, '{\"entity\":[{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1516236834,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"},{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat rancangan 2\",\"due_date\":1516409228,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/2\"}],\"intent\":\"assignment\"}', 4, '2018-01-18 07:03:55', 1);
INSERT INTO `bot_logs` VALUES (107, 'nilai', 4, '2018-01-18 07:04:09', 0);
INSERT INTO `bot_logs` VALUES (108, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"name\":\"Tugas 1\",\"score\":\"100\",\"scored_time\":1516237037,\"url\":\"/api/v1/assignment/1\"}],\"intent\":\"grade\"}', 4, '2018-01-18 07:04:09', 1);
INSERT INTO `bot_logs` VALUES (109, 'nilai metode numerik', 4, '2018-01-24 04:15:52', 0);
INSERT INTO `bot_logs` VALUES (110, '{\"entity\":[],\"intent\":\"grade\"}', 4, '2018-01-24 04:15:52', 1);
INSERT INTO `bot_logs` VALUES (111, 'tugas ane', 4, '2018-01-24 04:16:00', 0);
INSERT INTO `bot_logs` VALUES (112, '{\"entity\":[{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1516236834,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"},{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat rancangan 2\",\"due_date\":1516409228,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/2\"}],\"intent\":\"assignment\"}', 4, '2018-01-24 04:16:00', 1);
INSERT INTO `bot_logs` VALUES (113, 'asisten yang ngajar saya kemarin', 4, '2018-01-24 04:16:16', 0);
INSERT INTO `bot_logs` VALUES (114, '{\"entity\":[],\"intent\":\"assistant\"}', 4, '2018-01-24 04:16:16', 1);
INSERT INTO `bot_logs` VALUES (115, 'asisten yang harusnya ngajar saya hari ini', 4, '2018-01-24 04:16:27', 0);
INSERT INTO `bot_logs` VALUES (116, '{\"entity\":[{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199751435760215.217800.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199927657053139.403577.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199124102442230.279765.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199666274331738.434554.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"}],\"intent\":\"assistant\"}', 4, '2018-01-24 04:16:27', 1);
INSERT INTO `bot_logs` VALUES (117, 'asdfjhasdkfjhaksdf', 4, '2018-01-24 04:20:18', 0);
INSERT INTO `bot_logs` VALUES (118, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"day\":\"Wednesday\",\"place\":\"UDJT-0206\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Metode Numerik\",\"day\":\"Monday\",\"place\":\"UDJT-0207\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Struktur Data\",\"day\":\"Thursday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\"}', 4, '2018-01-24 04:20:18', 1);
INSERT INTO `bot_logs` VALUES (119, 'halo halo asdfjkhsakdjhsdf', 4, '2018-01-24 04:20:28', 0);
INSERT INTO `bot_logs` VALUES (120, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"day\":\"Wednesday\",\"place\":\"UDJT-0206\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Metode Numerik\",\"day\":\"Monday\",\"place\":\"UDJT-0207\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Struktur Data\",\"day\":\"Thursday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\"}', 4, '2018-01-24 04:20:28', 1);
INSERT INTO `bot_logs` VALUES (121, 'bro bro bro kenapa gini?', 4, '2018-01-24 04:20:40', 0);
INSERT INTO `bot_logs` VALUES (122, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"day\":\"Wednesday\",\"place\":\"UDJT-0206\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Metode Numerik\",\"day\":\"Monday\",\"place\":\"UDJT-0207\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Struktur Data\",\"day\":\"Thursday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\"}', 4, '2018-01-24 04:20:40', 1);
INSERT INTO `bot_logs` VALUES (123, 'Asisten', 4, '2018-02-01 09:50:56', 0);
INSERT INTO `bot_logs` VALUES (124, '{\"entity\":[{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199124102442230.279765.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199666274331738.434554.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199556090164418.385200.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199848912061665.429961.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199751435760215.217800.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199927657053139.403577.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"}],\"intent\":\"assistant\"}', 4, '2018-02-01 09:50:56', 1);
INSERT INTO `bot_logs` VALUES (125, 'Bro kasih tau dong yang ngajar gue hari ini siapa?', 4, '2018-02-01 23:14:15', 0);
INSERT INTO `bot_logs` VALUES (126, '{\"entity\":[{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199927657053139.403577.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199556090164418.385200.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199848912061665.429961.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\"}', 4, '2018-02-01 23:14:15', 1);
INSERT INTO `bot_logs` VALUES (127, 'Kalau semua yang ngajar saya siapa?', 4, '2018-02-01 23:14:46', 0);
INSERT INTO `bot_logs` VALUES (128, '{\"entity\":[{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199751435760215.217800.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199927657053139.403577.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199124102442230.279765.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199666274331738.434554.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199556090164418.385200.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199848912061665.429961.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\"}', 4, '2018-02-01 23:14:46', 1);
INSERT INTO `bot_logs` VALUES (129, 'Hari ini ada kuliah?', 4, '2018-02-01 23:15:09', 0);
INSERT INTO `bot_logs` VALUES (130, '{\"entity\":[{\"course_name\":\"Metode Numerik\",\"day\":\"Friday\",\"place\":\"UDJT-0207\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\"}', 4, '2018-02-01 23:15:09', 1);
INSERT INTO `bot_logs` VALUES (131, 'Info 2 bulan lalu', 4, '2018-02-01 23:16:46', 0);
INSERT INTO `bot_logs` VALUES (132, '{\"entity\":[],\"intent\":\"information\"}', 4, '2018-02-01 23:16:46', 1);
INSERT INTO `bot_logs` VALUES (133, 'Semua berita?', 4, '2018-02-01 23:16:58', 0);
INSERT INTO `bot_logs` VALUES (134, '{\"entity\":[{\"description\":\"Ini adalah desc informasi Informasi 17 Januari 2018\",\"id\":2,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197473,\"title\":\"Informasi Informasi 18 Januari 2018\"},{\"description\":\"Ini adalah deskripsi informasi tanggal Informasi 17 Januari 2018\",\"id\":1,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197440,\"title\":\"Informasi 17 Januari 2018\"}],\"intent\":\"information\"}', 4, '2018-02-01 23:16:58', 1);
INSERT INTO `bot_logs` VALUES (135, 'Informasi 1 bulan lalu?', 4, '2018-02-01 23:17:17', 0);
INSERT INTO `bot_logs` VALUES (136, '{\"entity\":[{\"description\":\"Ini adalah desc informasi Informasi 17 Januari 2018\",\"id\":2,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197473,\"title\":\"Informasi Informasi 18 Januari 2018\"},{\"description\":\"Ini adalah deskripsi informasi tanggal Informasi 17 Januari 2018\",\"id\":1,\"image\":\"/api/v1/file/default/information.png\",\"posted_at\":1516197440,\"title\":\"Informasi 17 Januari 2018\"}],\"intent\":\"information\"}', 4, '2018-02-01 23:17:17', 1);
INSERT INTO `bot_logs` VALUES (137, 'Aduh ada tugas ga?', 4, '2018-02-01 23:17:45', 0);
INSERT INTO `bot_logs` VALUES (138, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat rancangan 2\",\"due_date\":1517532428,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/2\"},{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1518915234,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"}],\"intent\":\"assignment\"}', 4, '2018-02-01 23:17:45', 1);
INSERT INTO `bot_logs` VALUES (139, 'Bot.. tugas metode numerik', 4, '2018-02-01 23:18:07', 0);
INSERT INTO `bot_logs` VALUES (140, '{\"entity\":[{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1518915234,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"}],\"intent\":\"assignment\"}', 4, '2018-02-01 23:18:07', 1);
INSERT INTO `bot_logs` VALUES (141, 'Ada pr ga?', 4, '2018-02-01 23:18:17', 0);
INSERT INTO `bot_logs` VALUES (142, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat rancangan 2\",\"due_date\":1517532428,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/2\"},{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1518915234,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"}],\"intent\":\"assignment\"}', 4, '2018-02-01 23:18:17', 1);
INSERT INTO `bot_logs` VALUES (143, 'Bro bro bro... Nilai ane berape?', 4, '2018-02-01 23:18:37', 0);
INSERT INTO `bot_logs` VALUES (144, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"name\":\"Tugas 1\",\"score\":\"100\",\"scored_time\":1516237037,\"url\":\"/api/v1/assignment/1\"}],\"intent\":\"grade\"}', 4, '2018-02-01 23:18:37', 1);
INSERT INTO `bot_logs` VALUES (145, 'bro kasih tau dong asisten yang ngajar saya', 4, '2018-02-02 02:07:02', 0);
INSERT INTO `bot_logs` VALUES (146, '{\"entity\":[{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199927657053139.403577.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199124102442230.279765.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199666274331738.434554.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199556090164418.385200.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199848912061665.429961.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199751435760215.217800.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"}],\"intent\":\"assistant\"}', 4, '2018-02-02 02:07:02', 1);
INSERT INTO `bot_logs` VALUES (147, 'pengajar metode numerik', 4, '2018-02-02 02:07:48', 0);
INSERT INTO `bot_logs` VALUES (148, '{\"entity\":[{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199848912061665.429961.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199927657053139.403577.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199556090164418.385200.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"}],\"intent\":\"assistant\"}', 4, '2018-02-02 02:07:48', 1);
INSERT INTO `bot_logs` VALUES (149, 'hahaha kasih hahaha siapa kamu', 4, '2018-02-02 02:08:09', 0);
INSERT INTO `bot_logs` VALUES (150, '{\"entity\":[{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199751435760215.217800.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199927657053139.403577.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1516199124102442230.279765.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199666274331738.434554.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1516199556090164418.385200.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1516199848912061665.429961.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\"}', 4, '2018-02-02 02:08:09', 1);
INSERT INTO `bot_logs` VALUES (151, 'asdjfkhaskjdfhaskjdfhasdf', 4, '2018-02-02 02:08:54', 0);
INSERT INTO `bot_logs` VALUES (152, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"day\":\"Wednesday\",\"place\":\"UDJT-0206\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Metode Numerik\",\"day\":\"Friday\",\"place\":\"UDJT-0207\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Struktur Data\",\"day\":\"Monday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\"}', 4, '2018-02-02 02:08:54', 1);
INSERT INTO `bot_logs` VALUES (153, 'xxxxxx', 4, '2018-02-02 02:08:57', 0);
INSERT INTO `bot_logs` VALUES (154, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"day\":\"Wednesday\",\"place\":\"UDJT-0206\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Metode Numerik\",\"day\":\"Friday\",\"place\":\"UDJT-0207\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Struktur Data\",\"day\":\"Monday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\"}', 4, '2018-02-02 02:08:57', 1);
INSERT INTO `bot_logs` VALUES (155, 'tunjukin tugas', 4, '2018-02-02 02:47:32', 0);
INSERT INTO `bot_logs` VALUES (156, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat rancangan 2\",\"due_date\":1517532428,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/2\"},{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1518915234,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"}],\"intent\":\"assignment\"}', 4, '2018-02-02 02:47:32', 1);
INSERT INTO `bot_logs` VALUES (157, 'Lihat tugas', 4, '2018-02-10 11:03:18', 0);
INSERT INTO `bot_logs` VALUES (158, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat rancangan 2\",\"due_date\":1517536028,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/2\"},{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1518918834,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"}],\"intent\":\"assignment\",\"text\":\"Pasti ini yang kamu maksud\"}', 4, '2018-02-10 11:03:18', 1);
INSERT INTO `bot_logs` VALUES (159, 'Praktikum besok', 4, '2018-02-10 11:03:41', 0);
INSERT INTO `bot_logs` VALUES (160, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Alhamdulilah tidak ada praktikum. yuk perbanyak berdzikir\"}', 4, '2018-02-10 11:03:41', 1);
INSERT INTO `bot_logs` VALUES (161, 'Praktikum besok', 4, '2018-02-10 11:05:04', 0);
INSERT INTO `bot_logs` VALUES (162, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Tidak ada praktikum nih, kamu bisa berlibur sejenak, jangan lupa kerjain tugas-tugas yaa...\"}', 4, '2018-02-10 11:05:04', 1);
INSERT INTO `bot_logs` VALUES (163, 'Lihat nilai', 4, '2018-02-10 11:05:14', 0);
INSERT INTO `bot_logs` VALUES (164, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"name\":\"Tugas 1\",\"score\":\"100\",\"scored_time\":1516240637,\"url\":\"/api/v1/assignment/1\"}],\"intent\":\"grade\",\"text\":\"Cus!\"}', 4, '2018-02-10 11:05:14', 1);
INSERT INTO `bot_logs` VALUES (165, 'Oioi', 4, '2018-02-10 11:05:25', 0);
INSERT INTO `bot_logs` VALUES (166, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Ask me anything\"}', 4, '2018-02-10 11:05:25', 1);
INSERT INTO `bot_logs` VALUES (167, 'Oioi', 4, '2018-02-10 11:05:31', 0);
INSERT INTO `bot_logs` VALUES (168, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Whatsup bro? What can i help?\"}', 4, '2018-02-10 11:05:31', 1);
INSERT INTO `bot_logs` VALUES (169, 'Tugas yang belum dikerjakan', 4, '2018-02-10 11:05:34', 0);
INSERT INTO `bot_logs` VALUES (170, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat rancangan 2\",\"due_date\":1517536028,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/2\"},{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1518918834,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"}],\"intent\":\"assignment\",\"text\":\"Aha!\"}', 4, '2018-02-10 11:05:34', 1);
INSERT INTO `bot_logs` VALUES (171, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat makalah sederhana\",\"due_date\":1545134400,\"name\":\"Pemrograman Web\",\"url\":\"/api/v1/assignment/5\"}],\"intent\":\"assignment\",\"text\":\"Ini tugas baru\"}', 4, '2018-02-10 11:08:39', 1);
INSERT INTO `bot_logs` VALUES (172, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat makalah sederhana\",\"due_date\":1545134400,\"name\":\"Pemrograman Web\",\"url\":\"/api/v1/assignment/6\"}],\"intent\":\"assignment\",\"text\":\"Ini tugas baru\"}', 4, '2018-02-10 11:08:59', 1);
INSERT INTO `bot_logs` VALUES (173, 'Asisten', 4, '2018-02-10 11:19:02', 0);
INSERT INTO `bot_logs` VALUES (174, '{\"entity\":[{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236147055364000.417315.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236182357116000.275531.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236205084959000.765766.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236227727802000.031189.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\",\"text\":\"Ini dia!\"}', 4, '2018-02-10 11:19:02', 1);
INSERT INTO `bot_logs` VALUES (175, 'Lihat nilai', 4, '2018-02-10 11:19:35', 0);
INSERT INTO `bot_logs` VALUES (176, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"name\":\"Tugas 1\",\"score\":\"100\",\"scored_time\":1516240637,\"url\":\"/api/v1/assignment/1\"}],\"intent\":\"grade\",\"text\":\"Ini dia bro!\"}', 4, '2018-02-10 11:19:36', 1);
INSERT INTO `bot_logs` VALUES (177, 'Saya siapa', 4, '2018-02-10 11:19:44', 0);
INSERT INTO `bot_logs` VALUES (178, '{\"entity\":[{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236147055364000.417315.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236182357116000.275531.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236205084959000.765766.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236227727802000.031189.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\",\"text\":\"Semoga ini yang kamu maksud\"}', 4, '2018-02-10 11:19:44', 1);
INSERT INTO `bot_logs` VALUES (179, 'Aku siapa?', 4, '2018-02-10 11:19:54', 0);
INSERT INTO `bot_logs` VALUES (180, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Emang gue pikirin :p\"}', 4, '2018-02-10 11:19:54', 1);
INSERT INTO `bot_logs` VALUES (181, 'Kamu cantik', 4, '2018-02-10 11:20:04', 0);
INSERT INTO `bot_logs` VALUES (182, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Emang gue pikirin :p\"}', 4, '2018-02-10 11:20:04', 1);
INSERT INTO `bot_logs` VALUES (183, 'Kamu siapa?', 4, '2018-02-10 11:20:16', 0);
INSERT INTO `bot_logs` VALUES (184, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Wkwkwkw\"}', 4, '2018-02-10 11:20:16', 1);
INSERT INTO `bot_logs` VALUES (185, 'Eh', 4, '2018-02-10 11:20:20', 0);
INSERT INTO `bot_logs` VALUES (186, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat rancangan 2\",\"due_date\":1517536028,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/2\"},{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1518918834,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"},{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat makalah sederhana\",\"due_date\":1545109200,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/5\"},{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat makalah sederhana\",\"due_date\":1545109200,\"name\":\"Tugas 3\",\"url\":\"/api/v1/assignment/6\"}],\"intent\":\"assignment\",\"text\":\"Ieu wa!\"}', 4, '2018-02-10 11:20:20', 1);
INSERT INTO `bot_logs` VALUES (187, 'Hai', 4, '2018-02-10 11:20:26', 0);
INSERT INTO `bot_logs` VALUES (188, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hello! Risal Falah\"}', 4, '2018-02-10 11:20:26', 1);
INSERT INTO `bot_logs` VALUES (189, 'Hello', 4, '2018-02-10 11:20:31', 0);
INSERT INTO `bot_logs` VALUES (190, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Ask me anything\"}', 4, '2018-02-10 11:20:31', 1);
INSERT INTO `bot_logs` VALUES (191, 'Keren', 4, '2018-02-10 11:20:37', 0);
INSERT INTO `bot_logs` VALUES (192, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Heyhoo..\"}', 4, '2018-02-10 11:20:37', 1);
INSERT INTO `bot_logs` VALUES (193, 'Kamu kapan?', 4, '2018-02-10 11:20:52', 0);
INSERT INTO `bot_logs` VALUES (194, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Wkwkwkw\"}', 4, '2018-02-10 11:20:52', 1);
INSERT INTO `bot_logs` VALUES (195, 'Jadwal hari ini?', 4, '2018-02-10 11:21:00', 0);
INSERT INTO `bot_logs` VALUES (196, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"yes, tidak ada praktikum. kamu boleh bersenang-senang\"}', 4, '2018-02-10 11:21:00', 1);
INSERT INTO `bot_logs` VALUES (197, 'Asisten', 4, '2018-02-10 11:21:11', 0);
INSERT INTO `bot_logs` VALUES (198, '{\"entity\":[{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236227727802000.031189.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236147055364000.417315.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236182357116000.275531.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236205084959000.765766.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"}],\"intent\":\"assistant\",\"text\":\"Easy..\"}', 4, '2018-02-10 11:21:11', 1);
INSERT INTO `bot_logs` VALUES (199, 'Halo', 4, '2018-02-10 11:21:24', 0);
INSERT INTO `bot_logs` VALUES (200, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hello\"}', 4, '2018-02-10 11:21:24', 1);
INSERT INTO `bot_logs` VALUES (201, 'Hei', 4, '2018-02-10 11:21:27', 0);
INSERT INTO `bot_logs` VALUES (202, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hei\"}', 4, '2018-02-10 11:21:27', 1);
INSERT INTO `bot_logs` VALUES (203, 'Kamu siapa?', 4, '2018-02-10 11:21:33', 0);
INSERT INTO `bot_logs` VALUES (204, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hehehe\"}', 4, '2018-02-10 11:21:33', 1);
INSERT INTO `bot_logs` VALUES (205, 'Halow', 4, '2018-02-10 11:21:49', 0);
INSERT INTO `bot_logs` VALUES (206, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hello Risal Falah\"}', 4, '2018-02-10 11:21:49', 1);
INSERT INTO `bot_logs` VALUES (207, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"\",\"due_date\":1519819920,\"name\":\"Pemrograman Web\",\"url\":\"/api/v1/assignment/7\"}],\"intent\":\"assignment\",\"text\":\"Ini tugas baru\"}', 4, '2018-02-10 11:43:06', 1);
INSERT INTO `bot_logs` VALUES (208, 'Hello cantik..', 4, '2018-02-10 11:47:44', 0);
INSERT INTO `bot_logs` VALUES (209, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Oioi, ada yang bisa bot bantu?\"}', 4, '2018-02-10 11:47:44', 1);
INSERT INTO `bot_logs` VALUES (210, 'Praktikum hari ini', 4, '2018-02-10 11:47:59', 0);
INSERT INTO `bot_logs` VALUES (211, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Yes, tidak ada praktikum, jangan lupa kerjain tugas-tugas yaa...\"}', 4, '2018-02-10 11:47:59', 1);
INSERT INTO `bot_logs` VALUES (212, 'Hai cantik. :)', 4, '2018-02-10 11:48:30', 0);
INSERT INTO `bot_logs` VALUES (213, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hello\"}', 4, '2018-02-10 11:48:30', 1);
INSERT INTO `bot_logs` VALUES (214, 'Hari ini ada kuliah ga?', 4, '2018-02-10 11:48:50', 0);
INSERT INTO `bot_logs` VALUES (215, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Kamu bisa berlibur, tidak ada praktikum\"}', 4, '2018-02-10 11:48:50', 1);
INSERT INTO `bot_logs` VALUES (216, 'Kalau tugas ada ga sih?', 4, '2018-02-10 11:49:27', 0);
INSERT INTO `bot_logs` VALUES (217, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat rancangan 2\",\"due_date\":1517536028,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/2\"},{\"course_name\":\"Metode Numerik\",\"description\":\"Kerjakan soal interpolasi terlampir\",\"due_date\":1518918834,\"name\":\"Tugas 1\",\"url\":\"/api/v1/assignment/3\"},{\"course_name\":\"Pemrograman Web\",\"description\":\"-\",\"due_date\":1519794720,\"name\":\"Midtest 1 Take Home\",\"url\":\"/api/v1/assignment/7\"},{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat makalah sederhana\",\"due_date\":1545109200,\"name\":\"Tugas 2\",\"url\":\"/api/v1/assignment/5\"},{\"course_name\":\"Pemrograman Web\",\"description\":\"Membuat makalah sederhana\",\"due_date\":1545109200,\"name\":\"Tugas 3\",\"url\":\"/api/v1/assignment/6\"}],\"intent\":\"assignment\",\"text\":\"Ini dia yang kamu maksud\"}', 4, '2018-02-10 11:49:27', 1);
INSERT INTO `bot_logs` VALUES (218, 'Kalau berita terbaru ada apa aja?', 4, '2018-02-10 11:50:36', 0);
INSERT INTO `bot_logs` VALUES (219, '{\"entity\":[{\"description\":\"Ini adalah desc informasi Informasi 17 Januari 2018\",\"posted_at\":1516201073,\"title\":\"Informasi Informasi 18 Januari 2018\"},{\"description\":\"Ini adalah deskripsi informasi tanggal Informasi 17 Januari 2018\",\"posted_at\":1516201040,\"title\":\"Informasi 17 Januari 2018\"}],\"intent\":\"information\",\"text\":\"Silahkan ini dia\"}', 4, '2018-02-10 11:50:36', 1);
INSERT INTO `bot_logs` VALUES (220, 'Berarti berita hari ini ga ada ya?', 4, '2018-02-10 11:51:06', 0);
INSERT INTO `bot_logs` VALUES (221, '{\"entity\":[],\"intent\":\"information\",\"text\":\"\"}', 4, '2018-02-10 11:51:06', 1);
INSERT INTO `bot_logs` VALUES (222, 'Berarti berita hari ini ga ada ya?', 4, '2018-02-10 11:54:58', 0);
INSERT INTO `bot_logs` VALUES (223, '{\"entity\":[{\"description\":\"Ada teman kita nih yang lagi lomba arkavidia. Nama timnya meiko\",\"posted_at\":1518238468,\"title\":\"Informasi 10 Februari 2018\"}],\"intent\":\"information\",\"text\":\"Ini dia yang kamu maksud\"}', 4, '2018-02-10 11:54:58', 1);
INSERT INTO `bot_logs` VALUES (224, 'Terima kasih bot. Btw kamu kapan nikah?', 4, '2018-02-10 11:55:59', 0);
INSERT INTO `bot_logs` VALUES (225, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Jangan bercanda deh..\"}', 4, '2018-02-10 11:55:59', 1);
INSERT INTO `bot_logs` VALUES (226, 'Serius nih. Kamu kapan nikah', 4, '2018-02-10 11:56:20', 0);
INSERT INTO `bot_logs` VALUES (227, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Ckckckck\"}', 4, '2018-02-10 11:56:20', 1);
INSERT INTO `bot_logs` VALUES (228, 'Oh iya asisten yg ngajar aku lusa siapa?', 4, '2018-02-10 11:57:02', 0);
INSERT INTO `bot_logs` VALUES (229, '{\"entity\":[{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\",\"text\":\"Sundul Gan\"}', 4, '2018-02-10 11:57:02', 1);
INSERT INTO `bot_logs` VALUES (230, 'Jadwal lusa?', 4, '2018-02-10 11:57:13', 0);
INSERT INTO `bot_logs` VALUES (231, '{\"entity\":[{\"course_name\":\"Struktur Data\",\"day\":\"Monday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\",\"text\":\"Sundul Gan\"}', 4, '2018-02-10 11:57:13', 1);
INSERT INTO `bot_logs` VALUES (232, '{\"entity\":[{\"description\":\"Hai teman teman. Mohon doanya ya. Teman kita sedang mengikuti lomba\",\"posted_at\":1518239476,\"title\":\"Mohon doanya ya\"}],\"intent\":\"Information\",\"text\":\"Informasi baru\"}', 4, '2018-02-10 12:11:17', 1);
INSERT INTO `bot_logs` VALUES (233, '{\"entity\":[{\"description\":\"123\",\"posted_at\":1518239648,\"title\":\"123\"}],\"intent\":\"information\",\"text\":\"Informasi baru\"}', 4, '2018-02-10 12:14:08', 1);
INSERT INTO `bot_logs` VALUES (234, '{\"entity\":[{\"description\":\"abcdef\",\"posted_at\":1518239724,\"title\":\"abc\"}],\"intent\":\"information\",\"text\":\"Informasi baru\"}', 4, '2018-02-10 12:15:24', 1);
INSERT INTO `bot_logs` VALUES (235, 'Praktikum hari ini', 4, '2018-02-10 13:25:10', 0);
INSERT INTO `bot_logs` VALUES (236, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"yes, tidak ada praktikum. kamu boleh bersenang-senang\"}', 4, '2018-02-10 13:25:10', 1);
INSERT INTO `bot_logs` VALUES (237, 'Oioi', 4, '2018-02-10 13:25:18', 0);
INSERT INTO `bot_logs` VALUES (238, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hai Hai\"}', 4, '2018-02-10 13:25:18', 1);
INSERT INTO `bot_logs` VALUES (239, 'Lihat nilai', 4, '2018-02-10 13:42:29', 0);
INSERT INTO `bot_logs` VALUES (240, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"name\":\"Tugas 1\",\"score\":\"100\",\"scored_time\":1516240637,\"url\":\"/api/v1/assignment/1\"}],\"intent\":\"grade\",\"text\":\"Silahkan ini dia\"}', 4, '2018-02-10 13:42:29', 1);
INSERT INTO `bot_logs` VALUES (241, 'Praktikum hari ini', 4, '2018-02-10 14:02:35', 0);
INSERT INTO `bot_logs` VALUES (242, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Alhamdulilah tidak ada praktikum. yuk perbanyak berdzikir\"}', 4, '2018-02-10 14:02:35', 1);
INSERT INTO `bot_logs` VALUES (243, 'Hello ganteng', 4, '2018-02-10 14:02:57', 0);
INSERT INTO `bot_logs` VALUES (244, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Wah kamu bercanda deh...\"}', 4, '2018-02-10 14:02:57', 1);
INSERT INTO `bot_logs` VALUES (245, 'Hello cantik', 4, '2018-02-10 14:03:03', 0);
INSERT INTO `bot_logs` VALUES (246, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"oioi Risal Falah\"}', 4, '2018-02-10 14:03:03', 1);
INSERT INTO `bot_logs` VALUES (247, 'Mau nanya dong. Besok ada praktikum ga?', 4, '2018-02-10 14:03:26', 0);
INSERT INTO `bot_logs` VALUES (248, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Jawab gak yaa...\"}', 4, '2018-02-10 14:03:26', 1);
INSERT INTO `bot_logs` VALUES (249, 'Kalau hari ini ada praktikum ga sih?', 4, '2018-02-10 14:03:40', 0);
INSERT INTO `bot_logs` VALUES (250, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Kamu bisa berlibur, tidak ada praktikum\"}', 4, '2018-02-10 14:03:40', 1);
INSERT INTO `bot_logs` VALUES (251, 'Informasi yang lagi hot', 4, '2018-02-10 14:03:55', 0);
INSERT INTO `bot_logs` VALUES (252, '{\"entity\":[{\"description\":\"Ada teman kita nih yang lagi lomba arkavidia. Nama timnya meiko\",\"posted_at\":1518238468,\"title\":\"Informasi 10 Februari 2018\"},{\"description\":\"Ini adalah desc informasi Informasi 17 Januari 2018\",\"posted_at\":1516201073,\"title\":\"Informasi Informasi 18 Januari 2018\"},{\"description\":\"Ini adalah deskripsi informasi tanggal Informasi 17 Januari 2018\",\"posted_at\":1516201040,\"title\":\"Informasi 17 Januari 2018\"}],\"intent\":\"information\",\"text\":\"Ieu wa!\"}', 4, '2018-02-10 14:03:55', 1);
INSERT INTO `bot_logs` VALUES (253, 'Jadwal untuk lusa apa aja ya?', 4, '2018-02-10 14:04:19', 0);
INSERT INTO `bot_logs` VALUES (254, '{\"entity\":[{\"course_name\":\"Struktur Data\",\"day\":\"Monday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\",\"text\":\"Ini nih jawabannya\"}', 4, '2018-02-10 14:04:19', 1);
INSERT INTO `bot_logs` VALUES (255, 'Kalau yang yang ngajar nanti lusa?', 4, '2018-02-10 14:04:30', 0);
INSERT INTO `bot_logs` VALUES (256, '{\"entity\":[{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\",\"text\":\"Sikat gan!\"}', 4, '2018-02-10 14:04:30', 1);
INSERT INTO `bot_logs` VALUES (257, 'Wah terima kasih... Btw kapan kamu nikah?', 4, '2018-02-10 14:05:02', 0);
INSERT INTO `bot_logs` VALUES (258, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Ckckckck\"}', 4, '2018-02-10 14:05:02', 1);
INSERT INTO `bot_logs` VALUES (259, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"\",\"due_date\":1519300800,\"name\":\"Pemrograman Web\",\"url\":\"/api/v1/assignment/8\"}],\"intent\":\"assignment\",\"text\":\"Ini tugas baru\"}', 4, '2018-02-10 14:05:34', 1);
INSERT INTO `bot_logs` VALUES (260, '{\"entity\":[{\"description\":\"Hai teman teman. Mohon doanya ya. Teman kita sedang mengikuti lomba hackathon di Arkavidia ITB. Doakan semoga bisa menang. Terima Kasih\",\"posted_at\":1518246360,\"title\":\"Mohon doanya ya\"}],\"intent\":\"information\",\"text\":\"Informasi baru\"}', 4, '2018-02-10 14:06:00', 1);
INSERT INTO `bot_logs` VALUES (261, 'Praktikum hari ini', 4, '2018-02-10 14:18:09', 0);
INSERT INTO `bot_logs` VALUES (262, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Yes, tidak ada praktikum, selamat berlibur, jangan lupa kerjain tugas yang belum ya\"}', 4, '2018-02-10 14:18:09', 1);
INSERT INTO `bot_logs` VALUES (263, 'Kalau misalnya berita hari ini ada ga?', 4, '2018-02-10 14:18:22', 0);
INSERT INTO `bot_logs` VALUES (264, '{\"entity\":[{\"description\":\"Hai teman teman. Mohon doanya ya. Teman kita sedang mengikuti lomba hackathon di Arkavidia ITB. Doakan semoga bisa menang. Terima Kasih\",\"posted_at\":1518246360,\"title\":\"Mohon doanya ya\"},{\"description\":\"Ada teman kita nih yang lagi lomba arkavidia. Nama timnya meiko\",\"posted_at\":1518238468,\"title\":\"Informasi 10 Februari 2018\"}],\"intent\":\"information\",\"text\":\"Sundul Gan\"}', 4, '2018-02-10 14:18:22', 1);
INSERT INTO `bot_logs` VALUES (265, 'Hai kalau lusa ada praktikum ga ya?', 4, '2018-02-10 14:18:50', 0);
INSERT INTO `bot_logs` VALUES (266, '{\"entity\":[{\"course_name\":\"Struktur Data\",\"day\":\"Monday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\",\"text\":\"Cus!\"}', 4, '2018-02-10 14:18:50', 1);
INSERT INTO `bot_logs` VALUES (267, 'Siapa sih yang ngajar praktikum lusa?', 4, '2018-02-10 14:19:00', 0);
INSERT INTO `bot_logs` VALUES (268, '{\"entity\":[{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\",\"text\":\"Ieu wa!\"}', 4, '2018-02-10 14:19:00', 1);
INSERT INTO `bot_logs` VALUES (269, 'Bro nilai gue berapa?', 4, '2018-02-10 14:19:24', 0);
INSERT INTO `bot_logs` VALUES (270, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"name\":\"Tugas 1\",\"score\":\"100\",\"scored_time\":1516240637,\"url\":\"/api/v1/assignment/1\"}],\"intent\":\"grade\",\"text\":\"Ini dia!\"}', 4, '2018-02-10 14:19:24', 1);
INSERT INTO `bot_logs` VALUES (271, 'Hello cantik', 4, '2018-02-10 14:19:42', 0);
INSERT INTO `bot_logs` VALUES (272, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hei\"}', 4, '2018-02-10 14:19:42', 1);
INSERT INTO `bot_logs` VALUES (273, 'Kamu udah nikah belum?', 4, '2018-02-10 14:19:51', 0);
INSERT INTO `bot_logs` VALUES (274, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Wah kamu bercanda deh...\"}', 4, '2018-02-10 14:19:51', 1);
INSERT INTO `bot_logs` VALUES (275, 'Siapa sih yang ngajar praktikum lusa?', 4, '2018-02-10 14:31:58', 0);
INSERT INTO `bot_logs` VALUES (276, '{\"entity\":[{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\",\"text\":\"Silahkan ini dia\"}', 4, '2018-02-10 14:31:58', 1);
INSERT INTO `bot_logs` VALUES (277, 'Kasih tahu dong semua asisten mengajar saya', 4, '2018-02-10 14:32:12', 0);
INSERT INTO `bot_logs` VALUES (278, '{\"entity\":[{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236227727802000.031189.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236147055364000.417315.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236182357116000.275531.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236205084959000.765766.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"}],\"intent\":\"assistant\",\"text\":\"Sikat gan!\"}', 4, '2018-02-10 14:32:12', 1);
INSERT INTO `bot_logs` VALUES (279, 'Tes', 4, '2018-02-10 14:37:31', 0);
INSERT INTO `bot_logs` VALUES (280, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hello my friend\"}', 4, '2018-02-10 14:37:31', 1);
INSERT INTO `bot_logs` VALUES (281, 'Hi', 4, '2018-02-10 14:37:58', 0);
INSERT INTO `bot_logs` VALUES (282, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hey\"}', 4, '2018-02-10 14:37:58', 1);
INSERT INTO `bot_logs` VALUES (283, 'Hey', 4, '2018-02-10 14:38:28', 0);
INSERT INTO `bot_logs` VALUES (284, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hallo, ada yang bisa bot bantu?\"}', 4, '2018-02-10 14:38:28', 1);
INSERT INTO `bot_logs` VALUES (285, 'Jadwal hari ini', 4, '2018-02-10 14:38:35', 0);
INSERT INTO `bot_logs` VALUES (286, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Hari ini tidak ada jadwal, yeeee.. jangan lupa bahagia\"}', 4, '2018-02-10 14:38:35', 1);
INSERT INTO `bot_logs` VALUES (287, 'Jadwal hari ini', 4, '2018-02-10 14:44:02', 0);
INSERT INTO `bot_logs` VALUES (288, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Kamu bisa berlibur, tidak ada praktikum\"}', 4, '2018-02-10 14:44:02', 1);
INSERT INTO `bot_logs` VALUES (289, 'Praktikum hari ini', 4, '2018-02-10 14:45:46', 0);
INSERT INTO `bot_logs` VALUES (290, '{\"entity\":[],\"intent\":\"schedule\",\"text\":\"Hari ini tidak ada jadwal, yeeee.. selamat berlibur :)\"}', 4, '2018-02-10 14:45:46', 1);
INSERT INTO `bot_logs` VALUES (291, 'Jadwal untuk apa aja ya?', 4, '2018-02-10 14:46:04', 0);
INSERT INTO `bot_logs` VALUES (292, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"day\":\"Wednesday\",\"place\":\"UDJT-0206\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Metode Numerik\",\"day\":\"Friday\",\"place\":\"UDJT-0207\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Struktur Data\",\"day\":\"Monday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\",\"text\":\"Here it is!\"}', 4, '2018-02-10 14:46:04', 1);
INSERT INTO `bot_logs` VALUES (293, 'Jadwal untuk lusa apa aja ya?', 4, '2018-02-10 14:46:14', 0);
INSERT INTO `bot_logs` VALUES (294, '{\"entity\":[{\"course_name\":\"Struktur Data\",\"day\":\"Monday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\",\"text\":\"Silahkan ini dia\"}', 4, '2018-02-10 14:46:14', 1);
INSERT INTO `bot_logs` VALUES (295, 'Bro siapa sih yang ngajar untuk lusa?', 4, '2018-02-10 14:46:46', 0);
INSERT INTO `bot_logs` VALUES (296, '{\"entity\":[{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\",\"text\":\"Sundul Gan\"}', 4, '2018-02-10 14:46:46', 1);
INSERT INTO `bot_logs` VALUES (297, 'Yang ngajar siapa aja ya', 4, '2018-02-10 14:46:59', 0);
INSERT INTO `bot_logs` VALUES (298, '{\"entity\":[{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236147055364000.417315.2.jpg\",\"line_id\":\"ghifarigue\",\"name\":\"Mohammad Ghifari\",\"phone\":\"81312312303\"},{\"courses\":[\"Pemrograman Web\",\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236182357116000.275531.2.jpg\",\"line_id\":\"ibepgue\",\"name\":\"Febryani Pertiwi Puteri\",\"phone\":\"81312312304\"},{\"courses\":[\"Pemrograman Web\"],\"image\":\"/api/v1/file/profile/1518236205084959000.765766.2.jpg\",\"line_id\":\"-\",\"name\":\"Asep Nur Muhammad\",\"phone\":\"81312312301\"},{\"courses\":[\"Pemrograman Web\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236273807452000.559803.2.jpg\",\"line_id\":\"-\",\"name\":\"Yunilucky Siswantari\",\"phone\":\"81312312302\"},{\"courses\":[\"Metode Numerik\"],\"image\":\"/api/v1/file/profile/1518236227727802000.031189.2.jpg\",\"line_id\":\"dayatura\",\"name\":\"Hidayaturrahman\",\"phone\":\"81312312305\"},{\"courses\":[\"Metode Numerik\",\"Struktur Data\"],\"image\":\"/api/v1/file/profile/1518236252208989000.596264.2.jpg\",\"line_id\":\"rifkirifkigue\",\"name\":\"Muhammad Rifki\",\"phone\":\"81312312306\"}],\"intent\":\"assistant\",\"text\":\"Pasti ini yang kamu maksud\"}', 4, '2018-02-10 14:46:59', 1);
INSERT INTO `bot_logs` VALUES (299, 'Titip absen ya?', 4, '2018-02-10 14:47:14', 0);
INSERT INTO `bot_logs` VALUES (300, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"day\":\"Wednesday\",\"place\":\"UDJT-0206\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Metode Numerik\",\"day\":\"Friday\",\"place\":\"UDJT-0207\",\"time\":\"10:00 - 13:20\"},{\"course_name\":\"Struktur Data\",\"day\":\"Monday\",\"place\":\"UDJT-305\",\"time\":\"10:00 - 13:20\"}],\"intent\":\"schedule\",\"text\":\"Easy..\"}', 4, '2018-02-10 14:47:14', 1);
INSERT INTO `bot_logs` VALUES (301, 'Hello cantik', 4, '2018-02-10 14:47:25', 0);
INSERT INTO `bot_logs` VALUES (302, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hello! Risal Falah\"}', 4, '2018-02-10 14:47:25', 1);
INSERT INTO `bot_logs` VALUES (303, 'Kapan kamu nikah sih cantik?', 4, '2018-02-10 14:47:36', 0);
INSERT INTO `bot_logs` VALUES (304, '{\"entity\":[],\"intent\":\"messageonly\",\"text\":\"Hehehe\"}', 4, '2018-02-10 14:47:36', 1);
INSERT INTO `bot_logs` VALUES (305, '{\"entity\":[{\"course_name\":\"Pemrograman Web\",\"description\":\"\",\"due_date\":1519128000,\"name\":\"Pemrograman Web\",\"url\":\"/api/v1/assignment/9\"}],\"intent\":\"assignment\",\"text\":\"Ini tugas baru\"}', 4, '2018-02-10 14:47:56', 1);
COMMIT;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of courses
-- ----------------------------
BEGIN;
INSERT INTO `courses` VALUES ('1', 'Bahasa Indonesia I', 'Bahasa Indonesia I merupakan mata kuliah bahasa indonesia dasar, yang meliputi konsep dasar tata bahasa dan Ejaan yang di sempurnakan (EYD)', 3, '2018-01-17 19:38:35', '2018-01-17 19:38:36');
INSERT INTO `courses` VALUES ('2', 'Matematika Dasar', 'Matematika Dasar merupakan mata kuliah matematika dasar yang mempelajari tentang konsep matematika dasar I', 3, '2018-01-17 19:38:48', '2018-01-17 14:46:46');
INSERT INTO `courses` VALUES ('3', 'Fisika Lanjutan', 'Fisika Lanjutan merupakan fisika yang membahasa tentang lanjutan dari fisika dasar, konsep tentang listrik statis, dinamis dan teori yang lainnya', 3, '2018-01-17 21:51:30', '2018-01-17 14:52:33');
COMMIT;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of files
-- ----------------------------
BEGIN;
INSERT INTO `files` VALUES ('1516199124102442230.058970.1', '23735012_174488926469402_6825093196763103232_', 'image/jpeg', 'jpg', 0, 3, 'PL-IMG-M', NULL, '3', '2018-01-17 14:25:24', '2018-02-10 11:16:45');
INSERT INTO `files` VALUES ('1516199124102442230.279765.2', '23735012_174488926469402_6825093196763103232_', 'image/jpeg', 'jpg', 0, 3, 'PL-IMG-T', NULL, '3', '2018-01-17 14:25:24', '2018-02-10 11:16:45');
INSERT INTO `files` VALUES ('1516199556090164418.344165.1', 'dayat', 'image/jpeg', 'jpg', 0, 6, 'PL-IMG-M', NULL, '6', '2018-01-17 14:32:36', '2018-02-10 11:17:07');
INSERT INTO `files` VALUES ('1516199556090164418.385200.2', 'dayat', 'image/jpeg', 'jpg', 0, 6, 'PL-IMG-T', NULL, '6', '2018-01-17 14:32:36', '2018-02-10 11:17:07');
INSERT INTO `files` VALUES ('1516199666274331738.434554.2', 'yuni', 'image/jpeg', 'jpg', 0, 5, 'PL-IMG-T', NULL, '5', '2018-01-17 14:34:26', '2018-02-10 11:17:53');
INSERT INTO `files` VALUES ('1516199666274331738.839623.1', 'yuni', 'image/jpeg', 'jpg', 0, 5, 'PL-IMG-M', NULL, '5', '2018-01-17 14:34:26', '2018-02-10 11:17:53');
INSERT INTO `files` VALUES ('1516199751435760215.217800.2', 'ghifari', 'image/jpeg', 'jpg', 0, 1, 'PL-IMG-T', NULL, '1', '2018-01-17 14:35:51', '2018-02-10 11:15:47');
INSERT INTO `files` VALUES ('1516199751435760215.315635.1', 'ghifari', 'image/jpeg', 'jpg', 0, 1, 'PL-IMG-M', NULL, '1', '2018-01-17 14:35:51', '2018-02-10 11:15:47');
INSERT INTO `files` VALUES ('1516199848912061665.360893.1', 'rifki', 'image/jpeg', 'jpg', 0, 7, 'PL-IMG-M', NULL, '7', '2018-01-17 14:37:28', '2018-02-10 11:17:32');
INSERT INTO `files` VALUES ('1516199848912061665.429961.2', 'rifki', 'image/jpeg', 'jpg', 0, 7, 'PL-IMG-T', NULL, '7', '2018-01-17 14:37:28', '2018-02-10 11:17:32');
INSERT INTO `files` VALUES ('1516199927657053139.307970.1', 'febryani', 'image/jpeg', 'jpg', 0, 2, 'PL-IMG-M', NULL, '2', '2018-01-17 14:38:47', '2018-02-10 11:16:22');
INSERT INTO `files` VALUES ('1516199927657053139.403577.2', 'febryani', 'image/jpeg', 'jpg', 0, 2, 'PL-IMG-T', NULL, '2', '2018-01-17 14:38:47', '2018-02-10 11:16:22');
INSERT INTO `files` VALUES ('1518236120263357000.321771.2', 'Screen Shot 2018-02-10 at 11.12.56 AM', 'image/png', 'png', 1, 4, 'PL-IMG-T', NULL, '4', '2018-02-10 11:15:20', '2018-02-10 11:15:20');
INSERT INTO `files` VALUES ('1518236120263357000.812694.1', 'Screen Shot 2018-02-10 at 11.12.56 AM', 'image/png', 'png', 1, 4, 'PL-IMG-M', NULL, '4', '2018-02-10 11:15:20', '2018-02-10 11:15:20');
INSERT INTO `files` VALUES ('1518236147055364000.243125.1', 'Screen Shot 2018-02-10 at 11.15.03 AM', 'image/png', 'png', 1, 1, 'PL-IMG-M', NULL, '1', '2018-02-10 11:15:47', '2018-02-10 11:15:47');
INSERT INTO `files` VALUES ('1518236147055364000.417315.2', 'Screen Shot 2018-02-10 at 11.15.03 AM', 'image/png', 'png', 1, 1, 'PL-IMG-T', NULL, '1', '2018-02-10 11:15:47', '2018-02-10 11:15:47');
INSERT INTO `files` VALUES ('1518236182357116000.275531.2', 'Screen Shot 2018-02-10 at 11.13.18 AM', 'image/png', 'png', 1, 2, 'PL-IMG-T', NULL, '2', '2018-02-10 11:16:22', '2018-02-10 11:16:22');
INSERT INTO `files` VALUES ('1518236182357116000.319768.1', 'Screen Shot 2018-02-10 at 11.13.18 AM', 'image/png', 'png', 1, 2, 'PL-IMG-M', NULL, '2', '2018-02-10 11:16:22', '2018-02-10 11:16:22');
INSERT INTO `files` VALUES ('1518236205084959000.765766.2', 'Screen Shot 2018-02-10 at 11.12.38 AM', 'image/png', 'png', 1, 3, 'PL-IMG-T', NULL, '3', '2018-02-10 11:16:45', '2018-02-10 11:16:45');
INSERT INTO `files` VALUES ('1518236205084959000.848208.1', 'Screen Shot 2018-02-10 at 11.12.38 AM', 'image/png', 'png', 1, 3, 'PL-IMG-M', NULL, '3', '2018-02-10 11:16:45', '2018-02-10 11:16:45');
INSERT INTO `files` VALUES ('1518236227727802000.031189.2', 'Screen Shot 2018-02-10 at 11.14.08 AM', 'image/png', 'png', 1, 6, 'PL-IMG-T', NULL, '6', '2018-02-10 11:17:07', '2018-02-10 11:17:07');
INSERT INTO `files` VALUES ('1518236227727802000.749117.1', 'Screen Shot 2018-02-10 at 11.14.08 AM', 'image/png', 'png', 1, 6, 'PL-IMG-M', NULL, '6', '2018-02-10 11:17:07', '2018-02-10 11:17:07');
INSERT INTO `files` VALUES ('1518236252208989000.340995.1', 'Screen Shot 2018-02-10 at 11.14.35 AM', 'image/png', 'png', 1, 7, 'PL-IMG-M', NULL, '7', '2018-02-10 11:17:32', '2018-02-10 11:17:32');
INSERT INTO `files` VALUES ('1518236252208989000.596264.2', 'Screen Shot 2018-02-10 at 11.14.35 AM', 'image/png', 'png', 1, 7, 'PL-IMG-T', NULL, '7', '2018-02-10 11:17:32', '2018-02-10 11:17:32');
INSERT INTO `files` VALUES ('1518236273807452000.559803.2', 'Screen Shot 2018-02-10 at 11.13.41 AM', 'image/png', 'png', 1, 5, 'PL-IMG-T', NULL, '5', '2018-02-10 11:17:53', '2018-02-10 11:17:53');
INSERT INTO `files` VALUES ('1518236273807452000.585875.1', 'Screen Shot 2018-02-10 at 11.13.41 AM', 'image/png', 'png', 1, 5, 'PL-IMG-M', NULL, '5', '2018-02-10 11:17:53', '2018-02-10 11:17:53');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of grade_parameters
-- ----------------------------
BEGIN;
INSERT INTO `grade_parameters` VALUES (1, 'ASSIGNMENT', 100.00, 0, 2, '2018-01-17 14:46:46', '2018-01-17 14:46:46');
INSERT INTO `grade_parameters` VALUES (2, 'ASSIGNMENT', 50.00, 0, 3, '2018-01-17 14:52:33', '2018-01-17 14:52:33');
INSERT INTO `grade_parameters` VALUES (3, 'QUIZ', 50.00, 0, 3, '2018-01-17 14:52:33', '2018-01-17 14:52:33');
INSERT INTO `grade_parameters` VALUES (4, 'ASSIGNMENT', 30.00, 0, 1, '2018-01-18 08:43:22', '2018-01-18 08:43:23');
INSERT INTO `grade_parameters` VALUES (5, 'MID', 35.00, 0, 1, '2018-01-18 08:43:43', '2018-01-18 08:43:45');
INSERT INTO `grade_parameters` VALUES (6, 'FINAL', 35.00, 0, 1, '2018-01-18 08:43:57', '2018-01-18 08:43:58');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of informations
-- ----------------------------
BEGIN;
INSERT INTO `informations` VALUES (1, 'Libur Tahun Baru', 'Kepada seluruh mahasiswa, sehubungan dengan penyambutan tahun baru maka pada tanggal 1 January 2019 Kegitan Belajar Mengajar diliburkan', NULL, '2018-12-30 21:57:20', '2008-12-30 08:11:28');
INSERT INTO `informations` VALUES (2, 'Meeting Dadakan', 'Sehubungan dengan adanya meeting pada pagi ini maka kuliah di liburkan', NULL, '2019-01-09 21:57:53', '2008-12-30 08:11:28');
INSERT INTO `informations` VALUES (3, 'Informasi Lomba', 'Ada teman kita nih yang lagi lomba arkavidia. Nama timnya meiko. Udah menang', NULL, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
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
  PRIMARY KEY (`id`) USING BTREE,
  KEY `fk_inventories_places` (`places_id`) USING BTREE,
  CONSTRAINT `fk_inventories_places` FOREIGN KEY (`places_id`) REFERENCES `places` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

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
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of meetings
-- ----------------------------
BEGIN;
INSERT INTO `meetings` VALUES (1, 1, 'meet-1', 'meet-1', '2018-11-05 01:04:57', 1, '2019-01-07 01:05:27', '2018-09-15 01:05:33');
INSERT INTO `meetings` VALUES (2, 2, 'meet-2', 'meet-2', '2019-01-12 01:05:59', 1, '2019-01-12 01:06:04', '2019-01-12 01:06:08');
INSERT INTO `meetings` VALUES (3, 3, 'meet-3', 'meet-3', '2018-11-05 01:04:57', 1, '2019-01-07 01:05:27', '2018-09-15 01:05:33');
INSERT INTO `meetings` VALUES (4, 4, 'meet-4', 'meet-4', '2019-01-12 01:05:59', 1, '2019-01-12 01:06:04', '2019-01-12 01:06:08');
INSERT INTO `meetings` VALUES (6, 6, 'meet-6', 'meet-6', '2019-01-12 01:05:59', 1, '2019-01-12 01:06:04', '2019-01-12 01:06:08');
INSERT INTO `meetings` VALUES (7, 7, 'meet-7', 'meet-7', '2018-11-05 01:04:57', 1, '2019-01-07 01:05:27', '2018-09-15 01:05:33');
INSERT INTO `meetings` VALUES (8, 1, 'meet-1', 'meet-1', '2018-11-05 01:04:57', 2, '2019-01-07 01:05:27', '2018-09-15 01:05:33');
INSERT INTO `meetings` VALUES (9, 2, 'meet-2', 'meet-2', '2019-01-12 01:05:59', 2, '2019-01-12 01:06:04', '2019-01-12 01:06:08');
INSERT INTO `meetings` VALUES (10, 3, 'meet-3', 'meet-3', '2018-11-05 01:04:57', 2, '2019-01-07 01:05:27', '2018-09-15 01:05:33');
INSERT INTO `meetings` VALUES (11, 4, 'meet-4', 'meet-4', '2019-01-12 01:05:59', 2, '2019-01-12 01:06:04', '2019-01-12 01:06:08');
INSERT INTO `meetings` VALUES (12, 6, 'meet-6', 'meet-6', '2019-01-12 01:05:59', 2, '2019-01-12 01:06:04', '2019-01-12 01:06:08');
INSERT INTO `meetings` VALUES (13, 7, 'meet-7', 'meet-7', '2018-11-05 01:04:57', 2, '2019-01-07 01:05:27', '2018-09-15 01:05:33');
INSERT INTO `meetings` VALUES (14, 1, 'meet-1', 'meet-1', '2018-11-05 01:04:57', 3, '2019-01-07 01:05:27', '2018-09-15 01:05:33');
INSERT INTO `meetings` VALUES (15, 2, 'meet-2', 'meet-2', '2019-01-12 01:05:59', 3, '2019-01-12 01:06:04', '2019-01-12 01:06:08');
INSERT INTO `meetings` VALUES (16, 3, 'meet-3', 'meet-3', '2018-11-05 01:04:57', 3, '2019-01-07 01:05:27', '2018-09-15 01:05:33');
INSERT INTO `meetings` VALUES (17, 4, 'meet-4', 'meet-4', '2019-01-12 01:05:59', 3, '2019-01-12 01:06:04', '2019-01-12 01:06:08');
INSERT INTO `meetings` VALUES (18, 6, 'meet-6', 'meet-6', '2019-01-12 01:05:59', 3, '2019-01-12 01:06:04', '2019-01-12 01:06:08');
INSERT INTO `meetings` VALUES (19, 7, 'meet-7', 'meet-7', '2018-11-05 01:04:57', 3, '2019-01-07 01:05:27', '2018-09-15 01:05:33');
COMMIT;

-- ----------------------------
-- Table structure for notifications
-- ----------------------------
DROP TABLE IF EXISTS `notifications`;
CREATE TABLE `notifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `users_id` int(11) unsigned NOT NULL,
  `onesignal_id` varchar(36) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `fk_notifications_users` (`users_id`,`onesignal_id`) USING BTREE,
  CONSTRAINT `fk_notifications_users` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of p_users_assignments
-- ----------------------------
BEGIN;
INSERT INTO `p_users_assignments` VALUES (1, 1, 80.00, 'Good', '2019-01-10 00:58:26', '2019-01-10 00:58:33');
INSERT INTO `p_users_assignments` VALUES (1, 2, 75.00, 'Enough', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (1, 3, 75.00, 'Enough', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (1, 4, 75.00, 'Enough', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (1, 5, 75.00, 'Enough', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (2, 1, 80.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (2, 2, 80.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (2, 3, 80.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (2, 4, 80.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (2, 5, 80.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (3, 1, 80.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (3, 2, 87.00, 'Good ', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (3, 3, 87.00, 'Good ', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (3, 4, 87.00, 'Good ', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (3, 5, 87.00, 'Good ', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (5, 1, 81.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (5, 2, 76.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (5, 3, 76.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (5, 4, 76.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (5, 5, 76.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (6, 1, 82.00, 'Good', '2019-01-12 01:00:17', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (6, 2, 89.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (6, 3, 89.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (6, 4, 89.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (6, 5, 89.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (7, 1, 89.00, 'Good', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (7, 2, 79.00, 'Enough', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (7, 3, 79.00, 'Enough', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (7, 4, 79.00, 'Enough', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
INSERT INTO `p_users_assignments` VALUES (7, 5, 79.00, 'Enough', '2019-01-10 00:58:52', '2019-01-10 00:58:52');
COMMIT;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of p_users_schedules
-- ----------------------------
BEGIN;
INSERT INTO `p_users_schedules` VALUES (1, 1, 1, '2018-11-01 23:59:25', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (1, 3, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (1, 4, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (2, 1, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (2, 2, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (2, 3, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (3, 1, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (3, 2, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (3, 3, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (4, 4, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (4, 5, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (4, 6, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (5, 4, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (5, 5, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (5, 6, 1, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (6, 1, 2, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (6, 2, 2, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (6, 4, 2, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (6, 5, 2, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (7, 3, 2, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
INSERT INTO `p_users_schedules` VALUES (7, 6, 2, '2019-01-10 11:54:28', '2019-01-10 11:54:28');
COMMIT;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of places
-- ----------------------------
BEGIN;
INSERT INTO `places` VALUES ('UDJT-0206', 'PPBS D', '2018-01-17 19:38:17', '2018-01-17 19:38:18');
INSERT INTO `places` VALUES ('UDJT-0207', 'PPBS D', '2018-01-17 14:46:46', '2018-01-17 14:46:46');
INSERT INTO `places` VALUES ('UDJT-305', 'PPBS A', '2018-01-17 14:52:33', '2018-01-17 14:52:33');
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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of rolegroups
-- ----------------------------
BEGIN;
INSERT INTO `rolegroups` VALUES (1, 'Assistant', '2018-01-17 19:37:05', '2018-01-17 19:37:07');
INSERT INTO `rolegroups` VALUES (2, 'Dosen', '2019-01-07 21:46:56', '2019-01-07 21:47:01');
INSERT INTO `rolegroups` VALUES (3, 'Assistant Dosen', '2019-01-07 21:48:08', '2019-01-07 21:48:13');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of rolegroups_modules
-- ----------------------------
BEGIN;
INSERT INTO `rolegroups_modules` VALUES (1, '', 'CREATE', '0000-00-00 00:00:00', '0000-00-00 00:00:00');
INSERT INTO `rolegroups_modules` VALUES (1, 'users', 'CREATE', '2018-01-17 21:14:38', '2018-01-17 21:14:40');
INSERT INTO `rolegroups_modules` VALUES (1, 'users', 'READ', '2018-01-17 21:14:47', '2018-01-17 21:14:48');
INSERT INTO `rolegroups_modules` VALUES (1, 'users', 'DELETE', '2018-01-17 21:15:14', '2018-01-17 21:15:15');
INSERT INTO `rolegroups_modules` VALUES (1, 'courses', 'CREATE', '2018-01-17 21:15:24', '2018-01-17 21:15:25');
INSERT INTO `rolegroups_modules` VALUES (1, 'courses', 'READ', '2018-01-17 21:15:32', '2018-01-17 21:15:34');
INSERT INTO `rolegroups_modules` VALUES (1, 'courses', 'UPDATE', '2018-01-17 21:15:41', '2018-01-17 21:15:48');
INSERT INTO `rolegroups_modules` VALUES (1, 'courses', 'DELETE', '2018-01-17 21:16:08', '2018-01-17 21:16:09');
INSERT INTO `rolegroups_modules` VALUES (1, 'attendances', 'CREATE', '2018-01-17 21:16:18', '2018-01-17 21:16:19');
INSERT INTO `rolegroups_modules` VALUES (1, 'attendances', 'READ', '2018-01-17 21:16:29', '2018-01-17 21:16:31');
INSERT INTO `rolegroups_modules` VALUES (1, 'attendances', 'UPDATE', '2018-01-17 21:16:43', '2018-01-17 21:16:44');
INSERT INTO `rolegroups_modules` VALUES (1, 'attendances', 'DELETE', '2018-01-17 21:16:51', '2018-01-17 21:16:52');
INSERT INTO `rolegroups_modules` VALUES (1, 'roles', 'CREATE', '2018-01-17 21:17:03', '2018-01-17 21:17:19');
INSERT INTO `rolegroups_modules` VALUES (1, 'roles', 'READ', '2018-01-17 21:17:17', '2018-01-17 21:17:18');
INSERT INTO `rolegroups_modules` VALUES (1, 'roles', 'UPDATE', '2018-01-17 21:17:26', '2018-01-17 21:17:27');
INSERT INTO `rolegroups_modules` VALUES (1, 'roles', 'DELETE', '2018-01-17 21:17:33', '2018-01-17 21:17:34');
INSERT INTO `rolegroups_modules` VALUES (1, 'schedules', 'CREATE', '2018-01-17 21:17:46', '2018-01-17 21:17:53');
INSERT INTO `rolegroups_modules` VALUES (1, 'schedules', 'READ', '2018-01-17 09:18:01', '2018-01-17 21:18:04');
INSERT INTO `rolegroups_modules` VALUES (1, 'schedules', 'UPDATE', '2018-01-17 21:18:17', '2018-01-17 21:18:18');
INSERT INTO `rolegroups_modules` VALUES (1, 'schedules', 'DELETE', '2018-01-17 21:18:25', '2018-01-17 21:18:26');
INSERT INTO `rolegroups_modules` VALUES (1, 'assignments', 'CREATE', '2018-01-17 21:18:38', '2018-01-17 21:18:39');
INSERT INTO `rolegroups_modules` VALUES (1, 'assignments', 'READ', '2018-01-17 21:18:47', '2018-01-17 21:18:49');
INSERT INTO `rolegroups_modules` VALUES (1, 'assignments', 'UPDATE', '2018-01-20 21:18:56', '2018-01-17 21:19:00');
INSERT INTO `rolegroups_modules` VALUES (1, 'assignments', 'DELETE', '2018-01-17 21:19:06', '2018-01-17 21:19:07');
INSERT INTO `rolegroups_modules` VALUES (1, 'informations', 'CREATE', '2018-01-17 21:19:21', '2018-01-17 21:19:22');
INSERT INTO `rolegroups_modules` VALUES (1, 'informations', 'READ', '2018-01-17 21:19:30', '2018-01-17 21:19:31');
INSERT INTO `rolegroups_modules` VALUES (1, 'informations', 'UPDATE', '2018-01-17 21:19:37', '2018-01-17 21:19:39');
INSERT INTO `rolegroups_modules` VALUES (1, 'informations', 'DELETE', '2018-01-17 21:19:45', '2018-01-17 21:19:46');
INSERT INTO `rolegroups_modules` VALUES (1, 'tutorials', 'CREATE', '2018-01-17 21:19:55', '2018-01-17 21:19:56');
INSERT INTO `rolegroups_modules` VALUES (1, 'tutorials', 'READ', '2018-01-17 21:20:04', '2018-01-17 21:20:06');
INSERT INTO `rolegroups_modules` VALUES (1, 'tutorials', 'UPDATE', '2018-01-17 21:20:12', '2018-01-17 21:20:14');
INSERT INTO `rolegroups_modules` VALUES (1, 'tutorials', 'DELETE', '2018-01-17 21:20:27', '2018-01-17 21:20:28');
INSERT INTO `rolegroups_modules` VALUES (2, 'users', 'CREATE', '2019-01-07 21:49:03', '2019-01-07 21:49:07');
INSERT INTO `rolegroups_modules` VALUES (2, 'users', 'READ', '2018-01-17 21:14:47', '2018-01-17 21:14:48');
INSERT INTO `rolegroups_modules` VALUES (2, 'users', 'DELETE', '2018-01-17 21:15:14', '2018-01-17 21:15:15');
INSERT INTO `rolegroups_modules` VALUES (2, 'courses', 'CREATE', '2018-01-17 21:15:24', '2018-01-17 21:15:25');
INSERT INTO `rolegroups_modules` VALUES (2, 'courses', 'READ', '2018-01-17 21:15:32', '2018-01-17 21:15:34');
INSERT INTO `rolegroups_modules` VALUES (2, 'courses', 'UPDATE', '2018-01-17 21:15:41', '2018-01-17 21:15:48');
INSERT INTO `rolegroups_modules` VALUES (2, 'courses', 'DELETE', '2018-01-17 21:16:08', '2018-01-17 21:16:09');
INSERT INTO `rolegroups_modules` VALUES (2, 'attendances', 'CREATE', '2018-01-17 21:16:18', '2018-01-17 21:16:19');
INSERT INTO `rolegroups_modules` VALUES (2, 'attendances', 'READ', '2018-01-17 21:16:29', '2018-01-17 21:16:31');
INSERT INTO `rolegroups_modules` VALUES (2, 'attendances', 'UPDATE', '2018-01-17 21:16:43', '2018-01-17 21:16:44');
INSERT INTO `rolegroups_modules` VALUES (2, 'attendances', 'DELETE', '2018-01-17 21:16:51', '2018-01-17 21:16:52');
INSERT INTO `rolegroups_modules` VALUES (2, 'roles', 'CREATE', '2018-01-17 21:17:03', '2018-01-17 21:17:19');
INSERT INTO `rolegroups_modules` VALUES (2, 'roles', 'READ', '2018-01-17 21:17:17', '2018-01-17 21:17:18');
INSERT INTO `rolegroups_modules` VALUES (2, 'roles', 'UPDATE', '2018-01-17 21:17:26', '2018-01-17 21:17:27');
INSERT INTO `rolegroups_modules` VALUES (2, 'roles', 'DELETE', '2018-01-17 21:17:33', '2018-01-17 21:17:34');
INSERT INTO `rolegroups_modules` VALUES (2, 'schedules', 'CREATE', '2018-01-17 21:17:46', '2018-01-17 21:17:53');
INSERT INTO `rolegroups_modules` VALUES (2, 'schedules', 'READ', '2018-01-17 09:18:01', '2018-01-17 21:18:04');
INSERT INTO `rolegroups_modules` VALUES (2, 'schedules', 'UPDATE', '2018-01-17 21:18:17', '2018-01-17 21:18:18');
INSERT INTO `rolegroups_modules` VALUES (2, 'schedules', 'DELETE', '2018-01-17 21:18:25', '2018-01-17 21:18:26');
INSERT INTO `rolegroups_modules` VALUES (2, 'assignments', 'CREATE', '2018-01-17 21:18:38', '2018-01-17 21:18:39');
INSERT INTO `rolegroups_modules` VALUES (2, 'assignments', 'READ', '2018-01-17 21:18:47', '2018-01-17 21:18:49');
INSERT INTO `rolegroups_modules` VALUES (2, 'assignments', 'UPDATE', '2018-01-20 21:18:56', '2018-01-17 21:19:00');
INSERT INTO `rolegroups_modules` VALUES (2, 'assignments', 'DELETE', '2018-01-17 21:19:06', '2018-01-17 21:19:07');
INSERT INTO `rolegroups_modules` VALUES (2, 'informations', 'CREATE', '2018-01-17 21:19:21', '2018-01-17 21:19:22');
INSERT INTO `rolegroups_modules` VALUES (2, 'informations', 'READ', '2018-01-17 21:19:30', '2018-01-17 21:19:31');
INSERT INTO `rolegroups_modules` VALUES (2, 'informations', 'UPDATE', '2018-01-17 21:19:37', '2018-01-17 21:19:39');
INSERT INTO `rolegroups_modules` VALUES (2, 'informations', 'DELETE', '2018-01-17 21:19:45', '2018-01-17 21:19:46');
INSERT INTO `rolegroups_modules` VALUES (2, 'tutorials', 'CREATE', '2018-01-17 21:19:55', '2018-01-17 21:19:56');
INSERT INTO `rolegroups_modules` VALUES (2, 'tutorials', 'READ', '2018-01-17 21:20:04', '2018-01-17 21:20:06');
INSERT INTO `rolegroups_modules` VALUES (2, 'tutorials', 'UPDATE', '2018-01-17 21:20:12', '2018-01-17 21:20:14');
INSERT INTO `rolegroups_modules` VALUES (2, 'tutorials', 'DELETE', '2018-01-17 21:20:27', '2018-01-17 21:20:28');
INSERT INTO `rolegroups_modules` VALUES (3, 'users', 'READ', '2018-01-17 21:14:47', '2018-01-17 21:14:48');
INSERT INTO `rolegroups_modules` VALUES (3, 'courses', 'READ', '2018-01-17 21:15:32', '2018-01-17 21:15:34');
INSERT INTO `rolegroups_modules` VALUES (3, 'attendances', 'READ', '2018-01-17 21:16:29', '2018-01-17 21:16:31');
INSERT INTO `rolegroups_modules` VALUES (3, 'roles', 'READ', '2018-01-17 21:17:17', '2018-01-17 21:17:18');
INSERT INTO `rolegroups_modules` VALUES (3, 'roles', 'UPDATE', '2018-01-17 21:17:26', '2018-01-17 21:17:27');
INSERT INTO `rolegroups_modules` VALUES (3, 'schedules', 'READ', '2018-01-17 09:18:01', '2018-01-17 21:18:04');
INSERT INTO `rolegroups_modules` VALUES (3, 'schedules', 'UPDATE', '2018-01-17 21:18:17', '2018-01-17 21:18:18');
INSERT INTO `rolegroups_modules` VALUES (3, 'assignments', 'READ', '2018-01-17 21:18:47', '2018-01-17 21:18:49');
INSERT INTO `rolegroups_modules` VALUES (3, 'assignments', 'UPDATE', '2018-01-20 21:18:56', '2018-01-17 21:19:00');
INSERT INTO `rolegroups_modules` VALUES (3, 'informations', 'READ', '2018-01-17 21:19:30', '2018-01-17 21:19:31');
INSERT INTO `rolegroups_modules` VALUES (3, 'informations', 'UPDATE', '2018-01-17 21:19:37', '2018-01-17 21:19:39');
INSERT INTO `rolegroups_modules` VALUES (3, 'tutorials', 'READ', '2018-01-17 21:20:04', '2018-01-17 21:20:06');
INSERT INTO `rolegroups_modules` VALUES (3, 'tutorials', 'UPDATE', '2018-01-17 21:20:12', '2018-01-17 21:20:14');
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of schedules
-- ----------------------------
BEGIN;
INSERT INTO `schedules` VALUES (1, 1, 600, 800, 1, 'B', 3, 2019, '1', 'UDJT-0206', 6, '2018-01-17 19:40:54', '2018-01-17 19:40:56');
INSERT INTO `schedules` VALUES (2, 1, 600, 800, 2, 'B', 3, 2019, '2', 'UDJT-0207', 6, '2018-01-17 14:46:46', '2018-01-17 14:46:46');
INSERT INTO `schedules` VALUES (3, 1, 600, 800, 3, 'B', 3, 2019, '3', 'UDJT-305', 7, '2018-01-17 14:52:33', '2018-01-17 14:52:33');
INSERT INTO `schedules` VALUES (4, 1, 1600, 1800, 1, 'A', 3, 2019, '1', 'UDJT-0206', 6, '2019-01-11 23:38:20', '2019-01-11 23:38:26');
INSERT INTO `schedules` VALUES (5, 1, 1600, 1800, 2, 'A', 3, 2019, '2', 'UDJT-0207', 6, '2019-01-11 23:39:28', '2019-01-11 23:39:33');
INSERT INTO `schedules` VALUES (6, 1, 1600, 1800, 4, 'A', 3, 2019, '3', 'UDJT-305', 7, '2019-01-11 23:40:18', '2019-01-11 23:40:23');
COMMIT;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for types
-- ----------------------------
DROP TABLE IF EXISTS `types`;
CREATE TABLE `types` (
  `name` varchar(10) NOT NULL,
  `assignments_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`assignments_id`,`name`) USING BTREE,
  CONSTRAINT `fk_types_assignments_types` FOREIGN KEY (`assignments_id`) REFERENCES `assignments` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of types
-- ----------------------------
BEGIN;
INSERT INTO `types` VALUES ('docx', 5);
INSERT INTO `types` VALUES ('jpg', 5);
INSERT INTO `types` VALUES ('pdf', 5);
INSERT INTO `types` VALUES ('docx', 6);
INSERT INTO `types` VALUES ('jpg', 6);
INSERT INTO `types` VALUES ('pdf', 6);
COMMIT;

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
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'Mohammad Ghifari', 1, 'ghifari@gmail.com', '2af9b1ba42dc5eb01743e6b3759b6e4b', '', NULL, 2, '081220058838', 'ghifari12', '1140810140004', NULL, NULL, NULL, '2018-01-17 21:27:01', '2018-01-17 14:36:07');
INSERT INTO `users` VALUES (2, 'Febryani Pertiwi Puteri', 2, 'febryani@gmail.com', '2af9b1ba42dc5eb01743e6b3759b6e4b', '', NULL, 2, '081229949939', 'febry11', '1140810140028', NULL, NULL, NULL, '2018-01-17 21:26:23', '2018-01-17 14:38:52');
INSERT INTO `users` VALUES (3, 'Asep Nur Muhammad', 1, 'asepnur.isk@gmail.com', '2af9b1ba42dc5eb01743e6b3759b6e4b', 'note', NULL, 2, '081230081128', 'asep10', '1140810140070', NULL, NULL, NULL, '2018-01-17 19:42:18', '2018-01-17 19:42:19');
INSERT INTO `users` VALUES (4, 'Risal Falah', 1, 'risal.falah@gmail.com', '2af9b1ba42dc5eb01743e6b3759b6e4b', '', NULL, 2, '081230081123', 'risal12', '1140810140016', 9238, '2018-01-17 14:52:46', 0, '2018-01-17 14:22:46', '2018-01-17 14:22:46');
INSERT INTO `users` VALUES (5, 'Yunilucky Siswantari', 2, 'yuni@gmail.com', '2af9b1ba42dc5eb01743e6b3759b6e4b', '', NULL, 2, '081274327712', 'yuni22', '1140810140029', NULL, NULL, NULL, '2018-01-17 21:27:51', '2018-01-17 21:27:52');
INSERT INTO `users` VALUES (6, 'Hidayaturrahman', 1, 'hidayat@gmail.com', '2af9b1ba42dc5eb01743e6b3759b6e4b', '', 2, 2, '081289089903', 'dayatura', '11140810140050', NULL, NULL, NULL, '2018-01-17 21:29:05', '2018-01-17 14:32:44');
INSERT INTO `users` VALUES (7, 'Muhammad Rifki', 1, 'rifki@gmail.com', '2af9b1ba42dc5eb01743e6b3759b6e4b', '', 2, 2, '081234561123', 'rifkimm', '1140810140020', NULL, NULL, NULL, '2018-01-17 21:30:01', '2018-01-17 14:37:38');
INSERT INTO `users` VALUES (8, 'fahmi irfan', 0, 'fahmiirfan909@gmail.com', '766d89ebaf59fddfb6a9d6b338ddf2d8', '', 3, 2, '081234781176', 'fahmi', '1140810160028', NULL, NULL, NULL, '2018-04-27 06:41:08', '2018-07-29 10:26:48');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
