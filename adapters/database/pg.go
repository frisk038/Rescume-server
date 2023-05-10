package database

import (
	"context"
	"fmt"
	"os"
	"rescueme-server/business/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	conn *pgxpool.Pool
}

const (
	insertCoordinates = `INSERT INTO user_locations (user_id, location)
		VALUES ($1, ST_GeographyFromText($2));`
	selectNearbyUsers = `SELECT user_id FROM user_locations
		WHERE ST_DWithin(location, ST_GeographyFromText($1), 1000) AND user <> $3; -- 1000 meters`
)

func NewPGStore() (*Store, error) {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &Store{conn: conn}, nil
}

func (s *Store) InsertUserCoordinates(ctx context.Context, userID string, coordinate models.Coordinates) error {
	point := fmt.Sprintf("POINT(%s %s)", coordinate.Longitude, coordinate.Latitiude)
	_, err := s.conn.Exec(ctx, insertCoordinates, userID, point)
	if err != nil {
		return err
	}
	return nil
}
