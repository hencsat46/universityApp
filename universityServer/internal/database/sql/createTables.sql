CREATE TABLE tempUni (

    uni_id SERIAL PRIMARY KEY,
    uni_name TEXT,
    uni_des TEXT,
    uni_img TEXT,
    min_point INT,
    seats_count INT

);

CREATE TABLE users (

    user_id SERIAL PRIMARY KEY,
    username TEXT,
    passwd TEXT,
    student_name TEXT,
    student_surname TEXT

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