package repositories

import (
	"database/sql"

	entities "github.com/LeoLion02/album-api/internal/domain/entities"
	"github.com/LeoLion02/album-api/internal/infra/data"
	"github.com/LeoLion02/album-api/internal/shared/log"
)

type IAlbumRepository interface {
	FindAll(baseLog *log.BaseLog) (*[]entities.Album, error)
	FindById(id int64, baseLog *log.BaseLog) (*entities.Album, error)
	Add(album *entities.Album, baseLog *log.BaseLog) (*int64, error)
	Update(album *entities.Album, baseLog *log.BaseLog) error
	Delete(id int64, baseLog *log.BaseLog) error
}

type AlbumRepository struct {
	dbContext data.IDbContext
}

func NewAlbumRepository(dbContext data.IDbContext) IAlbumRepository {
	return &AlbumRepository{dbContext: dbContext}
}

func (repository AlbumRepository) FindAll(baseLog *log.BaseLog) (*[]entities.Album, error) {
	conn, err := repository.dbContext.Connect(baseLog)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	logStep := log.NewLogStep(nil)
	baseLog.AddStep(log.SqlGetAlbums, logStep)
	defer logStep.Finish(&err)

	rows, err := conn.Query("select id, title, artist, price from albums")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var albums []entities.Album
	for rows.Next() {
		var album entities.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return &albums, nil
}

func (repository AlbumRepository) FindById(id int64, baseLog *log.BaseLog) (*entities.Album, error) {
	conn, err := repository.dbContext.Connect(baseLog)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	logStep := log.NewLogStep(id)
	baseLog.AddStep(log.SqlGetAlbumById, logStep)
	defer logStep.Finish(&err)

	row := conn.QueryRow("select id, title, artist, price from albums where id = @id", sql.Named("id", id))

	if err := row.Err(); err != nil {
		return nil, err
	}

	if row == nil {
		return nil, nil
	}

	result := &entities.Album{}

	err = row.Scan(&result.ID, &result.Title, &result.Artist, &result.Price)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repository AlbumRepository) Add(album *entities.Album, baseLog *log.BaseLog) (*int64, error) {
	conn, err := repository.dbContext.Connect(baseLog)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	logStep := log.NewLogStep(album)
	baseLog.AddStep(log.SqlAddAlbum, logStep)
	defer logStep.Finish(&err)

	sqlRow := conn.QueryRow(`
			insert into albums (title, artist, price) values (@title, @artist, @price);
			select SCOPE_IDENTITY();
		`,
		sql.Named("title", album.Title),
		sql.Named("artist", album.Artist),
		sql.Named("price", album.Price))

	if err := sqlRow.Err(); err != nil {
		return nil, err
	}

	var lastInsertId int64
	if err := sqlRow.Scan(&lastInsertId); err != nil {
		return nil, err
	}

	return &lastInsertId, nil
}

func (repository AlbumRepository) Update(album *entities.Album, baseLog *log.BaseLog) error {
	conn, err := repository.dbContext.Connect(baseLog)
	if err != nil {
		return err
	}

	logStep := log.NewLogStep(album)
	baseLog.AddStep(log.SqlUpdateAlbum, logStep)

	_, err = conn.Exec("UPDATE albums SET Title = @title, Artist = @artist, Price = @price where Id = @Id",
		sql.Named("title", album.Title),
		sql.Named("artist", album.Artist),
		sql.Named("price", album.Price),
		sql.Named("Id", album.ID))

	logStep.Finish(&err)
	return err
}

func (repository AlbumRepository) Delete(id int64, baseLog *log.BaseLog) error {

	conn, err := repository.dbContext.Connect(baseLog)

	if err != nil {
		return err
	}

	defer conn.Close()

	logStep := log.NewLogStep(id)
	baseLog.AddStep(log.SqlDeleteAlbum, logStep)

	_, err = conn.Exec("delete from albums where id = @id;", sql.Named("id", id))

	logStep.Finish(&err)

	return err
}
