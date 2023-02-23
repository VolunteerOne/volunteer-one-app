CREATE TABLE `users`
(
 `id`         int NOT NULL ,
 `handle`     linestring NOT NULL ,
 `email`      linestring NOT NULL ,
 `password`   linestring NOT NULL ,
 `birthdate`  date NOT NULL ,
 `firstName`  linestring NOT NULL ,
 `lastName`   linestring NOT NULL ,
 `profilePic` mediumblob,
 `interests`  text ,

PRIMARY KEY (`id`)
);