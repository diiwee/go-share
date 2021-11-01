package main

import "fmt"

//IFactory 抽象工厂接口
type IFactory interface {
	User() IUser
}

func NewFactory(name string) IFactory {
	if name == "mysql" {
		return &MysqlFacory{}
	} else if name == "mongo" {
		return &MongoFactory{}
	}
	return nil
}

type IUser interface {
	Create()
	Update()
	Get()
	Delete()
}

//mysql
type mysql struct {
}

func newMysql() *mysql {
	return &mysql{}
}

func (m *mysql) Create() {
	panic("implement me")
}

func (m *mysql) Update() {
	panic("implement me")
}

func (m *mysql) Get() {
	fmt.Println("Mysql get user")
}

func (m *mysql) Delete() {
	panic("implement me")
}

type MysqlFacory struct {
}

func (m *MysqlFacory) User() IUser {
	return newMysql()
}

//mongo
type mongo struct {
}

func newMongo() *mongo {
	return &mongo{}
}

func (m *mongo) Create() {
	panic("implement me")
}

func (m *mongo) Update() {
	panic("implement me")
}

func (m *mongo) Get() {
	fmt.Println("Mongo get user")
}

func (m *mongo) Delete() {
	panic("implement me")
}

type MongoFactory struct {
}

func (m *MongoFactory) User() IUser {
	return newMongo()
}

func main() {
	mysqlFactory := NewFactory("mysql")
	mysqlFactory.User().Get()

	mongoFactory := NewFactory("mongo")
	mongoFactory.User().Get()
}
