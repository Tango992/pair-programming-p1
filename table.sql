CREATE TABLE games (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    description VARCHAR(255),
    published DATE
);

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE user_games (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    game_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (game_id) REFERENCES games(id)
);

INSERT INTO games (name, description, published)
VALUES 
('dota2', 'RPG', '2013-07-09'),
('valorant', 'FPS', '2020-06-02'),
('cs:go', 'FPS', '2012-08-21'),
('battle realms', 'RPG', '2001-11-07');

INSERT INTO users (username, password)
VALUES
('admin', 'admin');

INSERT INTO user_games (user_id, game_id)
VALUES
(1, 1),
(1, 2);



