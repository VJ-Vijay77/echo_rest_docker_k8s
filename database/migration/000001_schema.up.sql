
CREATE TABLE person
(
 "id"        integer NOT NULL,
 firstname text,
 lastname  text,
 email     text,
 password  text,
 age       integer,
 CONSTRAINT PK_1 PRIMARY KEY ( "id" )
);
