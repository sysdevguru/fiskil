-- CREATE TYPE severity_type AS ENUM ('debug', 'info', 'warn', 'error', 'fatal');

CREATE TABLE logs (
    id serial,
    service_name VARCHAR(100) NOT NULL,
    payload VARCHAR(2048) NOT NULL,
    severity VARCHAR(10) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE severities (
    id serial,
    service_name VARCHAR(100) NOT NULL,
    severity VARCHAR(10) NOT NULL,
    count INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
