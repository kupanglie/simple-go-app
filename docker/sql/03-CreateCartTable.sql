CREATE TABLE `warung_alfin`.`cart` (
    `id` BINARY(16) PRIMARY KEY,
    `used_id` BINARY(16),
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);

CREATE TABLE `warung_alfin`.`cart_detail` (
    `id` BINARY(16) PRIMARY KEY,
    `product_id` BINARY(16),
    `price` INT UNSIGNED,
    `qty` INT UNSIGNED,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);