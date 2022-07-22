package Service

import (
	"awesomeSystem/app/model"
	"awesomeSystem/app/weight/aggregate"
	"awesomeSystem/app/weight/entity"
	"awesomeSystem/utils/ACS"
	"awesomeSystem/utils/APIResponse"
	"awesomeSystem/utils/DB/mongodb"
	"awesomeSystem/utils/DB/redisdb"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	collectionRecord          = "record"
	collectionWeighRecord     = "weigh_record"
	collectionCalculateRecord = "calculate_record"
	collectionParameter       = "parameter"
	collectionCraft           = "craft"
	collectionTexture         = "texture"
	collectionProcess         = "process"
	collectionPurchaseStatus  = "purchase_status"
	collectionMaterial        = "material"
	collectionSupplier        = "supplier"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("no document")
	NoErr              = errors.New("no errors")
	NoDocument         = errors.New("mongo: no documents in result")
	NoParameters       = errors.New("no parameters")
)

func GetUserAuth(c *gin.Context) {
	user := c.Param("user")
	per := ACS.Enforcer.GetFilteredNamedPolicy("p", 0, user)
	//fmt.Println(per)

	res := make(map[string][]string)
	for _, v := range per {
		//fmt.Println(k, v)
		//fmt.Println(res)
		pSplit := strings.Split(v[1], "/")
		//fmt.Println("pSplit", pSplit)
		//fmt.Println("pSplit-1", pSplit[1])
		if value, ok := res[pSplit[1]]; ok {
			//fmt.Println("value", value)
			//fmt.Println("ok", ok)
			value = append(value, fmt.Sprintf("%s %s", v[1], v[2]))
			//fmt.Println("value", value)
			res[pSplit[1]] = value
		} else {
			if pSplit[1] == "weight" {
				l := []string{fmt.Sprintf("%s %s", v[1], v[2])}
				res[pSplit[1]] = l
			}
		}
		//fmt.Println(res)
		//fmt.Println("****************")
	}
	//fmt.Println(res)
	APIResponse.Success(c, 200, "用户权限列表", res)

}

// *********************************************************
// Crate
// *********************************************************

// GetAllCraft 所有工艺
func GetAllCraft(c *gin.Context) {
	zap.L().Info("GetAllCraft", zap.Any("调用 Service", "GetAllCraft 处理请求"))
	var results []model.Craft

	// 查询总数
	_, size := mongodb.NewMgo(collectionCraft).Count()
	//fmt.Printf(" documents name documents size %d \n", size)
	cur := mongodb.NewMgo(collectionCraft).FindAll(0, size, 1)

	defer cur.Close(context.TODO())
	//if cur != nil {
	//	fmt.Println("FindAll :", cur)
	//	//err = errors.New("FindAll err")
	//}
	for cur.Next(context.TODO()) {
		var elem model.Craft
		err := cur.Decode(&elem)
		if err != nil {
			//err = errors.New("FindAll err")
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		//log.Fatal(err)
		//zap.L().Info("GetAllCraft", zap.String("GetAllCraft error", fmt.Sprintf("--------->%s", zap.Error(err).String)))
		APIResponse.Err(c, http.StatusBadRequest, 400, "数据错误", err.Error())
	}
	//fmt.Println(results)
	//zap.L().Info("GetAllCraft", zap.String("GetAllCraft", fmt.Sprintf("--------->")))
	APIResponse.Success(c, 200, "GetAllCraft success", results)

}

// AddCraft 新增工艺
func AddCraft(c *gin.Context) {
	zap.L().Info("AddCraft", zap.Any("调用 Service", "AddCraft 处理请求"))
	var craft entity.NewCraft
	err := c.BindJSON(&craft)
	if err != nil {
		return
	}
	if craft.Name == "" {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddCraft 参数错误", craft.Name)
	}

	//查询是否存在craft
	var existCraft model.Craft
	err = mongodb.NewMgo(collectionCraft).FindOne("name", craft.Name).Decode(&existCraft)
	//fmt.Println(err)
	//判断是否存在记录
	if err == nil {
		// 已存在
		APIResponse.Err(c, http.StatusOK, 400, "AddCraft 工艺已存在", fmt.Sprintf("工艺已存在：%s", craft.Name))

	} else if err.Error() == NoDocument.Error() {
		// 不存在，可以新增

		insertOneResult := mongodb.NewMgo(collectionCraft).InsertOne(craft)
		err = nil
		//fmt.Println(insertOneResult)
		APIResponse.Success(c, 200, "AddCraft success", fmt.Sprintf("%s", insertOneResult.InsertedID))

	} else {
		// 如果查询错误
		//fmt.Println("***********")
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "AddCraft 发生错误", err)
	}

}

