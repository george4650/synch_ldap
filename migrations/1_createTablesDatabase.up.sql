CREATE TABLE IF NOT EXISTS users
(
    id character varying PRIMARY KEY,
	surname character varying NOT NULL ,
	givenName character varying NOT NULL,
	createdAt character varying NOT NULL,
	sAMAccountName character varying NOT NULL UNIQUE, 
	telephoneNumber character varying,
	department character varying,
	title character varying,
	city character varying,
    mail character varying  
);

CREATE INDEX idx_user_UUID on users (id);