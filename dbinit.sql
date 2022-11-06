DROP DATABASE "is-theatre";
CREATE DATABASE "is-theatre";

DROP TABLE IF EXISTS AgeRating CASCADE ;
CREATE TABLE IF NOT EXISTS AgeRating (
    id SERIAL PRIMARY KEY,
    rating VARCHAR(255) DEFAULT '' UNIQUE
);

DROP TABLE IF EXISTS Genre CASCADE ;
CREATE TABLE IF NOT EXISTS Genre (
    genre VARCHAR(255) DEFAULT '' UNIQUE,
    id SERIAL PRIMARY KEY
) ;


DROP TABLE IF EXISTS Author CASCADE ;
CREATE TABLE IF NOT EXISTS Author (
    name VARCHAR(255) DEFAULT '',
    secondName VARCHAR(255) DEFAULT '',
    country VARCHAR(255) DEFAULT '',
    century VARCHAR(255) DEFAULT '',
    id SERIAL PRIMARY KEY
) ;

DROP TABLE IF EXISTS Composition CASCADE ;
CREATE TABLE IF NOT EXISTS Composition (
     compositionName VARCHAR(255) DEFAULT '',
     description VARCHAR DEFAULT '',
     authorId INT DEFAULT '0' REFERENCES Author(id),
     genreId INT DEFAULT '0' REFERENCES Genre(id),
     ageRatingId INT DEFAULT '0' REFERENCES AgeRating(id),
     id SERIAL PRIMARY KEY
);

DROP TABLE IF EXISTS Personage CASCADE ;
CREATE TABLE IF NOT EXISTS Personage (
    personageName VARCHAR(255) DEFAULT 'Mary Sue',
    compositionId INT NOT NULL DEFAULT '0' REFERENCES Composition(id),
    id SERIAL PRIMARY KEY
) ;

DROP TABLE IF EXISTS PersonageDescription CASCADE ;
CREATE TABLE IF NOT EXISTS PersonageDescription (
    actorId INT DEFAULT '0' REFERENCES Personage(id),
    age VARCHAR(255) DEFAULT '',
    voice VARCHAR(255) DEFAULT '',
    height VARCHAR(255) DEFAULT '',
    weight VARCHAR(255) DEFAULT '',
    gender BOOL DEFAULT '0',
    description VARCHAR DEFAULT '',
    id SERIAL PRIMARY KEY
);

DROP TABLE IF EXISTS Worker CASCADE ;
CREATE TABLE IF NOT EXISTS Worker (
    name VARCHAR(255) DEFAULT '',
    secondName VARCHAR(255) DEFAULT '',
    birthDate DATE DEFAULT '2000-01-01',
    salary DECIMAL DEFAULT '1000.00',
    id SERIAL PRIMARY KEY
);

DROP TABLE IF EXISTS Director CASCADE;
CREATE TABLE IF NOT EXISTS Director(
    workerId INT DEFAULT '0' REFERENCES Worker(id),
    id SERIAL PRIMARY KEY
);

DROP TABLE IF EXISTS Actor CASCADE ;
CREATE TABLE IF NOT EXISTS Actor (
    workerId INT DEFAULT '0' REFERENCES Worker(id),
    experience INT DEFAULT 0,
    id SERIAL PRIMARY KEY
);

DROP TABLE IF EXISTS ActorDescription CASCADE ;
CREATE TABLE IF NOT EXISTS ActorDescription (
    actorId INT DEFAULT '0' REFERENCES Actor(id),
    age VARCHAR(255) DEFAULT '',
    voice VARCHAR(255) DEFAULT '',
    height VARCHAR(255) DEFAULT '',
    weight VARCHAR(255) DEFAULT '',
    gender BOOL DEFAULT '0',
    id SERIAL PRIMARY KEY
);

DROP TABLE IF EXISTS Performance CASCADE ;
CREATE TABLE IF NOT EXISTS Performance (
    compositionId INT DEFAULT '0' REFERENCES Composition(id),
    performanceName VARCHAR(255) DEFAULT '',
    directorId INT DEFAULT 0 REFERENCES Director(id),
    performanceDate DATE DEFAULT '2000-01-01',
    performanceTime TIME DEFAULT '00:00:00',
    id SERIAL PRIMARY KEY
) ;

DROP TABLE IF EXISTS ActorsRole CASCADE;
CREATE TABLE IF NOT EXISTS ActorsRole(
    performanceId INT DEFAULT '0' REFERENCES Performance(id),
    actorId INT DEFAULT 0 REFERENCES Actor(id),
    id SERIAL PRIMARY KEY
);


DROP TABLE IF EXISTS Repertoire CASCADE ;
CREATE TABLE IF NOT EXISTS Repertoire (
    periodBegin DATE DEFAULT '2000-02-01',
    periodEnd DATE DEFAULT '2000-01-01',
    id SERIAL PRIMARY KEY
) ;

DROP TABLE IF EXISTS Userr CASCADE ;
CREATE TABLE IF NOT EXISTS Userr (
    nickName VARCHAR(255) DEFAULT '',
    email VARCHAR(255) CHECK (email ~* '[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+'),
    password VARCHAR(255) DEFAULT '',
    role VARCHAR(255) check ( role = 'admin' or role = 'user' ),
    verified BOOLEAN DEFAULT '0',
    emailCode INT DEFAULT '0',
    id SERIAL PRIMARY KEY
);

insert into Userr(nickName, email, password, role, verified, emailCode)
values ('nick', 'email@email.com', 'password', 'user', '0', 0);

select * from Userr;

DROP TABLE IF EXISTS Ticket CASCADE ;
CREATE TABLE IF NOT EXISTS Ticket (
    performanceId INT DEFAULT '0' REFERENCES PerformancesSet(id),
    place INT DEFAULT '0',
    ticketCost FLOAT DEFAULT '10.0',
    userId INT DEFAULT '0',
    id SERIAL PRIMARY KEY
) ;


