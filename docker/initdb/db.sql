CREATE DATABASE golang;
USE golang;

DROP TABLE IF EXISTS `employees`;
CREATE TABLE `employees` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `email` VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

LOCK TABLES `employees` WRITE;

INSERT INTO `employees` VALUES
    (2000, 'Steve', 'steve@work.com'),
    (2001, 'Arian', 'arian@work.com'),
    (2002, 'Rob', 'rob@work.com'),
    (2003, 'Samuel', 'samuel@work.com');

UNLOCK TABLES;