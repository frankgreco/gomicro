CREATE DATABASE <%= nounPluralLower %>;

CREATE TABLE <%= nounPluralLower %> (
    id           BIGSERIAL      PRIMARY KEY,
    param_one    VARCHAR(255),
    param_two    VARCHAR(255)
);
