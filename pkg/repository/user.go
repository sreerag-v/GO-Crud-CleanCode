package repository

import (
	"Clean/Mongo-Crud/pkg/domain"
	RepoInterface "Clean/Mongo-Crud/pkg/repository/interface"
	"context"
	"fmt"

	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserDatabase struct {
	DB *mongo.Client
}

func NewUserRepository(DB *mongo.Client) RepoInterface.UserRepository {
	return &UserDatabase{DB: DB}
}

func (usr *UserDatabase) CreateUser(ctx context.Context, user domain.Users) error {
	collection := usr.DB.Database("mongo_demo").Collection("users")

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return err
	}
	fmt.Println("id", id)

	return nil
}

func (usr *UserDatabase) GetUserByid(ctx context.Context, Uid string) (domain.UsersResponse, error) {
	collection := usr.DB.Database("mongo_demo").Collection("users")

	var User domain.UsersResponse

	// convert the userId string to mongdb ObjectId
	Oid, err := primitive.ObjectIDFromHex(Uid)
	if err != nil {
		return domain.UsersResponse{}, err
	}

	// Define a filter to find the user by their ID
	filter := bson.M{"_id": Oid}

	err = collection.FindOne(ctx, filter).Decode(&User)
	if err == mongo.ErrNoDocuments {
		return domain.UsersResponse{}, errors.New("user not found")
	}

	if err != nil {
		return domain.UsersResponse{}, errors.New("cant error when find the user")
	}

	return User, nil
}

func (usr *UserDatabase) UpdateUserById(ctx context.Context, Uid string, User domain.Users) error {
	collection := usr.DB.Database("mongo_demo").Collection("users")

	Oid, err := primitive.ObjectIDFromHex(Uid)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": Oid}

	update := bson.M{
		"$set": bson.M{
			"uname":    User.Uname,
			"fname":    User.Fname,
			"lname":    User.Lname,
			"email":    User.Email,
			"password": User.Password,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (usr *UserDatabase) DeleteUserById(ctx context.Context, Uid string) error {
	collection := usr.DB.Database("mongo_demo").Collection("users")

	Oid, err := primitive.ObjectIDFromHex(Uid)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": Oid}

	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
