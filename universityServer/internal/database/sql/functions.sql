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
	IF response > 0 THEN
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

CREATE OR REPLACE FUNCTION get_records()
RETURNS TABLE("Имя" TEXT, "Фамилия" TEXT, "Вуз" TEXT, "Баллы" INT)
LANGUAGE plpgsql AS $$
BEGIN
	RETURN query (SELECT u.student_name, u.student_surname, t.uni_name, sr.student_points FROM users u RIGHT JOIN students_records sr ON sr.student_id = u.user_id LEFT JOIN tempuni t ON t.uni_id = sr.student_university);
END; $$;

CREATE OR REPLACE FUNCTION get_user_data(u_name TEXT)
RETURNS TABLE("Имя" TEXT, "Фамилия" TEXT, "Логин" TEXT, "Университет" TEXT)
LANGUAGE plpgsql AS $$
BEGIN
	RETURN QUERY SELECT u.username, u.student_name, u.student_surname, t.uni_name FROM users u LEFT JOIN students_records sr ON sr.student_id = u.user_id LEFT JOIN tempuni t ON t.uni_id = sr.student_university WHERE u.username = u_name;
END; $$;

CREATE OR REPLACE FUNCTION update_record(login_name TEXT, university TEXT, points INT)
RETURNS INT
LANGUAGE plpgsql AS $$
DECLARE
	st_id INT;
	university_id INT;
BEGIN
	st_id = (SELECT user_id FROM users WHERE username = login_name);
	university_id = (SELECT uni_id FROM tempUni WHERE uni_name = university);
	UPDATE students_records SET student_university = university_id WHERE student_id = st_id;
	RETURN 0;
END; $$;