// DeleteCraftWithId 删除工艺
func DeleteCraftWithId(c *gin.Context) {
	zap.L().Info("DeleteCraftWithId", zap.Any("调用 Service", "DeleteCraftWithId 处理请求"))
	id := c.Param("id")

	//var craft entity.CraftId
	//err := c.BindJSON(&craft)
	//if err != nil {
	//	APIResponse.Err(c, http.StatusBadRequest, 400, "DeleteCraftWithId 参数错误", craft.Id)
	//	return
	//}

	objId, _ := primitive.ObjectIDFromHex(id)
	//查询是否存在craft
	var existCraft model.Craft
	err := mongodb.NewMgo(collectionCraft).FindOne("_id", objId).Decode(&existCraft)
	//fmt.Println(craftId)
	//fmt.Println(err)
	//判断是否存在记录
	if err == nil {
		// 已存在，可删除
		deleteResult := mongodb.NewMgo(collectionCraft).Delete("_id", objId)
		APIResponse.Success(c, 200, "DeleteCraftWithId success", fmt.Sprintf("工艺已删除：%d", deleteResult))

	} else {
		// 如果查询错误
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "DeleteCraftWithId id发生错误", err)
	}

}

// UpdateCraft 更新工艺
func UpdateCraft(c *gin.Context) {
	zap.L().Info("UpdateCraft", zap.Any("调用 Service", "UpdateCraft 处理请求"))
	id := c.Param("id")
	var craft entity.UpCraft
	err := c.BindJSON(&craft)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "UpdateCraft 参数错误", craft.Name)
		return
	}

	if craft.Name == "" || craft.ClientId == "" {
		APIResponse.Err(c, http.StatusBadRequest, 400, "UpdateCraft 参数错误", craft.Name)
		return
	}
	objId, _ := primitive.ObjectIDFromHex(id)

	//查询是否存在craft
	var existCraft model.Craft
	err = mongodb.NewMgo(collectionCraft).FindOne("_id", objId).Decode(&existCraft)

	//判断是否存在记录
	if err == nil {
		// 已存在,可更新
		// redis 分布式锁
		//res, err := redisdb.RdsClient.SetNX(clientId, existCraft.Id, time.Minute).Result()
		if r, _ := redisdb.GetLock(existCraft.Id, craft.ClientId); r {
			//fmt.Println("***********************")
			//fmt.Println(r)
			// 设置成功
			update := bson.D{
				{
					"$set", bson.D{
						{"name", craft.Name},
					}},
			}
			//time.Sleep(30 * time.Second)
			updateResult := mongodb.NewMgo(collectionCraft).UpdateOne("_id", objId, update)

			_, _ = redisdb.RunEvalDel(existCraft.Id, craft.ClientId)
			APIResponse.Success(c, 200, "UpdateCraft success", fmt.Sprintf("工艺已更新：%s", fmt.Sprintf("Matched %d documents and updated %d documents.", updateResult.MatchedCount, updateResult.ModifiedCount)))

		} else {
			ms, _ := redisdb.Pttl(existCraft.Id)
			APIResponse.Err(c, http.StatusOK, 400, "UpdateCraft 发生冲突", fmt.Sprintf("%d ms之后解锁", ms))

		}

	} else if err.Error() == NoDocument.Error() {
		// 不存在

		APIResponse.Err(c, http.StatusOK, 400, "UpdateCraft 没有该记录", "")
	} else {
		// 如果查询错误
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "UpdateCraft 发生错误", err)
	}

}

// *********************************************************
// process
// *********************************************************

