CREATE TABLE customer_acc (
    customer_id INT AUTO_INCREMENT,
    email VARCHAR(120) UNIQUE NOT NULL,
    username VARCHAR(120) UNIQUE NOT NULL,
    password CHAR(60) NOT NULL,
    image_path TEXT NOT NULL,
    balance INT NOT NULL DEFAULT 0,
    date_registered DATE NOT NULL,
    PRIMARY KEY(customer_id)
);

CREATE TABLE customer_req (
    customer_id INT AUTO_INCREMENT,
    date_requested DATE NOT NULL,
    req_type ENUM('ADD', 'SUBTRACT') NOT NULL,
    status ENUM('SUCCESS', 'FAILED', 'PENDING') NOT NULL,
    amount INT NOT NULL DEFAULT 0,
    PRIMARY KEY(customer_id),
    FOREIGN KEY(customer_id) REFERENCES customer_acc(customer_id)
);

CREATE TABLE customer_transfer (
    customer_id INT AUTO_INCREMENT,
    destination INT NOT NULL,
    date_transfer DATE NOT NULL,
    amount INT NOT NULL DEFAULT 0,
    PRIMARY KEY(customer_id),
    FOREIGN KEY(customer_id) REFERENCES customer_acc(customer_id)
);