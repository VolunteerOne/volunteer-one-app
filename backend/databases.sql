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

CREATE TABLE `organizations`
(
 `id`           int NOT NULL ,
 `name`         VARCHAR(255) NOT NULL ,
 `description`  VARCHAR(255) ,
 `verified`     tinyint ,
 `interests`    VARCHAR(255) ,

PRIMARY KEY (`id`)
);

CREATE TABLE `org_roles`
(
 `id`           int NOT NULL ,
 `admin_id`     int NOT NULL ,
 `org_id`       int NOT NULL ,
 `role`         set('owner', 'manager', 'member') NOT NULL ,

PRIMARY KEY (`id`) ,
FOREIGN KEY (`admin_id`) REFERENCES `users` (`id`) ,
FOREIGN KEY (`org_id`) REFERENCES `organizations` (`id`)
);

CREATE TABLE `friends`
(
 `id`           int NOT NULL AUTO_INCREMENT ,
 `sender_id`    int NOT NULL ,
 `receiver_id`  int NOT NULL ,
 `status`       set('accepted','pending','rejected') DEFAULT 'pending' NOT NULL ,

 PRIMARY KEY (`id`),
 FOREIGN KEY (`sender_id`) REFERENCES `users`(`id`),
 FOREIGN KEY (`sender_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `user_social`
(
 `id`           int NOT NULL AUTO_INCREMENT ,
 `user_id`      int NOT NULL ,
 `created_at`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `content`      text NOT NULL ,
 `hidden`       tinyint NOT NULL DEFAULT 0 ,

 PRIMARY KEY (`id`),
 FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);