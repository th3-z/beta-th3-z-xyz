package storage

var schema = `
    CREATE TABLE IF NOT EXISTS person (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        username VARCHAR NOT NULL,
        password VARCHAR NOT NULL,
        insert_date INTEGER NOT NULL,
        deleted_date INTEGER DEFAULT 0
    );

    CREATE TABLE IF NOT EXISTS signin_attempt (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        ip_address VARCHAR NOT NULL,
        insert_date INTEGER NOT NULL
    );

    CREATE TABLE IF NOT EXISTS server (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL,
        address VARCHAR NOT NULL,
        image VARCHAR NOT NULL,
        service VARCHAR NOT NULL,
        web_url VARCHAR NOT NULL,
        locked INT NOT NULL DEFAULT 0,
        max_players INT NOT NULL DEFAULT 0
    );

    CREATE TABLE IF NOT EXISTS infrastructure (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        hostname VARCHAR NOT NULL,
        address VARCHAR NOT NULL,
        os VARCHAR NOT NULL
    );

    CREATE TABLE IF NOT EXISTS software (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL,
        project_url VARCHAR NOT NULL,
        repo_url VARCHAR NOT NULL,
        description VARCHAR NOT NULL,
        status VARCHAR NOT NULL
    );

    CREATE TABLE IF NOT EXISTS paste (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        uploader_id VARCHAR NOT NULL,  -- Hash of IP address
        filename VARCHAR NOT NULL UNIQUE,
        insert_date INTEGER NOT NULL
    );
`
