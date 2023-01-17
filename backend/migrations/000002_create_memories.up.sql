CREATE TABLE `memories` (
    `id` int(255) NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(100) NOT NULL,
    `description` VARCHAR(255),
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `is_public` tinyint(1) NOT NULL DEFAULT '0',
    `user_id` int(255) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES users(`id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);