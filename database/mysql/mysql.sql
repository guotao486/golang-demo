create database go_db

CREATE TABLE user_tbl (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  username VARCHAR (20),
  PASSWORD VARCHAR (20)
)

INSERT INTO user_tbl (username, PASSWORD) VALUES ("tom", "123")
INSERT INTO user_tbl (username, PASSWORD) VALUES ("kite", "456")