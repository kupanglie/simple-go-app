CREATE TABLE `warung_alfin`.`order` (
    `id` BINARY(16) PRIMARY KEY,
    `used_id` BINARY(16),
    `total_price` INT UNSIGNED,
    `invoice` VARCHAR(6),
    `status` INT UNSIGNED,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);

CREATE TABLE `warung_alfin`.`order_detail` (
    `id` BINARY(16) PRIMARY KEY,
    `product_id` BINARY(16),
    `price` INT UNSIGNED,
    `qty` INT UNSIGNED,
    `total_price` INT UNSIGNED,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);

CREATE INDEX invoice_idx ON `warung_alfin`.`order` (invoice);
