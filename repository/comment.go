package repository

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/mongodb-play/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository interface {
	FindOneByID(id string) (*domain.Comment, error)
	FindByName(name string) ([]domain.Comment, error)
	FindByMovieID(movieID string) ([]domain.Comment, error)
	FindByDateRange(start, end time.Time) ([]domain.Comment, error)
	Find(query FindCondition) ([]domain.Comment, error)
	InsertOne(comment domain.Comment) (string, error)
	InsertMany(comments []domain.Comment) ([]string, error)
	UpdateOne(id string, filed UpdateFiled) error
	UpdateMany(cond FindCondition, filed UpdateFiled) error
	DeleteOne(id string) error
	DeleteMany(cond FindCondition) error
	ReplaceOne(id string, comment domain.Comment) error
	EstimatedDocumentCount() (int64, error)
}

type FindCondition struct {
	ID      *string
	Name    *string
	MovieID *string
	Start   *time.Time
	End     *time.Time
}

type Comment struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Email   string             `bson:"email"`
	MovieID primitive.ObjectID `bson:"movie_id"`
	Text    string             `bson:"text"`
	Date    time.Time          `bson:"date"`
}

type UpdateFiled struct {
	Name    *string
	Email   *string
	MovieID *string
	Text    *string
	Date    *time.Time
}

func (c *Comment) ToDomain() domain.Comment {
	return domain.Comment{
		ID:      c.ID.String(),
		Name:    c.Name,
		Email:   c.Email,
		MovieID: c.MovieID.String(),
		Text:    c.Text,
		Date:    c.Date,
	}
}

func FromDomain(comment *domain.Comment) (Comment, error) {
	movieID, err := primitive.ObjectIDFromHex(comment.MovieID)
	if err != nil {
		return Comment{}, err
	}

	id := primitive.NewObjectID()
	if comment.ID != "" {
		id, err = primitive.ObjectIDFromHex(comment.ID)
		if err != nil {
			return Comment{}, fmt.Errorf("failed to convert ID: %w", err)
		}
	}

	return Comment{
		ID:      id,
		Name:    comment.Name,
		Email:   comment.Email,
		MovieID: movieID,
		Text:    comment.Text,
		Date:    comment.Date,
	}, nil
}
