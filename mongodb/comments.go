package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/katsuokaisao/mongodb-play/domain"
	"github.com/katsuokaisao/mongodb-play/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type commentRepository struct {
	cli      *mongo.Client
	dbName   string
	collName string
}

func NewCommentRepository(cli *mongo.Client) repository.CommentRepository {
	return &commentRepository{
		cli:      cli,
		dbName:   "sample_mflix",
		collName: "comments",
	}
}

func (c *commentRepository) coll() *mongo.Collection {
	return c.cli.Database(c.dbName).Collection(c.collName)
}

func (c *commentRepository) FindOneByID(id string) (*domain.Comment, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id: %w", err)
	}
	filter := bson.M{"_id": bson.M{"$eq": _id}}

	var comment *repository.Comment
	if err := c.coll().FindOne(context.TODO(), filter).Decode(&comment); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find one comment: %w", err)
	}
	com := comment.ToDomain()
	return &com, nil
}

func (c *commentRepository) FindByName(name string) ([]domain.Comment, error) {
	filter := bson.M{"name": bson.M{"$eq": name}}

	return c.find(filter)
}

func (c *commentRepository) FindByMovieID(movieID string) ([]domain.Comment, error) {
	_movieID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert movieID: %w", err)
	}
	filter := bson.M{"movie_id": bson.M{"$eq": _movieID}}

	return c.find(filter)
}

func (c *commentRepository) FindByDateRange(start, end time.Time) ([]domain.Comment, error) {
	filter := bson.M{
		"date": bson.M{
			"$gte": start,
			"$lt":  end,
		},
	}

	return c.find(filter)
}

func (c *commentRepository) Find(cond repository.FindCondition) ([]domain.Comment, error) {
	filter := c.convertFilter(cond)
	return c.find(filter)
}

func (c *commentRepository) find(filter bson.M) ([]domain.Comment, error) {
	cursor, err := c.coll().Find(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find comments: %w", err)
	}
	defer cursor.Close(context.Background())

	var comments []domain.Comment
	for cursor.Next(context.Background()) {
		var comment repository.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, fmt.Errorf("failed to decode comment: %w", err)
		}
		comments = append(comments, comment.ToDomain())
	}
	return comments, nil
}

func (c *commentRepository) InsertOne(comment domain.Comment) (string, error) {
	com, err := repository.FromDomain(&comment)
	if err != nil {
		return "", fmt.Errorf("failed to convert domain to repository: %w", err)
	}

	res, err := c.coll().InsertOne(context.TODO(), com)
	if err != nil {
		return "", fmt.Errorf("error in insert one %v", err)
	}
	insertID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert inserted id: %w", err)
	}

	return insertID.String(), nil
}

func (c *commentRepository) InsertMany(comments []domain.Comment) ([]string, error) {
	var docs []interface{}
	for _, comment := range comments {
		com, err := repository.FromDomain(&comment)
		if err != nil {
			return nil, fmt.Errorf("failed to convert domain to repository: %w", err)
		}
		docs = append(docs, com)
	}

	res, err := c.coll().InsertMany(context.TODO(), docs)
	if err != nil {
		return nil, fmt.Errorf("error in insert many %v", err)
	}

	var ids []string
	for _, id := range res.InsertedIDs {
		insertID, ok := id.(primitive.ObjectID)
		if !ok {
			return nil, fmt.Errorf("failed to convert inserted id: %w", err)
		}
		ids = append(ids, insertID.String())
	}
	return ids, nil
}

func (c *commentRepository) UpdateOne(id string, filed repository.UpdateFiled) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert id: %w", err)
	}
	fil := bson.M{"_id": bson.M{"$eq": _id}}

	update := bson.M{
		"$set": bson.M{},
	}
	if filed.Name != nil {
		update["$set"].(bson.M)["name"] = *filed.Name
	}
	if filed.Email != nil {
		update["$set"].(bson.M)["email"] = *filed.Email
	}
	if filed.MovieID != nil {
		_movieID, err := primitive.ObjectIDFromHex(*filed.MovieID)
		if err != nil {
			return fmt.Errorf("failed to convert movieID: %w", err)
		}
		update["$set"].(bson.M)["movie_id"] = _movieID
	}
	if filed.Text != nil {
		update["$set"].(bson.M)["text"] = *filed.Text
	}
	if filed.Date != nil {
		update["$set"].(bson.M)["date"] = *filed.Date
	}

	out, err := c.coll().UpdateOne(context.TODO(), fil, update)
	if err != nil {
		return fmt.Errorf("error in update one %v", err)
	}
	if out.MatchedCount == 0 {
		return fmt.Errorf("no matched document")
	}
	fmt.Printf("matched count: %d, modified count: %d, upserted count: %d\n", out.MatchedCount, out.ModifiedCount, out.UpsertedCount)

	return nil
}

