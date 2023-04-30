
DROP TABLE IF EXISTS `Users`;
DROP TABLE IF EXISTS `Instructors`;
DROP TABLE IF EXISTS `Courses`;

SET SQL_MODE='ALLOW_INVALID_DATES';
CREATE TABLE `Users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `firstName` varchar(255) NOT NULL,
                         `lastName` varchar(255) NOT NULL,
                         `email` varchar(255) NOT NULL,
                         `password` varchar(255) NOT NULL,
                         `phoneNumber` varchar(255) NOT NULL,
                         `address` varchar(255) NOT NULL,
                         `role` enum('guest','admin') NOT NULL DEFAULT 'guest',
                         `is_instructor` tinyint(1) DEFAULT '0',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` timestamp NULL DEFAULT NULL,
                         `picture` json DEFAULT NULL,
                         `lastLogin` timestamp NULL DEFAULT '0000-00-00 00:00:00',
                         `verified` tinyint(1) DEFAULT '0',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `Users` (`id`, `firstName`, `lastName`, `email`, `password`, `phoneNumber`, `address`, `role`, `is_instructor`, `created_at`, `updated_at`, `deleted_at`, `picture`, `lastLogin`, `verified`)
VALUES
    (1, 'Nguyen Van', 'Nui', 'zhoangp@gmail.com', '$2a$10$sfhixhLkD3JciHzsVKeW7Oa1pWwTzuj6cBe60TU8/bvQ2i0w1Q5iG', '0922212121', 'cccc', 'guest', 1, '2023-03-10 01:06:53', '2023-04-20 11:01:30', NULL, NULL, '0000-00-00 00:00:00', 0);
