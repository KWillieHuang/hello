CREATE TABLE `user` (
    `userid` INT(10) NOT NULL AUTO_INCREMENT, 
    `username` VARCHAR(32) NOT NULL UNIQUE,
    `password` VARCHAR(32) NOT NULL,
    `email` VARCHAR(32) NOT NULL UNIQUE,
    `role` INT(5) NOT NULL,
    PRIMARY KEY(`userid`)
);
INSERT INTO user VALUES("1","admin","4fdebe67f1dbe1e7e7f002f73bc86e1b","1@celes.com","2");
INSERT INTO user VALUES("2","user1","81dc9bdb52d04dc20036dbd8313ed055","12@celes.com","1");
INSERT INTO user VALUES("3","user2","4fdebe67f1dbe1e7e7f002f73bc86e1b","123@celes.com","1");