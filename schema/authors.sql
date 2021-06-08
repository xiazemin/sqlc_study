CREATE TABLE `authors` (
  `id` int NOT NULL,
  `name` text NOT NULL,
  `bio` varchar(255) DEFAULT NULL,
  `company_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;