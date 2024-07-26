CREATE TABLE if not exists Users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    pid VARCHAR(191) NOT NULL UNIQUE default uuid() ,
    email VARCHAR(191) NOT NULL UNIQUE,
    password VARCHAR(191) NOT NULL,
    role VARCHAR(191) NOT NULL,
    createdAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);