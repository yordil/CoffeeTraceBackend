package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl              string
	Port                     int
	DbName                   string
	UserCollection           string
	OrderCollection           string
	ProductCollection		string
	ContextTimeout           int
	AccessTokenExpiryHour    int
	RefreshTokenExpiryHour   int
	AccessTokenSecret        string
	RefreshTokenSecret       string
	NotificationCollection   string



	TransactionCollection string
	ReplyCollection		  	string
	BlogCollection			string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	portStr := os.Getenv("PORT")
	dbname := os.Getenv("DB_NAME")
	usercoll := os.Getenv("user_collection")
	orderColl := os.Getenv("order_collection")
	productColl := os.Getenv("product_collection")
	contextTimeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	notifycoll := os.Getenv("notification_collection")
	replycoll := os.Getenv("reply_collection")
	blogcoll := os.Getenv("blog_collection")
	accessTokenExpiryHourStr := os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR")
	refreshTokenExpiryHourStr := os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	transcoll := os.Getenv("Transaction_collection")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatal("Invalid PORT value")
		return nil, err
	}

	contextTimeout, err := strconv.Atoi(contextTimeoutStr)
	if err != nil {
		log.Fatal("Invalid CONTEXT_TIMEOUT value")
		return nil, err
	}

	accessTokenExpiryHour, err := strconv.Atoi(accessTokenExpiryHourStr)
	if err != nil {
		log.Fatal("Invalid ACCESS_TOKEN_EXPIRY_HOUR value")
		return nil, err
	}

	refreshTokenExpiryHour, err := strconv.Atoi(refreshTokenExpiryHourStr)
	if err != nil {
		log.Fatal("Invalid REFRESH_TOKEN_EXPIRY_HOUR value")
		return nil, err
	}

	config := &Config{
		DatabaseUrl:            dbURL,
		Port:                   port,
		DbName:                 dbname,
		UserCollection:         usercoll,
		OrderCollection:         orderColl,
		ProductCollection:		productColl,
		ContextTimeout:         contextTimeout,
		AccessTokenExpiryHour:  accessTokenExpiryHour,
		RefreshTokenExpiryHour: refreshTokenExpiryHour,
		AccessTokenSecret:      accessTokenSecret,
		RefreshTokenSecret:     refreshTokenSecret,
		NotificationCollection: notifycoll,

		TransactionCollection: transcoll,
		ReplyCollection:		replycoll,
		BlogCollection:			blogcoll,
	}

	return config, nil
}
