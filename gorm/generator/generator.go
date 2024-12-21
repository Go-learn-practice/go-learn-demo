package generator

import (
	"fmt"
	"gorm.io/gen"
	"gorm/config"
	"gorm/model"
)

func Generate() {
	// 创建生成器实例
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dao/query",
		ModelPkgPath: "model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式
	})

	// 设置数据库连接
	g.UseDB(config.DB)

	//g.ApplyBasic(g.GenerateAllTable()...)

	// 指定需要生成代码的表模型 g.GenerateModel("users") 表示数据库已经存在users表
	g.ApplyBasic(g.GenerateModel("users"), model.Students{})
	g.Execute()

	fmt.Println("generate success")
}
