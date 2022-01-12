CREATE TABLE `authors` (
  `id` int NOT NULL,
  `name` text NOT NULL,
  `bio` varchar(255) DEFAULT NULL,
  `company_id` int NOT NULL,
  `size` bigint DEFAULT NULL,
  `empty_col` int DEFAULT NULL,
  `default_col` int NOT NULL,
  `size1` int DEFAULT NULL,
  `default_col1` int NOT NULL,
  `type` enum('互联网/游戏/软件','教育/培训','1-30人') DEFAULT NULL,
  `type1` enum('互联网/游戏/软件','教育/培训','1-30人') DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;