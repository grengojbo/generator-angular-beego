CREATE TABLE auth_user (
  id INT(11) NOT NULL AUTO_INCREMENT,
  password VARCHAR(128) NOT NULL,
  last_login DATETIME NOT NULL,
  is_superuser TINYINT(1) NOT NULL,
  username VARCHAR(30) NOT NULL,
  first_name VARCHAR(30) NOT NULL,
  last_name VARCHAR(30) NOT NULL,
  email VARCHAR(75) NOT NULL,
  is_staff TINYINT(1) NOT NULL,
  is_active TINYINT(1) NOT NULL,
  date_joined DATETIME NOT NULL,
  PRIMARY KEY (id),
  UNIQUE INDEX username USING BTREE (username)
)
ENGINE = TOKUDB
CHARACTER SET utf8
COLLATE utf8_general_ci;