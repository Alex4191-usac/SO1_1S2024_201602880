
CREATE DATABASE IF NOT EXISTS sample_db;

USE sample_db;

CREATE TABLE IF NOT EXISTS example_table (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

INSERT INTO example_table (name) VALUES ('Example Data 1');
INSERT INTO example_table (name) VALUES ('Example Data 2');

CREATE TABLE IF NOT EXISTS ram_module (
  id INT AUTO_INCREMENT PRIMARY KEY,
  total_memory FLOAT,
  used_memory FLOAT,
  free_memory FLOAT,
  percentage_used FLOAT,
  created_at TIMESTAMP 
);

CREATE TABLE IF NOT EXISTS cpu_module (
  id INT AUTO_INCREMENT PRIMARY KEY,
  percentage_used FLOAT,
  created_at TIMESTAMP 
);
