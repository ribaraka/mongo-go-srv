package db

import (
	"context"
	"github.com/ribaraka/mongo-go-srv/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoQueries interface {
	SignUp(models.User) (models.User, error)
	Update(string, interface{}) (models.User, error)
	Delete(string) (models.User, error)
	Get(string) (models.User, error)
	Search(interface{}) ([]models.User, error)
}

type DBClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *DBClient) SignUp(user models.User) (models.User, error) {
	todo := models.User{}

	res, err := c.Col.InsertOne(c.Ctx, user)
	if err != nil {
		return todo, err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return c.Get(id)
}

func (c *DBClient) Get(id string) (models.User, error) {
	todo := models.User{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return todo, err
	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&todo)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

//func (c *DBClient) Update(id string, update interface{}) (models.User, error) {
//	result := models.User{
//		ModifiedCount: 0,
//	}
//	_id, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return result, err
//	}
//
//	todo, err := c.Get(id)
//	if err != nil {
//		return result, err
//	}
//	var exist map[string]interface{}
//	b, err := json.Marshal(todo)
//	if err != nil {
//		return result, err
//	}
//	json.Unmarshal(b, &exist)
//
//	change := update.(map[string]interface{})
//	for k := range change {
//		if change[k] == exist[k] {
//			delete(change, k)
//		}
//	}
//
//	if len(change) == 0 {
//		return result, nil
//	}
//
//	res, err := c.Col.UpdateOne(c.Ctx, bson.M{"_id": _id}, bson.M{"$set": change})
//	if err != nil {
//		return result, err
//	}
//
//	newTodo, err := c.Get(id)
//	if err != nil {
//		return result, err
//	}
//
//	result.ModifiedCount = res.ModifiedCount
//	result.Result = newTodo
//	return result, nil
//}
//
//func (c *DBClient) Delete(id string) (models.User, error) {
//	result := models.UserDelete{
//		DeletedCount: 0,
//	}
//	_id, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return result, err
//	}
//
//	res, err := c.Col.DeleteOne(c.Ctx, bson.M{"_id": _id})
//	if err != nil {
//		return result, err
//	}
//	result.DeletedCount = res.DeletedCount
//	return result, nil
//}
//

//
//func (c *DBClient) Search(filter interface{}) ([]models.User, error) {
//	todos := []models.User{}
//	if filter == nil {
//		filter = bson.M{}
//	}
//
//	cursor, err := c.Col.Find(c.Ctx, filter)
//	if err != nil {
//		return todos, err
//	}
//
//	for cursor.Next(c.Ctx) {
//		row := models.User{}
//		cursor.Decode(&row)
//		todos = append(todos, row)
//	}
//
//	return todos, nil
//}

//func (sr *SignUpRepository) AddUser(ctx context.Context, user models.User, conf config.Config) error {
//	tx, err := sr.pool.Begin(ctx)
//	if err != nil {
//		return fmt.Errorf("Unable to commit transaction: %v\n", err)
//	}
//	defer tx.Rollback(ctx)
//
//	insertUser := `INSERT INTO users (firstName, lastName, email) VALUES ($1, $2, $3)`
//	_, err = tx.Exec(ctx, insertUser, user.FirstName, user.LastName, user.Email)
//	if err != nil {
//		return fmt.Errorf("Unable to insert data into database: %v\n", err)
//	}
//
//	hash, err := crypto.HashAndSalt([]byte(user.Password))
//	if err != nil {
//		return fmt.Errorf("failed to hash crypto: %w", err)
//	}
//	insertPasswordHash := `INSERT INTO credentials (password_hash) VALUES ($1)`
//	_, err = tx.Exec(ctx, insertPasswordHash, hash)
//	if err != nil {
//		return fmt.Errorf("Unable to insert hash into password_hash:: %v\n", err)
//	}
//
//	emailToken := crypto.GenerateToken(32)
//	insertEmailToken := `INSERT INTO email_verification_tokens (verification_token) VALUES ($1)`
//	_, err = tx.Exec(ctx, insertEmailToken, emailToken)
//	if err != nil {
//		return fmt.Errorf("Unable to insert token: %v\n", err)
//	}
//
//
//	err = tx.Commit(ctx)
//	if err != nil {
//		return fmt.Errorf("Unable insert user to database %v\n", err)
//
//	}

	//err = email.SendVerifyMassage(conf, user.Email, emailToken)
	//if err != nil {
	//	return fmt.Errorf("unable to send email %w", err)
	//}

//	return nil
//}
//
//func (sr *SignUpRepository) GetByID(ctx context.Context, id int) (*models.TableUser, error) {
//	user := &models.TableUser{}
//	err := sr.pool.QueryRow(ctx,
//		`SELECT * FROM users WHERE id=$1`, id).Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Verified)
//	if err != nil {
//		return nil, err
//	}
//
//	return user, nil
//}
//
//func (sr *SignUpRepository) UpdateUserByEmail(ctx context.Context, user *models.TableUser) error {
//	_, err := sr.pool.Exec(ctx,
//		`UPDATE users SET firstname = $2, lastname = $3, email = $4, verified = $5 WHERE email = $1`,
//		user.Email, user.Firstname, user.Lastname, user.Email, user.Verified)
//	if err != nil {
//		return fmt.Errorf("Unable to update row: %v\n", err)
//	}
//
//	return nil
//}
//
//func (sr *SignUpRepository) GetByEmail(ctx context.Context, email string) (*models.TableUser, error) {
//	user := &models.TableUser{}
//	err := sr.pool.QueryRow(ctx,
//		`SELECT * FROM users WHERE email=$1`, email).Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Verified)
//	if err != nil {
//		return nil, err
//	}
//
//	return user, nil
//}