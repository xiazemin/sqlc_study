CREATE TABLE `measurements` (
  `name` varchar(256) NOT NULL,
  `time_stamp` timestamp NOT NULL,
  `value` double NOT NULL,
  `comment` text NOT NULL,
  PRIMARY KEY (`name`,`time_stamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;