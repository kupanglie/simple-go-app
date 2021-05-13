CREATE TABLE `warung_alfin_user`.`user` (
    `id` BINARY(16) PRIMARY KEY,
    `name` VARCHAR(255),
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);