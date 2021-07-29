package log

import (
	"Raven/src/service"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

//TimePoint 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

//Record 一条日志
type Record struct {
	fileName  string      `bson:"fileName"`
	line      int         `bson:"line"`
	Err       interface{} `bson:"err"`
	TimePoint TimePoint   `bson:"timePoint"` // 执行时间点
}

func InsertLog(fileName string, line int, message interface{}) {
	var (
		client *mongo.Client
		err    error
		result *mongo.InsertOneResult
	)

	v := service.GetLogConnectString()
	path := strings.Join([]string{"mongodb://", v.UserName, ":", v.Password, "@", v.Ip, ":", v.Port}, "")
	// 建立mongodb连接
	clientOptions := options.Client().ApplyURI(path)
	if client, err = mongo.Connect(
		context.TODO(), clientOptions); err != nil {
		return
	}

	// 2, 选择数据库my_db
	database := client.Database("Helheim")

	// 3, 选择表my_collection
	collection := database.Collection("cron_log")
	// 4, 插入记录(bson)
	record := &Record{
		fileName:  fileName,
		line:      line,
		Err:       message,
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	if result, err = collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err)
		return
	}

	if result.InsertedID == nil {
		fmt.Println("Failed to insert log")
	}
}
