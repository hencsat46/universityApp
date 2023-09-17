CREATE OR REPLACE FUNCTION login(uname TEXT, pass TEXT)
RETURNS INT
LANGUAGE plpgsql AS $$
DECLARE
	response INT;
BEGIN
	response = (SELECT COUNT(*) FROM users WHERE username = uname AND passwd = pass);
	IF response = 1 THEN
		RETURN 0;
	END IF;
	RETURN -1;
END; $$

CREATE OR REPLACE FUNCTION checkUser(uname TEXT)
RETURNS INT
LANGUAGE plpgsql AS $$
DECLARE
	response INT;
BEGIN
	response = (SELECT COUNT(*) FROM users WHERE username = uname);
	IF response = 1 THEN
		RETURN 0;
	END IF;
	RETURN -1;
END; $$