// GetAllProcess 所有工序
func GetAllProcess(c *gin.Context) {
	zap.L().Info("GetAllProcess", zap.Any("调用 Service", "GetAllProcess 处理请求"))
	var results []model.Process

	m := mongodb.NewMgo(collectionProcess)
	// 查询总数
	_, size := m.Count()
	//fmt.Printf(" documents name documents size %d \n", size)
	cur := m.FindAll(0, size, 1)

	defer cur.Close(context.TODO())
	//if cur != nil {
	//	fmt.Println("FindAll :", cur)
	//	//err = errors.New("FindAll err")
	//}
	for cur.Next(context.TODO()) {
		var elem model.Process
		err := cur.Decode(&elem)
		if err != nil {
			//err = errors.New("FindAll err")
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		//log.Fatal(err)
		//zap.L().Info("GetAllProcess", zap.String("GetAllProcess error", fmt.Sprintf("--------->%s", zap.Error(err).String)))
		APIResponse.Err(c, http.StatusBadRequest, 400, "数据错误", err.Error())
	}
	//fmt.Println(results)
	//zap.L().Info("GetAllProcess", zap.String("GetAllProcess", fmt.Sprintf("--------->")))
	APIResponse.Success(c, 200, "GetAllProcess success", results)

}

// AddProcess 新增工序
func AddProcess(c *gin.Context) {
	zap.L().Info("AddProcess", zap.Any("调用 Service", "AddProcess 处理请求"))
	var n entity.NewProcess
	err := c.BindJSON(&n)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddProcess 参数错误", n.Name)
	}
	if n.Name == "" {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddProcess 参数错误", n.Name)
		return
	}

	m := mongodb.NewMgo(collectionProcess)
	//查询是否存在
	var exist model.Process
	err = m.FindOne("name", n.Name).Decode(&exist)
	//fmt.Println(err)
	//判断是否存在记录
	if err == nil {
		// 已存在
		APIResponse.Err(c, http.StatusOK, 400, "AddProcess 工序已存在", fmt.Sprintf("工序已存在：%s", n.Name))

	} else if err.Error() == NoDocument.Error() {
		// 不存在，可以新增

		insertOneResult := m.InsertOne(n)
		APIResponse.Success(c, 200, "AddProcess success", fmt.Sprintf("%s", insertOneResult.InsertedID))

	} else {
		// 如果查询错误
		//fmt.Println("***********")
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "AddProcess 发生错误", err)
	}

}

// DeleteProcessWithId 删除工序
func DeleteProcessWithId(c *gin.Context) {
	zap.L().Info("DeleteProcessWithId", zap.Any("调用 Service", "DeleteProcessWithId 处理请求"))
	id := c.Param("id")

	//var process entity.ProcessId
	//err := c.BindJSON(&process)
	//if err != nil {
	//	APIResponse.Err(c, http.StatusBadRequest, 400, "DeleteCraftWithId 参数错误", process.Id)
	//	return
	//}

	objId, _ := primitive.ObjectIDFromHex(id)
	m := mongodb.NewMgo(collectionProcess)
	//查询是否存在
	var exist model.Process
	err := m.FindOne("_id", objId).Decode(&exist)
	//fmt.Println(craftId)
	//fmt.Println(err)
	//判断是否存在记录
	if err == nil {
		// 已存在，可删除
		deleteResult := m.Delete("_id", objId)
		APIResponse.Success(c, 200, "success", fmt.Sprintf("DeleteProcessWithId 工序已删除：%d", deleteResult))

	} else {
		// 如果查询错误
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "DeleteProcessWithId id发生错误", err)
	}

}

// UpdateProcess 更新工序
func UpdateProcess(c *gin.Context) {
	zap.L().Info("UpdateProcess", zap.Any("调用 Service", "UpdateProcess 处理请求"))
	id := c.Param("id")
	var u entity.UpProcess
	err := c.BindJSON(&u)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "UpdateProcess 参数错误", u.Name)
		return
	}

	if u.Name == "" || u.ClientId == "" {
		APIResponse.Err(c, http.StatusBadRequest, 400, "UpdateProcess 参数错误", u.Name)
		return
	}
	objId, _ := primitive.ObjectIDFromHex(id)

	m := mongodb.NewMgo(collectionProcess)
	//查询是否存在
	var exist model.Process
	err = m.FindOne("_id", objId).Decode(&exist)

	//判断是否存在记录
	if err == nil {
		// 已存在,可更新
		// redis 分布式锁
		//res, err := redisdb.RdsClient.SetNX(clientId, existCraft.Id, time.Minute).Result()
		if r, _ := redisdb.GetLock(exist.Id, u.ClientId); r {
			//fmt.Println("***********************")
			//fmt.Println(r)
			// 设置成功
			update := bson.D{
				{
					"$set", bson.D{
						{"name", u.Name},
					}},
			}
			//time.Sleep(30 * time.Second)
			updateResult := m.UpdateOne("_id", objId, update)

			_, _ = redisdb.RunEvalDel(exist.Id, u.ClientId)
			APIResponse.Success(c, 200, "UpdateProcess success", fmt.Sprintf("工序已更新：%s", fmt.Sprintf("Matched %d documents and updated %d documents.", updateResult.MatchedCount, updateResult.ModifiedCount)))

		} else {
			ms, _ := redisdb.Pttl(exist.Id)
			APIResponse.Err(c, http.StatusOK, 400, "UpdateProcess 发生冲突", fmt.Sprintf("%d ms之后解锁", ms))

		}

	} else if err.Error() == NoDocument.Error() {
		// 不存在

		APIResponse.Err(c, http.StatusOK, 400, "UpdateProcess 没有该记录", "")
	} else {
		// 如果查询错误
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "UpdateProcess 发生错误", err)
	}

}

