CREATE TABLE IF NOT EXISTS classes (
    class_id INT AUTO_INCREMENT PRIMARY KEY,
    class_name VARCHAR(50) NOT NULL,
    student_id INT NOT NULL
);

INSERT INTO classes (class_name, student_id) VALUES
('Math', 1),
('Science', 2),
('History', 1),
('English', 2),
('Physics', 1);