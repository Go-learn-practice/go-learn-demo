package generator

import (
	"fmt"
	"gorm.io/gen"
	"gorm/config"
)

func Generate() {
	// 创建生成器实例
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dao",
		ModelPkgPath: "./dao/model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式
	})

	g.UseDB(config.DB)
	//g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(g.GenerateModel("users"))
	g.Execute()

	fmt.Println("generate success")
}
