CREATE TABLE `memories` (
    `id` int(255) NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(100) NOT NULL,
    `description` VARCHAR(255),
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `is_public` tinyint(1) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
);