create database if not exists db_contacts;

CREATE TABLE db_contacts.contact (
	id INT auto_increment NOT NULL,
	name VARCHAR(255) NOT NULL,
	email varchar(255) NULL,
	phone VARCHAR(20) NULL,
	CONSTRAINT contact_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

insert into db_contacts.contact (name, email, phone) values("Juan Acosta", "juancho@email.com", "365-956-7854"), ("Maria Camila", "MC@email.com", "324-568-7941"), ("Melo", "Caramelo@email.com", "318-472-1473");

insert into db_contacts.contact (name, email, phone) values("Bandy", null, "314-231-4444");

insert into db_contacts.contact (name, email, phone) values("Isabella", "Isix@email.com", null);