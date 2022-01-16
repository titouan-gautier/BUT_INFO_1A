create table Personne(
nucond number(2) primary key,
nomcond varchar2(30) not null,
nompass varchar2(30) not null,
nuag varchar2(15) not null,
dateag varchar(10) not null,
mail varchar2(50),
nutel varchar2(10) not null,
genre varchar2(20),
age number(2) not null
);

create table VilleD(
nuvilled number(2) primary key,
nomvilled varchar2(50) not null
);

create table Avis(
nuavis number(2) primary key,
etoile number(1) not null,
nompass varchar2(30) not null,
nucond number(2),
constraint Avis foreign key(nucond) references Personne,
comm varchar2(500)
);

create table Vehicule(
nuvehi number(2) primary key,
marque varchar2(20) not null,
typev varchar2(20) not null,
energie varchar2(20) not null,
dcircu number(4) not null,
nbplace varchar(1) not null
);

create table Trajet(
nutrajet number(2) primary key,
nomvilleA varchar(50) not null,
hdep varchar(8) not null,
harr varchar(8) not null,
timetrajet varchar(8) not null,
prix number(4) not null
)

insert into Personne (nucond, nomcond, nompass, nuag, dateag, mail, nutel, genre, age) values(01, 'Dupond', 'Dupond', 'FR4415603CE', '15/11/2021', 'dupond01@gmail.com', 0654221400, 'homme', 25);
insert into Personne (nucond, nomcond, nompass, nuag, dateag, nutel, genre, age) values (02, 'Michelle', 'Jules-Titouan', 'FR1717001CE', '01/11/1989', 0748251502, 'femme', 60);
insert into Personne (nucond, nomcond, nompass, nuag, dateag, mail, nutel, genre, age) values (03, 'Schweppes', 'Paul', 'CH5532502CE', '11/11/2010', 'mr.schweppes@gmail.com', 0225391383,'homme', 40);

insert into VilleD (nuvilled, nomvilled) values (01, 'Paris');
insert into VilleD (nuvilled, nomvilled) values (02, 'Nantes');
insert into VilleD (nuvilled, nomvilled) values (03, 'Nantes');

insert into Avis (nuavis, etoile, nompass, nucond) values (01, 5, 'Dupond', 01);
insert into Avis (nuavis, etoile, nompass, nucond, comm) values (02, 2, 'Jules', 02, 'In omnia paratus...');
insert into Avis (nuavis, etoile, nompass, nucond, comm) values (03, 3, 'Titouan', 02, 'Duis aute irure dolor...');
insert into Avis (nuavis, etoile, nompass, nucond, comm) values (04, 4, 'Paul', 03, 'Carpe diem.');

insert into Vehicule (nuvehi, marque, typev, energie, nbplace, dcircu) values (01, 'Tesla', 'Model 3', 'electrique', 5, 2020);
insert into Vehicule (nuvehi, marque, typev, energie, nbplace, dcircu) values (02, 'Fiat', 'Multipla', 'essence', 6, 1999);
insert into Vehicule (nuvehi, marque, typev, energie, nbplace, dcircu) values (03, 'Renault', 'Alpine V6 GT', 'essence', 5, 1989);

insert into Trajet (nutrajet, nomvilleA, hdep, harr, timetrajet, prix) values (01, 'Nantes','14h20', '18h40', '4h20', 30);
insert into Trajet (nutrajet, nomvilleA, hdep, harr, timetrajet, prix) values (02, 'Nantes', '6h', '6h20', '20min', 5);
insert into Trajet (nutrajet, nomvilleA, hdep, harr, timetrajet, prix) values (03, 'Cholet', '17h', '17h50', '50min', 10);

drop table avis;
drop table personne;
drop table vehicule;
drop table villed;
drop table trajet;