// *********************************************************
// texture
// *********************************************************

// GetAllTexture 所有材质
func GetAllTexture(c *gin.Context) {
	zap.L().Info("GetAllTexture", zap.Any("调用 Service", "GetAllTexture 处理请求"))
	var results []model.Texture

	m := mongodb.NewMgo(collectionTexture)
	// 查询总数
	_, size := m.Count()
	//fmt.Printf(" documents name documents size %d \n", size)
	cur := m.FindAll(0, size, 1)

	defer cur.Close(context.TODO())
	//if cur != nil {
	//	fmt.Println("FindAll :", cur)
	//	//err = errors.New("FindAll err")
	//}
	for cur.Next(context.TODO()) {
		var elem model.Texture
		err := cur.Decode(&elem)
		if err != nil {
			//err = errors.New("FindAll err")
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		//log.Fatal(err)
		//zap.L().Info("GetAllTexture", zap.String("GetAllTexture error", fmt.Sprintf("--------->%s", zap.Error(err).String)))
		APIResponse.Err(c, http.StatusBadRequest, 400, "数据错误", err.Error())
	}
	//fmt.Println(results)
	//zap.L().Info("GetAllTexture", zap.String("GetAllTexture", fmt.Sprintf("--------->")))
	APIResponse.Success(c, 200, "GetAllTexture success", results)

}

// AddTexture 新增材质
func AddTexture(c *gin.Context) {
	zap.L().Info("AddTexture", zap.Any("调用 Service", "AddTexture 处理请求"))
	var n entity.NewTexture
	err := c.BindJSON(&n)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddTexture 参数错误", n.Name)
		return
	}
	if n.Name == "" {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddTexture 参数错误", n.Name)
		return
	}

	m := mongodb.NewMgo(collectionTexture)
	//查询是否存在
	var exist model.Texture
	err = m.FindOne("name", n.Name).Decode(&exist)
	//fmt.Println(err)
	//判断是否存在记录
	if err == nil {
		// 已存在
		APIResponse.Err(c, http.StatusOK, 400, "AddTexture 材质已存在", fmt.Sprintf("材质已存在：%s", n.Name))

	} else if err.Error() == NoDocument.Error() {
		// 不存在，可以新增

		insertOneResult := m.InsertOne(n)
		APIResponse.Success(c, 200, "AddTexture success", fmt.Sprintf("%s", insertOneResult.InsertedID))

	} else {
		// 如果查询错误
		//fmt.Println("***********")
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "AddTexture 发生错误", err)
	}

}

// DeleteTextureWithId 删除材质
func DeleteTextureWithId(c *gin.Context) {
	zap.L().Info("DeleteTextureWithId", zap.Any("调用 Service", "DeleteTextureWithId 处理请求"))
	id := c.Param("id")

	//var texture entity.TextureId
	//err := c.BindJSON(&texture)
	//if err != nil {
	//	APIResponse.Err(c, http.StatusBadRequest, 400, "DeleteCraftWithId 参数错误", texture.Id)
	//	return
	//}

	objId, _ := primitive.ObjectIDFromHex(id)
	m := mongodb.NewMgo(collectionTexture)
	//查询是否存在
	var exist model.Texture
	err := m.FindOne("_id", objId).Decode(&exist)
	//fmt.Println(craftId)
	//fmt.Println(err)
	//判断是否存在记录
	if err == nil {
		// 已存在，可删除
		deleteResult := m.Delete("_id", objId)
		APIResponse.Success(c, 200, "DeleteTextureWithId success", fmt.Sprintf("材质已删除：%d", deleteResult))

	} else {
		// 如果查询错误
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "DeleteTextureWithId id发生错误", err)
	}

}

