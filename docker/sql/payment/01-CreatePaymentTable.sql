CREATE TABLE `warung_alfin_payment`.`payment` (
    `id` BINARY(16) PRIMARY KEY,
    `invoice` BINARY(16),
    `total_payment` INT UNSIGNED,
    `status` INT UNSIGNED,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NULL,
    `deleted_at` DATETIME NULL
);

CREATE INDEX invoice_idx ON `warung_alfin_payment`.`payment` (invoice);
