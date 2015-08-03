/*drop database cnt;*/
CREATE DATABASE cnt DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
use cnt;
create table Товар
(
Артикул int NOT NULL AUTO_INCREMENT,
Наименование varchar(255),
Цена double,
Склад1 int,
Склад2 int,
Склад3 int,
PRIMARY KEY (Артикул)
);
INSERT INTO Товар (Наименование, Цена, Склад1, Склад2, Склад3)
VALUES 
	('Скрепки', 10, 0, 50, 0),
	('Гвозди', 20, 10, 0, 30),
	('Кнопки', 30, 20, 10, 10), 
	('Шурупы', 40, 0, 0, 10);
    