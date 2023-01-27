-- SELECT 'CREATE DATABASE pecunia' 
-- WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'pecunia')\gexec;

-- CREATE USER user_pecunia;

-- ALTER USER user_pecunia WITH ENCRYPTED PASSWORD '1122';

-- GRANT ALL PRIVILEGES ON DATABASE pecunia TO user_pecunia;

-- \connect pecunia;

-- GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA PUBLIC TO user_pecunia;

CREATE TABLE IF NOT EXISTS users (
  id varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL, 
  PRIMARY KEY (id)
);
