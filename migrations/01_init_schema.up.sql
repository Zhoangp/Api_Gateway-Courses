DROP TABLE IF EXISTS `Users`;
CREATE TABLE `Users`
(
    `id`            int          NOT NULL AUTO_INCREMENT,
    `firstName`     varchar(255) NOT NULL,
    `lastName`      varchar(255) NOT NULL,
    `email`         varchar(255) NOT NULL,
    `password`      varchar(255) NOT NULL,
    `phoneNumber`   varchar(255) NOT NULL,
    `address`       varchar(255) NOT NULL,
    `role`          enum('guest','admin') NOT NULL DEFAULT 'guest',
    `is_instructor` tinyint(1) DEFAULT '0',
    `created_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`    timestamp NULL DEFAULT NULL,
    `picture`       json DEFAULT NULL,
    `verified`      tinyint(1) DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;