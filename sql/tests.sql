CREATE table Customer (
cst_id SERIAL PRIMARY KEY,
nationality_id int not null,
cst_email varchar(50) not null,
cst_phoneNum varchar(20) not null,
cst_dobdate varchar(20) not null,
CONSTRAINT fk_CustomerNationality FOREIGN KEY (nationality_id) REFERENCES Nationality(nationality_id) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE table Family_list(
fl_id SERIAL PRIMARY KEY,
cst_id int not null,
fl_relation varchar(50) not null,
fl_name varchar(50) not null,
fl_dob varchar(50)not null,
CONSTRAINT fk_CustomerFamily FOREIGN KEY (cst_id) REFERENCES Customer(cst_id) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE table Nationality(
nationality_id SERIAL PRIMARY KEY,
nationality_name VARCHAR(50) NOT NULL,
nationality_code CHAR(2) NOT NULL
);

SELECT * FROM Nationality;
SELECT * FROM Customer;
SELECT * FROM Family_list;