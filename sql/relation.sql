CREATE TABLE relation
(
	id 			VARCHAR(64) NOT NULL,
	user_id INT NOT NULL,
	other_user_id INT NOT NULL,
	state VARCHAR (16) NOT NULL,
	PRIMARY KEY (id)
);


CREATE TABLE relation_0
(
	id 			VARCHAR(64) NOT NULL,
	user_id INT NOT NULL,
	other_user_id INT NOT NULL,
	state VARCHAR (16) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE relation_1
(
	id 			VARCHAR(64) NOT NULL,
	user_id INT NOT NULL,
	other_user_id INT NOT NULL,
	state VARCHAR (16) NOT NULL,
	PRIMARY KEY (id)
);


 INSERT INTO relation (id, user_id, other_user_id, state)
 VALUES (11, 11, 11, '11')

