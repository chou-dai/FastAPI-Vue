CREATE TABLE `users` (
    `id` int(255) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(200) NOT NULL UNIQUE,
    `password` VARCHAR(200) NOT NULL,
    `session_id` VARCHAR(200) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
);