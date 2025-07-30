package db

import (
	"common"
	"fmt"
	"strconv"
	"time"

	"github.com/gofrs/uuid/v5"
)

const (
	SubscriptionTableName = "subscriptions"
)

// TODO do all the methods

// CRUD db methods
// Create
func CreateNewRecord(price int, serviceName, userIdString string) (common.Subscription, error) {
	userId, _ := uuid.FromString(userIdString)
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
func GetRecord(userId, service string) (common.Subscription, error) {
	psqlQuery := fmt.Sprintf(`SELECT * FROM %s WHERE user_id='%s' AND service_name='%s'`, SubscriptionTableName, userId, service)
	//_, err := PostgresDB.Exec(psqlQuery)
	row := PostgresDB.QueryRow(psqlQuery)
	var subscription common.Subscription
	if err := row.Scan(&subscription.ServiceName, &subscription.Price, &subscription.UserId, &subscription.StartDate); err != nil {
		fmt.Println("couldnt scan from the subscription rows with error ", err.Error())
		return common.Subscription{}, err
	}

	return subscription, nil
}

// Update
func UpdateRecord(name, description string, id, projectId int, query string) (common.Subscription, error) {
	return common.Subscription{}, nil
}

// Delete
func DeleteRecord(userIdString, service string) error {
	userId, _ := uuid.FromString(userIdString)
	psqlQuery := fmt.Sprintf(`DELETE FROM %s WHERE user_id='%s' AND service_name='%s'`, SubscriptionTableName, userId, service)
	_, err := PostgresDB.Exec(psqlQuery)
	if err != nil {
		fmt.Println("couldnt insert into the subscription table with error ", err.Error())
		return err
	}
	return nil
}

// List
func ListRecords() ([]common.Subscription, error) {
	return []common.Subscription{}, nil
}
