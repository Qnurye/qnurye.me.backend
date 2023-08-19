CREATE TABLE users (
    userId SERIAL PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL unique,
    password VARCHAR(255) NOT NULL,
    website  VARCHAR(255) NOT NULL,
    subscribed BOOLEAN NOT NULL DEFAULT FALSE,
    updateTime TIMESTAMP ON UPDATE now(),
    registerTime TIMESTAMP DEFAULT now()
);
