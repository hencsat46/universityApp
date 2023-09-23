CREATE TABLE tempUni (

    uni_id SERIAL PRIMARY KEY,
    uni_name TEXT,
    uni_des TEXT,
    uni_img TEXT,

);

CREATE TABLE users (

    user_id SERIAL,
    username TEXT,
    passwd TEXT

);

CREATE TABLE jwtKey (

    secretkey TEXT

);

CREATE TABLE students_records (

	record_id SERIAL PRIMARY KEY,
    student_id INT,
	student_university INT,
	student_points INT,
    FOREIGN KEY (student_id) REFERENCES users(user_id),
	FOREIGN KEY (student_university) REFERENCES tempUni(uni_id)

);