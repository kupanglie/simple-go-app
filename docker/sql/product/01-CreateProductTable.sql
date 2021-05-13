CREATE TABLE `warung_alfin_product`.`product` (
    `id` BINARY(16) PRIMARY KEY,
    `name` VARCHAR(255),
    `price` INT UNSIGNED,
    `stock` INT UNSIGNED,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);