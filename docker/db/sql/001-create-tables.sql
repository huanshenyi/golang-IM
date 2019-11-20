---- drop ----
DROP TABLE IF EXISTS `chat`;

---- create ----
create table IF not exists `chat`
(
 `id`               INT (64) AUTO_INCREMENT,
 `mobile`             VARCHAR (20) NOT NULL,
 `passwd`       VARCHAR (40) NOT NULL,
 `avatar`       VARCHAR (150) DEFAULT NULL,
 `sex`          VARCHAR (2) DEFAULT NULL,
 `Nickname`     VARCHAR (20) DEFAULT NULL,
 `salt`         VARCHAR (10) DEFAULT  NULL,
 `online`       INT (10) DEFAULT NULL,
 `token`        VARCHAR (40) DEFAULT NULL,
 `memo`         VARCHAR (140) DEFAULT NULL,
 `created_at`   Datetime DEFAULT NULL,
 `updated_at`   Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;