// UpdateTexture 更新材质
func UpdateTexture(c *gin.Context) {
	zap.L().Info("UpdateTexture", zap.Any("调用 Service", "UpdateTexture 处理请求"))
	id := c.Param("id")
	var u entity.UpTexture
	err := c.BindJSON(&u)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "UpdateTexture 参数错误", u.Name)
		return
	}

	if u.Name == "" || u.ClientId == "" {
		APIResponse.Err(c, http.StatusBadRequest, 400, "UpdateTexture 参数错误", u.Name)
		return
	}
	objId, _ := primitive.ObjectIDFromHex(id)

	m := mongodb.NewMgo(collectionTexture)
	//查询是否存在
	var exist model.Texture
	err = m.FindOne("_id", objId).Decode(&exist)

	//判断是否存在记录
	if err == nil {
		// 已存在,可更新
		// redis 分布式锁
		//res, err := redisdb.RdsClient.SetNX(clientId, existCraft.Id, time.Minute).Result()
		if r, _ := redisdb.GetLock(exist.Id, u.ClientId); r {
			//fmt.Println("***********************")
			//fmt.Println(r)
			// 设置成功
			update := bson.D{
				{
					"$set", bson.D{
						{"name", u.Name},
					}},
			}
			//time.Sleep(30 * time.Second)
			updateResult := m.UpdateOne("_id", objId, update)

			_, _ = redisdb.RunEvalDel(exist.Id, u.ClientId)
			APIResponse.Success(c, 200, "UpdateTexture success", fmt.Sprintf("材质已更新：%s", fmt.Sprintf("Matched %d documents and updated %d documents.", updateResult.MatchedCount, updateResult.ModifiedCount)))

		} else {
			ms, _ := redisdb.Pttl(exist.Id)
			APIResponse.Err(c, http.StatusOK, 400, "UpdateTexture 发生冲突", fmt.Sprintf("%d ms之后解锁", ms))

		}

	} else if err.Error() == NoDocument.Error() {
		// 不存在

		APIResponse.Err(c, http.StatusOK, 400, "UpdateTexture 没有该记录", "")
	} else {
		// 如果查询错误
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "UpdateTexture 发生错误", err)
	}

}

// *********************************************************
// purchase status
// *********************************************************

// GetAllPurchaseStatus 所有采购状态
func GetAllPurchaseStatus(c *gin.Context) {
	zap.L().Info("GetAllPurchaseStatus", zap.Any("调用 Service", "GetAllPurchaseStatus 处理请求"))
	var results []model.PurchaseStatus

	m := mongodb.NewMgo(collectionPurchaseStatus)
	// 查询总数
	_, size := m.Count()
	//fmt.Printf(" documents name documents size %d \n", size)
	cur := m.FindAll(0, size, 1)

	defer cur.Close(context.TODO())
	//if cur != nil {
	//	fmt.Println("FindAll :", cur)
	//	//err = errors.New("FindAll err")
	//}
	for cur.Next(context.TODO()) {
		var elem model.PurchaseStatus
		err := cur.Decode(&elem)
		if err != nil {
			//err = errors.New("FindAll err")
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		//log.Fatal(err)
		//zap.L().Info("GetAllPurchaseStatus", zap.String("GetAllPurchaseStatus error", fmt.Sprintf("--------->%s", zap.Error(err).String)))
		APIResponse.Err(c, http.StatusBadRequest, 400, "数据错误", err.Error())
	}
	//fmt.Println(results)
	//zap.L().Info("GetAllProcess", zap.String("GetAllPurchaseStatus", fmt.Sprintf("--------->")))
	APIResponse.Success(c, 200, "GetAllPurchaseStatus success", results)

}

// AddPurchaseStatus 新增工序
func AddPurchaseStatus(c *gin.Context) {
	zap.L().Info("AddPurchaseStatus", zap.Any("调用 Service", "AddPurchaseStatus 处理请求"))
	var n entity.NewPurchaseStatus
	err := c.BindJSON(&n)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddPurchaseStatus 参数错误", n.Name)
		return
	}
	if n.Name == "" {
		APIResponse.Err(c, http.StatusBadRequest, 400, "AddPurchaseStatus 参数错误", n.Name)
		return
	}

	m := mongodb.NewMgo(collectionPurchaseStatus)
	//查询是否存在
	var exist model.PurchaseStatus
	err = m.FindOne("name", n.Name).Decode(&exist)
	//fmt.Println(err)
	//判断是否存在记录
	if err == nil {
		// 已存在
		APIResponse.Err(c, http.StatusOK, 400, "AddPurchaseStatus 采购状态已存在", fmt.Sprintf("采购状态已存在：%s", n.Name))

	} else if err.Error() == NoDocument.Error() {
		// 不存在，可以新增

		insertOneResult := m.InsertOne(n)
		APIResponse.Success(c, 200, "AddPurchaseStatus success", fmt.Sprintf("%s", insertOneResult.InsertedID))

	} else {
		// 如果查询错误
		//fmt.Println("***********")
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "AddPurchaseStatus 发生错误", err)
	}

}

