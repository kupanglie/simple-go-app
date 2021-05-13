INSERT INTO `warung_alfin`.`user`
VALUES (
        UUID_TO_BIN(UUID()),
        "GOD",
        NOW(),
        NULL,
        NULL
    );
INSERT INTO `warung_alfin`.`user`
VALUES (
        UUID_TO_BIN(UUID()),
        "Anak Warung 1",
        NOW(),
        NULL,
        NULL
    );
INSERT INTO `warung_alfin`.`product`
VALUES (
        UUID_TO_BIN(UUID()),
        "KOPI CAP TIKUS",
        1000,
        50,
        NOW(),
        NULL,
        NULL
    );
INSERT INTO `warung_alfin`.`product`
VALUES (
        UUID_TO_BIN(UUID()),
        "KOPI CAP BUAYA",
        100,
        999,
        NOW(),
        NULL,
        NULL
    );