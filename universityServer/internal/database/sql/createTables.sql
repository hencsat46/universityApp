CREATE TABLE tempUni (

    uni_id SERIAL,
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