CREATE TABLE `Instructors` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `user_id` int NOT NULL,
                               `website` varchar(255) DEFAULT NULL,
                               `linkedin` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                               `youtube` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                               `bio` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                               `total_students` int DEFAULT '0',
                               `total_reviews` int DEFAULT '0',
                               `rating` float DEFAULT '0',
                               `total_courses` bigint DEFAULT '0',
                               `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                               `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               `deleted_at` timestamp NULL DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `user_id` (`user_id`) USING BTREE,
                               CONSTRAINT `Instructors_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `Instructors` (`id`, `user_id`, `website`, `linkedin`, `youtube`, `bio`, `total_students`, `total_reviews`, `rating`, `total_courses`, `created_at`, `updated_at`, `deleted_at`) VALUES
    (1, 1, NULL, NULL, NULL, NULL, 0, 0, 0, 0, '2023-03-09 19:34:19', '2023-03-09 19:34:19', NULL);
CREATE TABLE `Courses` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `description` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `level` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `language` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                           `price` decimal(12,2) NOT NULL,
                           `discount` float DEFAULT NULL,
                           `duration` time NOT NULL,
                           `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                           `rating` float DEFAULT NULL,
                           `instructor_id` int NOT NULL,
                           `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           `deleted_at` timestamp NULL DEFAULT NULL,
                           `currency` varchar(20) NOT NULL,
                           `thumbnail` json NOT NULL,
                           PRIMARY KEY (`id`),
                           KEY `instructor_id` (`instructor_id`),
                           CONSTRAINT `Courses_ibfk_2` FOREIGN KEY (`instructor_id`) REFERENCES `Instructors` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `Courses` (`id`, `title`, `description`, `level`, `language`, `price`, `discount`, `duration`, `status`, `rating`, `instructor_id`, `created_at`, `updated_at`, `deleted_at`, `currency`, `thumbnail`)
VALUES
    (1, 'Python For Beginners', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}'),
    (2, 'Photograph', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}'),
    (3, 'Adobe Photoshop', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}'),
    (4, 'Adobe XD', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}'),
    (5, 'Basic JS', 'This python for beginners course is geared to students who want to know how python works and also to those totally new to programming.\n\nThe python language has very simple syntax(way to write it) to learn and it is one of the most powerful languages to learn since it can used for a variety of things.\n\nHere are some:\n\nData analysis\nGame development\nVisualization\nWeb development\nRobotics\nand more....\n\n\nJobs in this field are really lucrative and knowing this language will give you an edge when finding a job and making a lot more money than other developers; python developers are not as many as in other languages since people think is hard. Python is super easy to learn but very powerful since it contains many possibilities.\n\nPython is growing faster and faster everyday and it has surpassed many other languages over the years for a lot of reasons, which you will find out soon enough.\n\n------------------------------------------------------------------------------\n\nWhy take this course?\n\nThis course starts with explaining what programming really is? Have you ever wondered how things actually work in a program?\n\n1. Teaches the foundation of GENERAL programming\n\nEvery student should have some foundations on what programming really is before learning any language, why? Because once you understand the core components of programming it will be a lot easier to learn any language and create better programs.\n\n2. New lectures added all the time\n\nWhen you like what you do, it reflects. This is not a job for me, so I wake up wanting to code and help my students. Basically you pay once, get the course forever and get extra values added all the time, you have to love that :)\n\n3. Fun place to be\n\nI love making my lectures fun and engaging, so no boring lectures here!\n\n4. The support you get in this course in unmatched\n\nhave you ever joined a class and received very little support or none at all? Well that is not going to happen in this class, because I love helping my students.\n\n\n\nWho this course is for:\nStudent totally new to programming\nStudent totally new to python\n',
     'Beginner', 'English', 13.52, NULL, '12:15:00', NULL, NULL, 1, '2023-03-09 19:34:27', '2023-04-29 16:10:27', NULL, 'usd', '{"id": 1, "url": "https://d338hmodjsosv2.cloudfront.net/img/avatar-zhoangp", "width": "250px", "height": "250px"}');

DROP TABLE IF EXISTS `Sections`;
CREATE TABLE `Sections` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `course_id` int NOT NULL,
                            `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `numberOfLectures` int NOT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `course_id` (`course_id`),
                            CONSTRAINT `Sections_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `Courses` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `Sections` (`id`, `course_id`, `title`, `numberOfLectures`, `created_at`, `updated_at`, `deleted_at`)
VALUES
    (1, 3, 'Introduce', 3, '2023-04-25 02:44:44', '2023-04-25 02:48:09', NULL),
    (2, 3, 'Main', 3, '2023-04-25 02:44:44', '2023-04-25 02:48:09', NULL);


DROP TABLE IF EXISTS `Lectures`;
CREATE TABLE `Lectures` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `section_id` int NOT NULL,
                            `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `section_id` (`section_id`),
                            CONSTRAINT `Lectures_ibfk_1` FOREIGN KEY (`section_id`) REFERENCES `Sections` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
INSERT INTO `Lectures` (`id`, `section_id`, `title`, `content`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
                                                                                                                        (1, 1, 'Hello', 'abc', 'ok', '2023-04-25 02:47:16', '2023-04-25 02:47:16', NULL),
                                                                                                                        (2, 2, 'Hello 2', 'xyz', 'ok', '2023-04-25 02:47:16', '2023-04-25 17:41:00', NULL);

DROP TABLE IF EXISTS `LectureResources`;
CREATE TABLE `LectureResources` (
                                    `id` int NOT NULL AUTO_INCREMENT,
                                    `lecture_id` int NOT NULL,
                                    `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                    `duration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                                    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    `deleted_at` timestamp NULL DEFAULT NULL,
                                    PRIMARY KEY (`id`),
                                    KEY `lecture_id` (`lecture_id`),
                                    CONSTRAINT `LectureResources_ibfk_1` FOREIGN KEY (`lecture_id`) REFERENCES `Lectures` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `LectureResources` (`id`, `lecture_id`, `url`, `duration`, `created_at`, `updated_at`, `deleted_at`) VALUES
    (1, 1, 'https://courses-storages.s3.ap-northeast-1.amazonaws.com/video/Golang+installation+and+hello+world.mp4', '5:00', '2023-04-25 03:15:36', '2023-04-25 11:02:41', NULL);
