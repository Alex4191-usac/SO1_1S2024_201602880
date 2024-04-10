CREATE DATABASE IF NOT EXISTS `tarea4`
USE `tarea4`

CREATE TABLE IF NOT EXISTS `tarea4`.`music_liderboard`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(150) NOT NULL,
    `album` VARCHAR(100) NOT NULL,
    `year` VARCHAR(100) NOT NULL,
    `rank` VARCHAR(50) NOT NULL,
    PRIMARY KEY(`id`)

)