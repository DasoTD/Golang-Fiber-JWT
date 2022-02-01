package data

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

func CreateDBEngine() (*xorm.Engine, error) {
	connectioninfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "root", "password", "authserver")
	engine, err := xorm.NewEngine("postgres", connectioninfo)
	if err != nil {
		return nil, err
	}
	if err := engine.Ping(); err != nil {
		return nil, err
	}
	if err := engine.Sync(new(User)); err != nil {
		return nil, err
	}
	return engine, nil
}

//https://github.com/learning-zone/nodejs-interview-questions#
//https://www.interviewbit.com/node-js-interview-questions/#difference-between-process-nexttick-and-setimmediate-methods
//https://www.edureka.co/blog/interview-questions/top-node-js-interview-questions-2016/
