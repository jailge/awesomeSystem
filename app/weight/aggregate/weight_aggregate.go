package aggregate

import (
	"awesomeSystem/app/weight/entity"
	"awesomeSystem/utils/DB/mongodb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"sync"
	"time"
)

func GenAllRecordPipeline(sort string, skip int, limit int) (bson.D, bson.D, bson.D, bson.D) {
	//groupPip := bson.D{
	//	{
	//		"$group",
	//		bson.D{
	//			{
	//				"_id",
	//				bson.D{
	//					{"material_code", "$material_code"},
	//					{"material_type", "$material_type"},
	//					{"material_name", "$material_name"},
	//					{"specifications", "$specifications"},
	//					{"supplier", "$supplier"},
	//					{"craft", "$craft"},
	//					{"texture", "$texture"},
	//					{"process", "$process"},
	//					{"purchase_status", "$purchase_status"},
	//					{"receiving_warehouse", "$receiving_warehouse"},
	//				},
	//			},
	//		},
	//	},
	//}
	groupPip := bson.D{
		{
			"$group",
			bson.D{
				{
					"_id",
					"$material_code",
				},
			},
		},
	}
	sortPip := bson.D{
		{
			"$sort",
			bson.D{
				{sort, -1},
			},
		},
	}
	limitPip := bson.D{
		{"$limit", limit},
	}

	skipPip := bson.D{
		{
			"$skip",
			skip,
		},
	}
	return groupPip, sortPip, limitPip, skipPip
}

func GenFirstRecordDatePipeline(code string, timePara string) (bson.D, bson.D, bson.D) {
	pip := bson.D{
		{
			"$match",
			bson.D{
				{"material_code",
					bson.D{
						{"$eq", code}},
				},
				{"validate",
					bson.D{
						{"$eq", true}},
				},
			},
		},
	}
	sortPip := bson.D{
		{
			"$sort",
			bson.D{
				{timePara, -1},
			},
		},
	}
	limitPip := bson.D{
		{"$limit", 1},
	}
	return pip, sortPip, limitPip
}

func GenSupplierSearchPipeline(code string, name string, skip int, limit int) (bson.D, bson.D, bson.D) {
	pip := bson.D{
		{
			"$match",
			bson.D{
				{"code",
					bson.D{
						{"$regex", code}},
				},
				{"name",
					bson.D{
						{"$regex", name}},
				},
			},
		},
	}

	skipPip := bson.D{
		{
			"$skip",
			skip,
		},
	}

	limitPipe := bson.D{
		{
			"$limit",
			limit,
		},
	}

	return pip, skipPip, limitPipe
}

func GenSearchCodePipeline(code string, skip int, limit int) (bson.D, bson.D, bson.D, bson.D) {
	pip := bson.D{
		{
			"$match",
			bson.D{
				{"code",
					bson.D{
						{"$regex", code}},
				},
				{"status",
					bson.D{
						{"$eq", true}},
				},
			},
		},
	}

	skipPip := bson.D{
		{
			"$skip",
			skip,
		},
	}

	limitPip := bson.D{
		{
			"$limit",
			limit,
		},
	}

	countPip := bson.D{
		{
			"$count",
			"total",
		},
	}

	return pip, skipPip, limitPip, countPip
}

