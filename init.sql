IF NOT EXISTS (SELECT name FROM sys.databases WHERE name = 'AlbumDb')
BEGIN
    CREATE DATABASE AlbumDb;
END
GO

use AlbumDb;

create table Albums(
    id int IDENTITY PRIMARY KEY,
    title varchar(200) NOT NULL,
    artist varchar(200) NOT NULL,
    price money NOT NULL,    
);
