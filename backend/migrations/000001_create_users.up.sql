CREATE TABLE `users` (
    `id` int(255) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(10) NOT NULL UNIQUE,
    `password` VARCHAR(30) NOT NULL,
    `session_id` VARCHAR(60) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
);