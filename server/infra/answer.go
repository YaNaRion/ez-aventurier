package infra

import (
	"context"
	"fmt"
	"main/infra/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (db *DB) ModifyAnswer(answerID string, userID string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.Client.Database("dev1").Collection("answers")

	filter := bson.M{
		"answer_id": answerID,
	}

	update := bson.M{
		"$set": bson.M{
			"is_available": false,
		},
	}

	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{
			bson.M{
				"answer_id":    answerID,
				"is_available": true,
				"answer_by":    userID,
			},
		},
	}

	updateOptions := options.UpdateOptions{
		ArrayFilters: &arrayFilters,
	}

	// We need to get the weight before updating or after
	var answer models.Answer
	err := collection.FindOne(ctx, bson.M{"answer_id": answerID}).Decode(&answer)
	if err != nil {
		return 0, fmt.Errorf("failed to decode answer: %w", err)
	}

	// Perform the update
	_, err = collection.UpdateOne(ctx, filter, update, &updateOptions)
	if err != nil {
		return 0, fmt.Errorf("failed to update answer: %w", err)
	}

	return answer.Weight, nil
}
