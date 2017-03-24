CREATE DATABASE IF NOT EXISTS <%= nounPluralLower %>;

USE <%= nounPluralLower %>;

CREATE TABLE <%= nounPluralLower %>(
    id           INT             AUTO_INCREMENT  PRIMARY KEY,
    param_one    VARCHAR(255),
    param_two    VARCHAR(255)
);
