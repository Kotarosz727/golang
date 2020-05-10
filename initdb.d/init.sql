DROP DATABASE IF EXISTS test;
CREATE DATABASE test;
USE test;
DROP TABLE IF EXISTS posts;
 
CREATE TABLE posts (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
content TEXT NOT NULL,
author varchar(255)
)DEFAULT CHARACTER SET=utf8;
 
INSERT INTO posts (content,author) 
VALUES 
("test1", "Bob")
,("test2", "Jack")
,("test3", "John");