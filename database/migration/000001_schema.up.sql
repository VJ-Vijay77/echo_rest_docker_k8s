
CREATE TABLE person
(
 id serial primary key,   
 firstname text,
 lastname  text,
 email     text unique,
 password  text,
 age       integer
);
