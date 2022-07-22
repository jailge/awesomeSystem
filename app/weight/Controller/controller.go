package Controller

import (
	"awesomeSystem/app/model"
	"awesomeSystem/utils/DB/mongodb"
	"context"
	"log"
)

func GetAllFun(results []interface{}, collection string) []interface{} {
	// 查询总数
	_, size := mongodb.NewMgo(collection).Count()
	//fmt.Printf(" documents name documents size %d \n", size)
	cur := mongodb.NewMgo(collection).FindAll(0, size, 1)

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
	return results
}