// DeletePurchaseStatusWithId 删除采购状态
func DeletePurchaseStatusWithId(c *gin.Context) {
	zap.L().Info("DeletePurchaseStatusWithId", zap.Any("调用 Service", "DeletePurchaseStatusWithId 处理请求"))
	id := c.Param("id")

	//var ps entity.PurchaseStatusId
	//err := c.BindJSON(&ps)
	//if err != nil {
	//	APIResponse.Err(c, http.StatusBadRequest, 400, "DeleteCraftWithId 参数错误", ps.Id)
	//	return
	//}

	objId, _ := primitive.ObjectIDFromHex(id)
	m := mongodb.NewMgo(collectionPurchaseStatus)
	//查询是否存在
	var exist model.PurchaseStatus
	err := m.FindOne("_id", objId).Decode(&exist)
	//fmt.Println(craftId)
	//fmt.Println(err)
	//判断是否存在记录
	if err == nil {
		// 已存在，可删除
		deleteResult := m.Delete("_id", objId)
		APIResponse.Success(c, 200, "DeletePurchaseStatusWithId success", fmt.Sprintf("采购状态已删除：%d", deleteResult))

	} else {
		// 如果查询错误
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "DeletePurchaseStatusWithId id发生错误", err)
	}

}

// UpdatePurchaseStatus 更新采购状态
func UpdatePurchaseStatus(c *gin.Context) {
	zap.L().Info("UpdatePurchaseStatus", zap.Any("调用 Service", "UpdatePurchaseStatus 处理请求"))
	id := c.Param("id")
	var u entity.UpPurchaseStatus
	err := c.BindJSON(&u)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "UpdatePurchaseStatus 参数错误", u.Name)
		return
	}

	if u.Name == "" || u.ClientId == "" {
		APIResponse.Err(c, http.StatusBadRequest, 400, "UpdatePurchaseStatus 参数错误", u.Name)
		return
	}
	objId, _ := primitive.ObjectIDFromHex(id)

	m := mongodb.NewMgo(collectionPurchaseStatus)
	//查询是否存在
	var exist model.PurchaseStatus
	err = m.FindOne("_id", objId).Decode(&exist)

	//判断是否存在记录
	if err == nil {
		// 已存在,可更新
		// redis 分布式锁
		//res, err := redisdb.RdsClient.SetNX(clientId, existCraft.Id, time.Minute).Result()
		if r, _ := redisdb.GetLock(exist.Id, u.ClientId); r {
			//fmt.Println("***********************")
			//fmt.Println(r)
			// 设置成功
			update := bson.D{
				{
					"$set", bson.D{
						{"name", u.Name},
					}},
			}
			//time.Sleep(30 * time.Second)
			updateResult := m.UpdateOne("_id", objId, update)

			_, _ = redisdb.RunEvalDel(exist.Id, u.ClientId)
			APIResponse.Success(c, 200, "UpdatePurchaseStatus success", fmt.Sprintf("采购状态已更新：%s", fmt.Sprintf("Matched %d documents and updated %d documents.", updateResult.MatchedCount, updateResult.ModifiedCount)))

		} else {
			ms, _ := redisdb.Pttl(exist.Id)
			APIResponse.Err(c, http.StatusOK, 400, "UpdatePurchaseStatus 发生冲突", fmt.Sprintf("%d ms之后解锁", ms))

		}

	} else if err.Error() == NoDocument.Error() {
		// 不存在

		APIResponse.Err(c, http.StatusOK, 400, "UpdatePurchaseStatus 没有该记录", "")
	} else {
		// 如果查询错误
		//fmt.Println(err)
		APIResponse.Err(c, http.StatusOK, 400, "UpdatePurchaseStatus 发生错误", err)
	}

}

// *********************************************************
// 多条件搜索
// *********************************************************

