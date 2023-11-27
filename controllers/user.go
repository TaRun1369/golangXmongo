package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github/TaRun1369/golangXmongo/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
    client *mongo.Client
}

func NewUserController(c *mongo.Client) *UserController {
    return &UserController{c}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    oid, _ := primitive.ObjectIDFromHex(id)
    u := models.User{}

    collection := uc.client.Database("golangXmongo").Collection("users")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    err := collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&u)

    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    uj, _ := json.Marshal(u)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    u := models.User{}
    json.NewDecoder(r.Body).Decode(&u)
    u.Id = primitive.NewObjectID()

    collection := uc.client.Database("golangXmongo").Collection("users")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    result, _ := collection.InsertOne(ctx, u)

    uj, _ := json.Marshal(result)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    oid, _ := primitive.ObjectIDFromHex(id)

    collection := uc.client.Database("golangXmongo").Collection("users")
    ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
    result, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Deleted user %s\n", id)
}