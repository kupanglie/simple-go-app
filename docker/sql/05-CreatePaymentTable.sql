CREATE TABLE `warung_alfin`.`payment` (
    `id` BINARY(16) PRIMARY KEY,
    `invoice` BINARY(16),
    `status` TINYINT(1) UNSIGNED,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);