func (c *commentRepository) UpdateMany(cond repository.FindCondition, filed repository.UpdateFiled) error {
	fil := c.convertFilter(cond)

	update := bson.M{
		"$set": bson.M{},
	}
	if filed.Name != nil {
		update["$set"].(bson.M)["name"] = *filed.Name
	}
	if filed.Email != nil {
		update["$set"].(bson.M)["email"] = *filed.Email
	}
	if filed.MovieID != nil {
		_movieID, err := primitive.ObjectIDFromHex(*filed.MovieID)
		if err != nil {
			return fmt.Errorf("failed to convert movieID: %w", err)
		}
		update["$set"].(bson.M)["movie_id"] = _movieID
	}
	if filed.Text != nil {
		update["$set"].(bson.M)["text"] = *filed.Text
	}
	if filed.Date != nil {
		update["$set"].(bson.M)["date"] = *filed.Date
	}

	out, err := c.coll().UpdateMany(context.TODO(), fil, update)
	if err != nil {
		return fmt.Errorf("error in update many %v", err)
	}
	if out.MatchedCount == 0 {
		return fmt.Errorf("no matched document")
	}
	fmt.Printf("matched count: %d, modified count: %d, upserted count: %d\n", out.MatchedCount, out.ModifiedCount, out.UpsertedCount)

	return nil
}

func (c *commentRepository) DeleteOne(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert id: %w", err)
	}
	fil := bson.M{"_id": bson.M{"$eq": _id}}

	out, err := c.coll().DeleteOne(context.TODO(), fil)
	if err != nil {
		return fmt.Errorf("error in delete one %v", err)
	}
	if out.DeletedCount == 0 {
		return fmt.Errorf("no matched document")
	}
	fmt.Printf("deleted count: %d\n", out.DeletedCount)
	return nil
}

func (c *commentRepository) DeleteMany(cond repository.FindCondition) error {
	fil := c.convertFilter(cond)

	out, err := c.coll().DeleteMany(context.TODO(), fil)
	if err != nil {
		return fmt.Errorf("error in delete many %v", err)
	}
	if out.DeletedCount == 0 {
		return fmt.Errorf("no matched document")
	}
	fmt.Printf("deleted count: %d\n", out.DeletedCount)
	return nil
}

func (c *commentRepository) ReplaceOne(id string, comment domain.Comment) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert id: %w", err)
	}
	fil := bson.M{"_id": bson.M{"$eq": _id}}

	com, err := repository.FromDomain(&comment)
	if err != nil {
		return fmt.Errorf("failed to convert domain to repository: %w", err)
	}

	out, err := c.coll().ReplaceOne(context.TODO(), fil, com)
	if err != nil {
		return fmt.Errorf("error in replace one %v", err)
	}
	if out.MatchedCount == 0 {
		return fmt.Errorf("no matched document")
	}
	fmt.Printf("matched count: %d, modified count: %d, upserted count: %d\n", out.MatchedCount, out.ModifiedCount, out.UpsertedCount)

	return nil
}


func (c *commentRepository) convertFilter(findCondition repository.FindCondition) bson.M {
	fil := bson.M{}
	if findCondition.ID != nil {
		_id, err := primitive.ObjectIDFromHex(*findCondition.ID)
		if err != nil {
			return nil
		}
		fil["_id"] = bson.M{"$eq": _id}
	}
	if findCondition.Name != nil {
		fil["name"] = bson.M{"$eq": *findCondition.Name}
	}
	if findCondition.MovieID != nil {
		_movieID, err := primitive.ObjectIDFromHex(*findCondition.MovieID)
		if err != nil {
			return nil
		}
		fil["movie_id"] = bson.M{"$eq": _movieID}
	}
	if findCondition.Start != nil && findCondition.End != nil {
		fil["date"] = bson.M{
			"$gte": *findCondition.Start,
			"$lt":  *findCondition.End,
		}
	}
	return fil
}
