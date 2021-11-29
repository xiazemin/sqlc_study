CREATE TABLE `users` (
  `user_id` varchar(256) NOT NULL,
  `labels` json NOT NULL,
  `comment` text NOT NULL,
  `subscribe` json NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;