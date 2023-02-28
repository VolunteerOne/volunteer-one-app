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
 `verified`   tinyint,

PRIMARY KEY (`id`)
);

CREATE TABLE `events` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `organization_id` int NOT NULL,
  `address` varchar(255),
  `date` date NOT NULL,
  `time` time NOT NULL,
  `description` text,

  PRIMARY KEY (`id`),
  FOREIGN KEY (`organization_id`) REFERENCES `volunteer_req` (`event_id`)
);