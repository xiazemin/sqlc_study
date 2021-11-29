CREATE TABLE `items` (
  `item_id` varchar(256) NOT NULL,
  `time_stamp` timestamp NOT NULL,
  `labels` json NOT NULL,
  `comment` text NOT NULL,
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;