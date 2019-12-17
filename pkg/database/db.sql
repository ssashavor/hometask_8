DROP TABLE IF EXISTS ‘contact‘;
CREATE TABLE ‘contact‘ (
    id int8 NOT NULL PRIMARY KEY,
    firstname varchar(255) NOT NULL,
    lastname varchar(255) NOT NULL,
    phone varchar(255) NOT NULL,
    email varchar(255) NOT NULL
);
