CREATE TABLE IF NOT EXISTS users (
    "user_id"         SERIAL PRIMARY KEY,
    "username"        VARCHAR(20) UNIQUE, 
    "email"           VARCHAR(255) UNIQUE NOT NULL,
    "password_hash"   VARCHAR(255) NOT NULL, 
    "display_name"    VARCHAR(50),
    "self_introduction" TEXT,
    "location"        VARCHAR(100), 
    "website"         VARCHAR(255),
    "date_of_birth"   DATE, 
    "profile_image"   VARCHAR(255),
    "header_image"    VARCHAR(255),
    "is_active"       BOOLEAN NOT NULL DEFAULT false,
    "created_at"      TIMESTAMP NOT NULL DEFAULT current_timestamp,
    "updated_at"      TIMESTAMP NOT NULL DEFAULT current_timestamp,
    "deleted_at"      TIMESTAMP 
);