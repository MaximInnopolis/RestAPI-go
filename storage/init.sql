-- +migrate Up
CREATE TABLE users(
                         id		 SERIAL,
                         name   VARCHAR(50),
                         surname VARCHAR(50),
                         login VARCHAR(50),
                         password VARCHAR(50),
                         birthdate date,
                         status VARCHAR(50),
                         role VARCHAR(50),
                         constraint users_pkey PRIMARY KEY (id)
);

CREATE TABLE keys(
                        id		 SERIAL,
                        user_id integer,
                        key varchar(50),
                        constraint keys_pkey PRIMARY KEY (id),
                        constraint keys_user_id_fkey foreign key (user_id) references users(id)
);

-- +migrate Down
DROP TABLE users;
DROP TABLE keys;