package database

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/common/config"
	"main/common/dbconn"
	"main/common/tabletypes"
	"main/core/sendtx"
	"reflect"
	"strings"
)

func InsertUserInfo(username, password, email, phone string, identity tabletypes.Identity) error {
	username = strings.ToLower(username)
	email = strings.ToLower(email)
	filter := bson.M{"username": username}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		return fmt.Errorf("InsertUserInfo:err in getting Trans data: %v", err)
	}
	if len(idres) == 0 {
		var res = tabletypes.UserInfo{
			Id:       utils.UUIDv4(),
			Username: username,
			Password: password,
			Email:    email,
			Phone:    phone,
			Identity: identity,
			Approved: 0,
		}
		err := InsertDocument(config.DbcollectionUserInfo, res)
		if err != nil {
			return fmt.Errorf("InsertUserInfo:err in inserting NFTData")
		}
		return nil
	}
	return fmt.Errorf("Already Registed Username: %s", email)
}

func InsertProdInfo(id string, ecode string, codata string, operator string, waterdata string) (error, string) {
	id = uuid.New().String()
	txhash := sendtx.AddOrUpdateProdData(id, ecode, codata, operator, waterdata)
	var res = tabletypes.EcoResData{
		Id:        id,
		Ecodata:   ecode,
		Codata:    codata,
		Operator:  operator,
		Waterdata: waterdata,
		Hash:      txhash,
	}

	err := InsertDocument(config.DbcollectionEcoInfo, res)
	if err != nil {
		return fmt.Errorf("InsertProdInfo:err in inserting"), ""
	}
	return nil, txhash

}

func QueryUserInfos(username, password string) (error, *tabletypes.UserInfo) {
	username = strings.ToLower(username)
	filter := bson.M{"username": username, "password": password, "approved": 1}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	var resdata *tabletypes.UserInfo
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err), resdata
	}
	if len(idres) == 0 {
		return fmt.Errorf("Non-Registed/Not-Approved user, Please register firstly."), resdata
	}
	resdata = idres[0].(*tabletypes.UserInfo)
	return nil, resdata
}

func QueryAllUserInfo() (error, []*tabletypes.UserInfo) {
	filter := bson.M{"identity": tabletypes.User}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("QueryAllUserInfo err : %s", err), nil
	}
	if len(idres) != 0 {
		resdata := make([]*tabletypes.UserInfo, 0)
		for _, data := range idres {
			res := data.(*tabletypes.UserInfo)
			resdata = append(resdata, res)
		}
		return nil, resdata
	}
	return nil, nil
}

func QueryChainData() (error, []*tabletypes.EcoResData) {
	filter := bson.M{}
	err, idres := GetDocuments(config.DbcollectionEcoInfo, filter, &tabletypes.EcoResData{})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("QueryChainData err : %s", err), nil
	}
	if len(idres) != 0 {
		resdata := make([]*tabletypes.EcoResData, 0)
		for _, data := range idres {
			res := data.(*tabletypes.EcoResData)
			resdata = append(resdata, res)
		}
		return nil, resdata
	}
	return nil, nil
}
func QueryUserInfoByName(username string) (error, *tabletypes.UserInfo) {
	filter := bson.M{"username": username}
	var res *tabletypes.UserInfo
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfoByName err :%s", err), res
	}
	if len(idres) != 0 {
		res := idres[0].(*tabletypes.UserInfo)
		return nil, res
	}
	return fmt.Errorf("QueryUserInfoByName err :%s", err), res
}

func ApproveUserInfo(id string) error {
	filter := bson.M{"id": id}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err)
	}
	if len(idres) != 0 {
		update := bson.M{"$set": bson.M{"approved": 1}}
		err := UpdateDocument(config.DbcollectionUserInfo, filter, update)
		if err != nil {
			return fmt.Errorf("ApproveUserInfo err")
		}
		return nil
	}
	return fmt.Errorf("Approve non-exist user.")
}

func DeleteChaindata(id string) error {
	filter := bson.M{"id": id}
	err, idres := GetDocuments(config.DbcollectionEcoInfo, filter, &tabletypes.EcoResData{})
	if err != nil {
		return fmt.Errorf("DeleteChaindata: %v", err)
	}
	if len(idres) != 0 {
		collection := dbconn.GetCollection(config.DbcollectionEcoInfo)
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("DbcollectionEcoInfo: %v", err)
		}
		return nil
	}
	return fmt.Errorf("Delete non-exist eco.")
}

func DeleteUserInfo(id string) error {
	filter := bson.M{"id": id}
	err, idres := GetDocuments(config.DbcollectionUserInfo, filter, &tabletypes.UserInfo{})
	if err != nil {
		return fmt.Errorf("QueryUserInfo: %v", err)
	}
	if len(idres) != 0 {
		collection := dbconn.GetCollection(config.DbcollectionUserInfo)
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("DeleteUserInfoError: %v", err)
		}
		return nil
	}
	return fmt.Errorf("Delete non-exist user.")
}

func UpdateAdminUserInfo(username, phone, email string) error {
	//filter := bson.M{"username": strings.ToLower(username), "approved": 1}
	filter := bson.M{"username": strings.ToLower(username)}
	update := bson.M{"$set": bson.M{"email": email, "phone": phone}}
	err := UpdateDocument(config.DbcollectionUserInfo, filter, update)
	if err != nil {
		return fmt.Errorf("UpdateUserInfo err")
	}
	return nil

}

func InsertDocument(collectionname string, data interface{}) error {
	collection := dbconn.GetCollection(collectionname)
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return fmt.Errorf("failed to insert document: %v", err)
	}
	return nil
}

func GetDocuments(collectionname string, filter bson.M, result interface{}, opts ...*options.FindOptions) (error, []interface{}) {
	collection := dbconn.GetCollection(collectionname)
	cur, err := collection.Find(context.Background(), filter, opts...)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("failed to get documents: %v", err)), nil
	}
	defer cur.Close(context.Background())

	var results []interface{}
	for cur.Next(context.Background()) {
		// Create a new instance of the result type for each document
		res := reflect.New(reflect.TypeOf(result).Elem()).Interface()
		err := cur.Decode(res)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("failed to decode document: %v", err)), nil
		}
		results = append(results, res)
	}

	if err := cur.Err(); err != nil {
		return fmt.Errorf(fmt.Sprintf("cursor error: %v", err)), nil
	}

	return nil, results
}
func UpdateDocument(collectionname string, filter bson.M, update bson.M) error {
	//	fmt.Println("up")
	collection := dbconn.GetCollection(collectionname)
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %v", err)
	}
	return nil
}
