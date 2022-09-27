CREATE TABLE `students` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `ruby` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `student_number` int unsigned NOT NULL,
  `date_of_birth` datetime NOT NULL,
  `address` varchar(256) COLLATE utf8mb4_bin NOT NULL,
  `expiration_date` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `student_number` (`student_number`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin