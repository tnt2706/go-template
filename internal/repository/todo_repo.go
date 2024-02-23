package repository

import (
	"context"
	"lineapp/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepo interface {
	CreateToDo(*model.Todo) (*model.Todo, error)
	GetTodoById(string) (*model.Todo, error)
}

type todoRepoImpl struct {
	DB *mongo.Collection
}

func NewTodoRepo(DB *mongo.Collection) TodoRepo {
	return &todoRepoImpl{DB: DB}
}

func (u *todoRepoImpl) CreateToDo(todo *model.Todo) (*model.Todo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	todo.Add()
	result, err := u.DB.InsertOne(ctx, todo)

	if err != nil {
		return nil, err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)
	todo.ID = oid.Hex()

	return todo, err
}

func (u *todoRepoImpl) GetTodoById(id string) (*model.Todo, error) {

	var todo *model.Todo
	todoObjId, _ := primitive.ObjectIDFromHex(id)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := u.DB.FindOne(ctx, bson.M{"_id": todoObjId}).Decode(&todo)
	return todo, err
}
