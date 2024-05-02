CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(15) NOT NULL,
    username VARCHAR(15) NOT NULL,
    password VARCHAR(255) NOT NULL,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    description TEXT
);

CREATE TABLE IF NOT EXISTS Chats (
    id SERIAL PRIMARY KEY,
    first_user_id INTEGER NOT NULL,
    second_user_id INTEGER NOT NULL,
    FOREIGN KEY (first_user_id) REFERENCES Users(id),
    FOREIGN KEY (second_user_id) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Messages (
    id SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL,
    receiver_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    chat_id INTEGER NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES Users(id),
    FOREIGN KEY (receiver_id) REFERENCES Users(id),
    FOREIGN KEY (chat_id) REFERENCES Chats(id)
);
