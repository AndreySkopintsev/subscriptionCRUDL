package db

import (
	"common"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	SubscriptionTableName = "subscriptions"
)

// TODO do all the methods

// CRUD db methods
// Create
func CreateNewRecord(price int, serviceName string) (common.Subscription, error) {
	// TODO ADD start date
	userId := uuid.New()
	// NOTE TO SELF: dont forget 'single quotes' around values (except for int)
	psqlQuery := fmt.Sprintf(`INSERT INTO %s (user_id, price, service_name, start_date) VALUES ('%s', %s, '%s', '%s')`, SubscriptionTableName, userId, strconv.Itoa(price), serviceName, time.Now().Format("2006-01-02"))
	fmt.Println("resulting query string is: ", psqlQuery)
	_, err := PostgresDB.Exec(psqlQuery)
	if err != nil {
		fmt.Println("couldnt insert into the subscription table with error ", err.Error())
		return common.Subscription{}, err
	}
	return common.Subscription{}, nil
}

// Read
func ListRecords(offset, upperLimit int) ([]common.Subscription, error) {
	return []common.Subscription{}, nil
}

// Update
func UpdateRecord(name, description string, id, projectId int, query string) (common.Subscription, error) {
	return common.Subscription{}, nil
}

// Delete
func DeleteRecord(id int) error {
	return nil
}
