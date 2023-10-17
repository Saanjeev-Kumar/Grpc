package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"goapp/internal/mongodb"
	schema "goapp/internal/mongodb/schema"

	user "goapp/proto/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser creates a User and stores in mongodb
func (s *Server) createUser(ctx context.Context, mongoClient *mongo.Client, req *user.CreateUserRequest) error {
	userInfo := &schema.User{}
	fmt.Println("CreateUser request in mongodb", userInfo)
	userInfo.ConvertToSchema(req.GetUser())
	fmt.Println("Convertion to schema")
	fmt.Println("s.database", s.database)
	fmt.Println("userInfo", userInfo)
	_, err := mongodb.InsertOne(ctx, s.database, "admin1", "Users", userInfo)
	// handle the error
	if err != nil {
		return err
	}
	return nil
}

// GetUser fetches User details from mongodb
func (s *Server) getUser(ctx context.Context, mongoClient *mongo.Client, userName string) (*user.User, error) {
	res, err := mongodb.FindOne(ctx, s.database, "admin1", "Users", bson.M{"name": userName})
	fmt.Println("get request", res)
	fmt.Println("name", userName)
	fmt.Println("bson.M name", bson.M{"name": userName})
	if err != nil {
		return nil, err
	}
	var userInfo *schema.User
	fmt.Println("userInfo", userInfo)
	if err = res.(*mongo.SingleResult).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo.ConvertToProto(), nil
}

// UpdateUser updates existing user in mongodb
func (s *Server) updateUser(ctx context.Context, mongoClient *mongo.Client, req *user.UpdateUserRequest) error {
	var updateParams map[string]interface{}
	userBytes, _ := json.Marshal(req.GetUser())
	json.Unmarshal(userBytes, &updateParams)
	filter := bson.M{"name": req.GetName()}
	fields := bson.M{"$set": updateParams}
	fmt.Println(fields)
	fmt.Println(filter)
	mongodb.UpdateOne(ctx, s.database, "admin1", "Users", filter, fields)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// DeleteUser deleltes the user from mongodb
func (s *Server) deleteUser(ctx context.Context, mongoClient *mongo.Client, userName string) (string, error) {
	err := mongodb.DeleteOne(ctx, s.database, "admin1", "Users", bson.M{"name": userName})
	if err != nil {
		return "", err
	}
	return "Deleted record successfully", nil
}
