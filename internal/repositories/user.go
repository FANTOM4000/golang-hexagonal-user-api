package repositories

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	mc  *mongo.Client
	db  string
	col string
}

func NewUserRepository(mc *mongo.Client, db string) ports.UserRepository {
	collection :=  mc.Database(db).Collection("user")
	indexName, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal("Create index error")
	}
	fmt.Printf("Created index %v\n", indexName)

	return userRepository{mc: mc, db: db, col: "user"}
}

func (ur userRepository) CreateUser(ctx context.Context, user domains.UserProperty) (domains.User, error) {
	userObj := domains.User{}

	collection := ur.mc.Database(ur.db).Collection(ur.col)
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return domains.User{}, err
	}
	id, _ := res.InsertedID.(primitive.ObjectID)
	
	userObj.ID = id
	userObj.Username = user.Username 
	userObj.Password = user.Password
	userObj.Mobile = user.Mobile 
	userObj.Role = user.Role 
	userObj.CreatedAt = user.CreatedAt  
	userObj.UpdatedAt = user.UpdatedAt   
	
	return userObj, nil
}
func (ur userRepository) UpdateUserById(ctx context.Context, id string, user domains.UserMetadata) error {

	collection := ur.mc.Database(ur.db).Collection(ur.col)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.D{
		{
			Key:"$set", Value:bson.D{
				{
					Key:"mobile", Value:user.Mobile,
				},
				{
					Key:"role", Value:user.Role,
				},
			},
		},
	}
	fillter := bson.D{{Key: "_id", Value: idObj}}
	if err := collection.FindOneAndUpdate(ctx, fillter, update).Decode(&user); err != nil {
		return err
	}
	return nil
}
func (ur userRepository) DeleteUserById(ctx context.Context, id string) error {
	collection := ur.mc.Database(ur.db).Collection(ur.col)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: idObj}}, nil); 
	if err != nil {
		return err
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return nil
}
func (ur userRepository) GetUserById(ctx context.Context, id string) (domains.User, error) {
	user := domains.User{}
	collection := ur.mc.Database(ur.db).Collection(ur.col)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domains.User{}, err
	}
	if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: idObj}}, nil).Decode(&user); err != nil {
		return domains.User{}, err
	}
	return user, nil
}
func (ur userRepository) GetUserByUsername(ctx context.Context, username string) (domains.User, error) {
	user := domains.User{}
	collection := ur.mc.Database(ur.db).Collection(ur.col)

	if err := collection.FindOne(ctx, bson.D{{Key: "username", Value: username}}, nil).Decode(&user); err != nil {
		return domains.User{}, err
	}
	return user, nil
}