// WeighMultiConditionSearch 称重多条件查询
func WeighMultiConditionSearch(c *gin.Context) {
	zap.L().Info("WeighMultiConditionSearch", zap.Any("调用 Service", "WeighMultiConditionSearch 处理请求"))
	var mc entity.WeighMultiCondition
	err := c.BindJSON(&mc)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "WeighMultiConditionSearch", err.Error())
	}

	//fmt.Println(mc)
	//dataStatResult := make(chan bson.M)
	//var output ResultSlice
	//var wg sync.WaitGroup
	m := mongodb.NewMgo(collectionWeighRecord)
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go aggregate.DataAggregate(m, mc, dataStatResult, &wg)
	//}
	//
	//fmt.Println("dataStatResult len:   ")
	//fmt.Println(dataStatResult)
	//fmt.Println(len(dataStatResult))
	////if len(dataStatResult) == 0 {
	////	return
	////}
	//
	//for value := range dataStatResult{
	//	fmt.Println("*********value**********")
	//	fmt.Println(value)
	//	output = append(output, OutPut{
	//		MaterialCode: value["material_code"].(string),
	//		Supplier:     value["supplier"].(string),
	//		Craft:        value["craft"].(string),
	//	})
	//	fmt.Println(len(output))
	//	fmt.Println(output)
	//	if len(output) == 200  {
	//		break
	//	}
	//}
	//wg.Wait()
	//
	//fmt.Println("#####################")
	//fmt.Println("wait")
	//for _, v := range output {
	//	result, err := json.Marshal(&v)
	//	if err != nil {
	//		fmt.Printf("json.marshal failed, err: %s", err)
	//	}
	//	fmt.Println(string(result))
	//}

	//var results [][]*processNode
	var results []model.WeighMaterialRecord

	// 分页逻辑
	//var pages int64 = int64(size / pageSize)
	var skip int = mc.PageSize * (mc.PageNum - 1)

	matchStage := aggregate.GenPipeline(mc, skip, mc.PageSize)
	opts := options.Aggregate().SetMaxTime(15 * time.Second)
	//cur := m.FindAggregation(mongo.Pipeline{matchStage}, opts)
	cur := m.FindAggregation(matchStage, opts)

	//if err := cur.All(context.TODO(), &results); err != nil {
	//	utils.GetLogger().Info(fmt.Sprintf("dataAggregate error: %v", err))
	//}

	defer cur.Close(context.TODO())
	if cur != nil {
		fmt.Println("Search :", cur)
		//err = errors.New("FindAll err")
		//APIResponse.Err(c, http.StatusBadRequest, 400, "WeighMultiConditionSearch", err.Error())
		//zap.L().Info("WeighMultiConditionSearch", zap.Any("返回错误", err.Error()))
	}
	for cur.Next(context.TODO()) {

		var elem model.WeighMaterialRecord
		err := cur.Decode(&elem)
		if err != nil {
			//log.Fatal(err)
			APIResponse.Err(c, http.StatusBadRequest, 400, "WeighMultiConditionSearch", err.Error())
		}
		//fmt.Println("*************")
		//fmt.Println(elem)
		//tmpList := make([]*processNode, 4, 6)
		//var tmpList []*processNode
		//for _, value := range elem.FlowProcess {
		//
		//		pn := processNode{
		//			Id:                 elem.Id,
		//			MaterialCode:       elem.MaterialCode,
		//			MaterialType:       elem.MaterialType,
		//			MaterialName:       elem.MaterialName,
		//			Specifications:     elem.Specifications,
		//			Supplier:           elem.Supplier,
		//			Craft:              elem.Craft,
		//			Texture:            elem.Texture,
		//			Process:            elem.Process,
		//			PurchaseStatus:     elem.PurchaseStatus,
		//			ReceivingWarehouse: elem.ReceivingWarehouse,
		//			WeighStage:         value.WeighStage,
		//			RecordLog:          value.RecordLog,
		//		}
		//		tmpList = append(tmpList, &pn)
		//	}

		results = append(results, elem)
	}

	rData := entity.WeightMaterialRecord{
		Data:  results,
		Total: int64(len(results)),
	}
	//fmt.Println(rData)

	if err := cur.Err(); err != nil {
		//log.Fatal(err)
		APIResponse.Err(c, http.StatusBadRequest, 400, "WeighMultiConditionSearch", err.Error())
	}

	//zap.L().Info("WeighMultiConditionSearch", zap.Any("调用 Service", "WeighMultiConditionSearch 处理请求"), zap.Any("处理返回值", rData))
	APIResponse.Success(c, 200, "WeighMultiConditionSearch success", rData)

}

