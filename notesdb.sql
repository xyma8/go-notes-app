create database gonotesdb

USE gonotesdb;

CREATE TABLE `notes` (
	`id` VARCHAR(36) KEY,
    `title` VARCHAR(128),
    `note_text` MEDIUMTEXT NOT NULL,
    `author` VARCHAR(64),
    `creation_time` TIMESTAMP NOT NULL,
    `is_public` BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO `notes` (`id`, `title`, `note_text`, `author`, `creation_time`, `is_public`)
VALUES(UUID(), 'привет', 'привет привет привет', 'pul', '2024-05-09 01:15:21', 1)