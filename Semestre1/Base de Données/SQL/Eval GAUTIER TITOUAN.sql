CREATE TABLE Fournisseur (
nuf number(2) PRIMARY KEY,
nomf varchar(10) NOT NULL,
ville varchar(15) NOT NULL
)

CREATE TABLE Fourniture (
nuf number(2),
nup number(5),
qte number(4) check(qte > 0),
constraint FK_Fourniture FOREIGN key (nuf) references Fournisseur(nuf),
constraint FK_Produit foreign key (nup) references Fourniture(nup),
constraint CLE_Fournisseur PRIMARY KEY (nuf,nup)
)

CREATE TABLE Produit (
nup number(5) PRIMARY KEY,
nomp varchar(10) NOT NULL,
couleur varchar(5) ,
origine varchar(15) NOT NULL,
pvu number(5) check(0 < pvu) check(pvu <= 1000)
)

Insert into Fournisseur values('1','DUPOND',Nantes);
Insert into Fournisseur values('1','Jean', Nantes);

Insert into Produit(1,'prod1','vert','Nantes',100);
Insert into Produit(2,'prod2','vert','Nantes',100);
Insert into Produit(3,'prod3','vert','Nantes',100);

Select nuf,nomf from fournisseur natural join produit where nuf = nup;