// CalMultiConditionSearch 理算多条件查询
func CalMultiConditionSearch(c *gin.Context) {
	zap.L().Info("CalMultiConditionSearch", zap.Any("调用 Service", "CalMultiConditionSearch 处理请求"))
	var mc entity.CalMultiCondition
	err := c.BindJSON(&mc)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "CalMultiConditionSearch", err.Error())
	}

	//fmt.Println(mc)
	//dataStatResult := make(chan bson.M)
	//var output ResultSlice
	//var wg sync.WaitGroup
	m := mongodb.NewMgo(collectionCalculateRecord)
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go aggregate.DataAggregate(m, mc, dataStatResult, &wg)
	//}
	//
	//fmt.Println("dataStatResult len:   ")
	//fmt.Println(dataStatResult)
	//fmt.Println(len(dataStatResult))
	////if len(dataStatResult) == 0 {
	////	return
	////}
	//
	//for value := range dataStatResult{
	//	fmt.Println("*********value**********")
	//	fmt.Println(value)
	//	output = append(output, OutPut{
	//		MaterialCode: value["material_code"].(string),
	//		Supplier:     value["supplier"].(string),
	//		Craft:        value["craft"].(string),
	//	})
	//	fmt.Println(len(output))
	//	fmt.Println(output)
	//	if len(output) == 200  {
	//		break
	//	}
	//}
	//wg.Wait()
	//
	//fmt.Println("#####################")
	//fmt.Println("wait")
	//for _, v := range output {
	//	result, err := json.Marshal(&v)
	//	if err != nil {
	//		fmt.Printf("json.marshal failed, err: %s", err)
	//	}
	//	fmt.Println(string(result))
	//}

	//var results [][]*processNode
	var results []model.CalculateMaterialRecord

	// 分页逻辑
	//var pages int64 = int64(size / pageSize)
	var skip int = mc.PageSize * (mc.PageNum - 1)

	matchStage := aggregate.GenCalPipeline(mc, skip, mc.PageSize)
	opts := options.Aggregate().SetMaxTime(15 * time.Second)
	//cur := m.FindAggregation(mongo.Pipeline{matchStage}, opts)
	cur := m.FindAggregation(matchStage, opts)

	//if err := cur.All(context.TODO(), &results); err != nil {
	//	utils.GetLogger().Info(fmt.Sprintf("dataAggregate error: %v", err))
	//}

	defer cur.Close(context.TODO())
	if cur != nil {
		fmt.Println("Search :", cur)
		//err = errors.New("FindAll err")
		//APIResponse.Err(c, http.StatusBadRequest, 400, "WeighMultiConditionSearch", err.Error())
		//zap.L().Info("WeighMultiConditionSearch", zap.Any("返回错误", err.Error()))
	}
	for cur.Next(context.TODO()) {

		var elem model.CalculateMaterialRecord
		err := cur.Decode(&elem)
		if err != nil {
			//log.Fatal(err)
			APIResponse.Err(c, http.StatusBadRequest, 400, "CalMultiConditionSearch", err.Error())
		}
		//fmt.Println("*************")
		//fmt.Println(elem)
		//tmpList := make([]*processNode, 4, 6)
		//var tmpList []*processNode
		//for _, value := range elem.FlowProcess {
		//
		//		pn := processNode{
		//			Id:                 elem.Id,
		//			MaterialCode:       elem.MaterialCode,
		//			MaterialType:       elem.MaterialType,
		//			MaterialName:       elem.MaterialName,
		//			Specifications:     elem.Specifications,
		//			Supplier:           elem.Supplier,
		//			Craft:              elem.Craft,
		//			Texture:            elem.Texture,
		//			Process:            elem.Process,
		//			PurchaseStatus:     elem.PurchaseStatus,
		//			ReceivingWarehouse: elem.ReceivingWarehouse,
		//			WeighStage:         value.WeighStage,
		//			RecordLog:          value.RecordLog,
		//		}
		//		tmpList = append(tmpList, &pn)
		//	}
		results = append(results, elem)
	}
	rData := entity.CalMaterialRecord{
		Data:  results,
		Total: int64(len(results)),
	}
	//fmt.Println(rData)

	if err := cur.Err(); err != nil {
		//log.Fatal(err)
		APIResponse.Err(c, http.StatusBadRequest, 400, "CalMultiConditionSearch", err.Error())
	}
	//zap.L().Info("WeighMultiConditionSearch", zap.Any("调用 Service", "WeighMultiConditionSearch 处理请求"), zap.Any("处理返回值", rData))
	APIResponse.Success(c, 200, "CalMultiConditionSearch success", rData)
}

// *********************************************************
// 权限
// *********************************************************

func AddPolicy(c *gin.Context) {
	newPolicy := entity.Policy{}
	err := c.BindJSON(&newPolicy)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "add Policy failed", "")
		return
	}
	//subject := "tom"
	//object := "/api/routers"
	//action := "POST"
	//cacheName := newPolicy.Subject + newPolicy.Object + newPolicy.Action
	result, _ := ACS.Enforcer.AddPolicy(newPolicy.Subject, newPolicy.Object, newPolicy.Action)
	if result {
		// 清除缓存
		//_ = Cache.GlobalCache.Delete(cacheName)
		APIResponse.Success(c, 200, "add Policy success", "")
	} else {
		APIResponse.Err(c, http.StatusBadRequest, 400, "add Policy failed", "")
	}
}

func DeletePolicy(c *gin.Context) {
	policy := entity.Policy{}
	err := c.BindJSON(&policy)
	if err != nil {
		APIResponse.Err(c, http.StatusBadRequest, 400, "delete Policy failed", "")
		return
	}
	result, _ := ACS.Enforcer.RemovePolicy(policy.Subject, policy.Object, policy.Action)
	if result {
		// 清除缓存 代码省略
		APIResponse.Success(c, 200, "delete Policy success", "")
	} else {
		APIResponse.Err(c, http.StatusBadRequest, 400, "delete Policy failed", "")
	}
}
