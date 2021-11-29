CREATE TABLE `feedback` (
  `feedback_type` varchar(256) NOT NULL,
  `user_id` varchar(256) NOT NULL,
  `item_id` varchar(256) NOT NULL,
  `time_stamp` timestamp NOT NULL,
  `comment` text NOT NULL,
  PRIMARY KEY (`feedback_type`,`user_id`,`item_id`),
  KEY `user_id` (`user_id`),
  KEY `user_id_2` (`user_id`),
  KEY `user_id_3` (`user_id`),
  KEY `user_id_4` (`user_id`),
  KEY `user_id_5` (`user_id`),
  KEY `user_id_6` (`user_id`),
  KEY `user_id_7` (`user_id`),
  KEY `user_id_8` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;