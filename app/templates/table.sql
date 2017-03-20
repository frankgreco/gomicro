CREATE DATABASE IF NOT EXISTS <%= nounPluralLower %>;

USE <%= nounPluralLower %>;

CREATE TABLE <%= nounPluralLower %>(
    id          INT             AUTO_INCREMENT  PRIMARY KEY,
    paramOne    VARCHAR(255),
    paramTwo    VARCHAR(255)
);