func GenPipeline(mc entity.WeighMultiCondition, skip int, limit int) (bl []bson.D) {

	pip0 := bson.D{

		{
			"$match",
			bson.D{
				{"material_code",
					bson.D{
						{"$regex", mc.MaterialCode}},
				},
			},
		},
	}

	bl = append(bl, pip0)

	if mc.Craft != "" {
		pip1 := bson.D{

			{
				"$match",
				bson.D{
					{"craft",
						bson.D{
							{"$eq", mc.Craft}},
					},
				},
			},
		}
		bl = append(bl, pip1)
	}

	if mc.Texture != "" {
		pip2 := bson.D{

			{
				"$match",
				bson.D{
					{"texture",
						bson.D{
							{"$eq", mc.Texture}},
					},
				},
			},
		}
		bl = append(bl, pip2)
	}

	if mc.Process != "" {
		pip3 := bson.D{

			{
				"$match",
				bson.D{
					{"process",
						bson.D{
							{"$eq", mc.Process}},
					},
				},
			},
		}
		bl = append(bl, pip3)
	}

	if mc.PurchaseStatus != "" {
		pip4 := bson.D{

			{
				"$match",
				bson.D{
					{"purchase_status",
						bson.D{
							{"$eq", mc.PurchaseStatus}},
					},
				},
			},
		}
		bl = append(bl, pip4)
	}

	if mc.Supplier != "" {
		pip5 := bson.D{

			{
				"$match",
				bson.D{
					{"supplier",
						bson.D{
							{"$regex", mc.Supplier}},
					},
				},
			},
		}
		bl = append(bl, pip5)
	}

	if mc.Validate != "" {
		pip6 := bson.D{

			{
				"$match",
				bson.D{
					{"validate",
						bson.D{
							{"$eq", mc.Validate}},
					},
				},
			},
		}
		bl = append(bl, pip6)
	}

	skipPip := bson.D{
		{
			"$skip",
			skip,
		},
	}

	bl = append(bl, skipPip)

	limitPip := bson.D{
		{"$limit", limit},
	}

	bl = append(bl, limitPip)

	fmt.Println(bl)

	//pip := bson.D{
	//
	//	{
	//		"$match",
	//		bson.D{
	//			{"material_code",
	//				bson.D{
	//					{"$regex", mc.MaterialCode}},
	//			},
	//			{"craft",
	//				bson.D{
	//					{"$eq", mc.Craft}},
	//			},
	//			{"texture",
	//				bson.D{
	//					{"$eq", mc.Texture}},
	//			},
	//			{"process",
	//				bson.D{
	//					{"$eq", mc.Process}},
	//			},
	//			{"purchase_status",
	//				bson.D{
	//					{"$eq", mc.PurchaseStatus}},
	//			},
	//		},
	//	},
	//	//{
	//	//	"$match",
	//	//	bson.D{
	//	//		{"supplier",
	//	//			bson.D{
	//	//				{"$eq", mc.Supplier}},
	//	//		},
	//	//	},
	//	//},
	//	//{
	//	//	"$match",
	//	//	bson.D{
	//	//		{"craft",
	//	//			bson.D{
	//	//				{"$eq", mc.Craft}},
	//	//		},
	//	//	},
	//	//},
	//}
	//fmt.Println("pip:   ")
	//fmt.Println(pip)

	//	projectStage := bson.D{
	//		{
	//			"$project",
	//			bson.D{
	//				{"flow_process",
	//					bson.D{
	//
	//						{"$filter",
	//							bson.D{
	//								{"input", "$flow_process"},
	//								{"as", "flow_process"},
	//								{"cond", bson.D{
	//									{"$eq", ["$$flow_process.weigh_stage", ""]},
	//								}},
	//							},
	//						},
	//
	//,					},
	//				},
	//			},
	//		},
	//	}

	return
}

func GenCalPipeline(mc entity.CalMultiCondition, skip int, limit int) (bl []bson.D) {

	pip0 := bson.D{

		{
			"$match",
			bson.D{
				{"material_code",
					bson.D{
						{"$regex", mc.MaterialCode}},
				},
			},
		},
	}

	bl = append(bl, pip0)

	if mc.Craft != "" {
		pip1 := bson.D{

			{
				"$match",
				bson.D{
					{"craft",
						bson.D{
							{"$eq", mc.Craft}},
					},
				},
			},
		}
		bl = append(bl, pip1)
	}

	if mc.Texture != "" {
		pip2 := bson.D{

			{
				"$match",
				bson.D{
					{"texture",
						bson.D{
							{"$eq", mc.Texture}},
					},
				},
			},
		}
		bl = append(bl, pip2)
	}

	if mc.Process != "" {
		pip3 := bson.D{

			{
				"$match",
				bson.D{
					{"process",
						bson.D{
							{"$eq", mc.Process}},
					},
				},
			},
		}
		bl = append(bl, pip3)
	}

	if mc.PurchaseStatus != "" {
		pip4 := bson.D{

			{
				"$match",
				bson.D{
					{"purchase_status",
						bson.D{
							{"$eq", mc.PurchaseStatus}},
					},
				},
			},
		}
		bl = append(bl, pip4)
	}

	if mc.Validate != "" {
		pip6 := bson.D{

			{
				"$match",
				bson.D{
					{"validate",
						bson.D{
							{"$eq", mc.Validate}},
					},
				},
			},
		}
		bl = append(bl, pip6)
	}

	skipPip := bson.D{
		{
			"$skip",
			skip,
		},
	}

	bl = append(bl, skipPip)

	limitPip := bson.D{
		{"$limit", limit},
	}

	bl = append(bl, limitPip)

	fmt.Println(bl)

	return
}

func DataAggregate(m *mongodb.Mgo, mc entity.WeighMultiCondition, resultChan chan bson.M, wg *sync.WaitGroup, skip int, limit int) {
	//fmt.Println("DataAggregate")
	matchStage := GenPipeline(mc, skip, limit)
	fmt.Println(matchStage)
	opts := options.Aggregate().SetMaxTime(15 * time.Second)
	//cursor := m.FindAggregation(mongo.Pipeline{matchStage}, opts)
	cursor := m.FindAggregation(matchStage, opts)

	// 打印文档内容
	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		//utils.GetLogger().Info(fmt.Sprintf("dataAggregate error: %v", err))
		zap.L().Info("dataAggregate", zap.String("dataAggregate error", fmt.Sprintf("--------->%s", zap.Error(err).String)))
	}
	for _, result := range results {
		fmt.Println(result)
		resultChan <- result
	}
	wg.Done()
}
