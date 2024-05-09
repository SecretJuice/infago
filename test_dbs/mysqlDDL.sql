CREATE TABLE IF NOT EXISTS students
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    fav_color VARCHAR(20),
    age INT
);

INSERT INTO students(name, fav_color, age) VALUES ('Blake McGhie', 'pthalo green', 22);
INSERT INTO students(name, fav_color, age) VALUES ('Connor Robb', 'idk blue ig', 23);