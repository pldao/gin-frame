package model

import (
	"context"
	"github.com/PLDao/gin-frame/data"
	"github.com/PLDao/gin-frame/internal/validator/form"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Social struct {
	UserName string         `json:"username" bson:"username"`
	Socials  []*SocialModel `json:"socials" bson:"socials"`
	*mongo.Collection
}

type SocialModel struct {
	SocialName string `json:"social_name" bson:"social_name"`
	SocialLink string `json:"social_link" bson:"social_link"`
}

func NewSocial() *Social {
	db := data.MongoDB.Database("social").Collection("user_social")
	return &Social{
		UserName:   "",
		Socials:    []*SocialModel{},
		Collection: db,
	}
}

func (s *Social) CollectionName() string {
	return "social"
}

func (s *Social) AddSocial(param *form.UserSocial) error {
	filter := bson.M{"username": param.UserName}
	update := bson.M{
		"$push": bson.M{
			"socials": SocialModel{param.SocialName, param.SocialLink},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.Collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}

func (s *Social) ListSocial(userName string) ([]*SocialModel, error) {
	filter := bson.M{"username": userName}
	var social *Social
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.Collection.FindOne(ctx, filter).Decode(&social)
	if err != nil {
		return nil, err
	}
	return social.Socials, nil
}

//func (s *Social) DelSocial(userName string, socialName string) error {
//	collection := s.Connect()
//
//	filter := bson.M{"username": userName}
//	update := bson.M{
//		"$pull": bson.M{
//			"socials": bson.M{"social_name": socialName},
//		},
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	_, err := collection.UpdateOne(ctx, filter, update)
//	return err
//}
//
