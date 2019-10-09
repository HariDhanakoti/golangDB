# golangDB
Golang database connectivity and retrieval options

Pull latest mysql:
docker pull mysql/mysql-server:latest

Run Database:
docker run -p 3306:3306 -d --name mysql -e MYSQL_ROOT_PASSWORD=password mysql/mysql-server

Login to the container bash
mysql -uroot -ppassword

Create user:
CREATE USER 'golang'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON * . * TO 'golang'@'%';

Test created user login:
mysql -ugolang -ppassword


Create schema:
CREATE DATABASE demo;

USE demo;

CREATE TABLE users ( id smallint unsigned not null auto_increment, name varchar(20) not null, constraint pk_example primary key (id) );
INSERT INTO users ( id, name ) VALUES (1, "Ashlee");
INSERT INTO users ( id, name ) VALUES (2, "Brown");
INSERT INTO users ( id, name ) VALUES (3, "Daniel");
INSERT INTO users ( id, name ) VALUES (4, "Samantha");
INSERT INTO users ( id, name ) VALUES (5, "Brady");

