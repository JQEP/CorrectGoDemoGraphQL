package database

import (
	"context"
	"fmt"
	model2 "go-graphql-mongodb-api/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type DB struct {
	client *mongo.Client
}

func Connect(dbUrl string) *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) InsertCourseById(course model2.NewCourse) *model2.Course {
	courseColl := db.client.Database("graphql-mongodb-api-db").Collection("course")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := courseColl.InsertOne(ctx, bson.M{"name": course.Name, "subject": course.Subject, "instructorID": course.InstructorID})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnCourse := model2.Course{ID: insertedID, Name: course.Name, Subject: course.Subject, InstructorID: course.InstructorID}

	return &returnCourse
}

func (db *DB) FindCourseById(id string) *model2.Course {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	courseColl := db.client.Database("graphql-mongodb-api-db").Collection("course")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := courseColl.FindOne(ctx, bson.M{"_id": ObjectID})

	course := model2.Course{ID: id}

	res.Decode(&course)

	return &course
}

func (db *DB) AllCourses() []*model2.Course {
	courseColl := db.client.Database("graphql-mongodb-api-db").Collection("course")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := courseColl.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var courses []*model2.Course
	for cur.Next(ctx) {
		sus, err := cur.Current.Elements()
		fmt.Println(sus)
		if err != nil {
			log.Fatal(err)
		}

		course := model2.Course{ID: (sus[0].String()), Name: (sus[1].String()), Subject: (sus[2].String()), InstructorID: (sus[3].String())}

		courses = append(courses, &course)
	}

	return courses
}
