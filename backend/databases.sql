CREATE TABLE `users`
(
 `id`         int NOT NULL ,
 `handle`     VARCHAR(255) NOT NULL ,
 `email`      VARCHAR(255) NOT NULL ,
 `password`   VARCHAR(255) NOT NULL ,
 `birthdate`  date NOT NULL ,
 `firstName`  VARCHAR(255) NOT NULL ,
 `lastName`   VARCHAR(255) NOT NULL ,
 `profilePic` mediumblob,
 `interests`  VARCHAR(255) ,

PRIMARY KEY (`id`)
);