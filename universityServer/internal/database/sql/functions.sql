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

CREATE OR REPLACE FUNCTION add_record(login_name TEXT, university TEXT, points INT)
RETURNS INT
LANGUAGE plpgsql AS $$
DECLARE
	st_id INT;
	university_id INT;
BEGIN
	st_id = (SELECT user_id FROM users WHERE username = login_name);
	university_id = (SELECT uni_id FROM tempUni WHERE uni_name = university);
	INSERT INTO students_records(student_id, student_university, student_points) VALUES (st_id, university_id, points);
	RETURN 0;
END; $$