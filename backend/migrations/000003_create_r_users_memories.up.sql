CREATE TABLE `r_users_memories` (
    `user_id` int(255) NOT NULL,
    `memory_id` int(255) NOT NULL,
    PRIMARY KEY (`user_id`, `memory_id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES users(`id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT `fk_memory_id` FOREIGN KEY (`memory_id`) REFERENCES memories(`id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE
) ENGINE = InnoDB;