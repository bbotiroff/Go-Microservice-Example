CREATE KEYSPACE IF NOT EXISTS employee_data
WITH REPLICATION = { 'class': 'SimpleStrategy', 'replication_factor': '1' }
AND durable_writes = true;

CREATE TABLE IF NOT EXISTS employee_data.employee (
    employee_id UUID,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    gender TEXT,
    PRIMARY KEY (employee_id)
);