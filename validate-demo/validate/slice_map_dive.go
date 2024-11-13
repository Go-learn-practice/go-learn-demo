package validate

// 切片与集合内元素格式验证

type sliceStruct struct {
	OpCode int    `v:"eq=1|eq=2"`
	Op     string `v:"required"`
}

func SliceValidate() {
	v := validate
	slice1 := []string{"12345", "67890", "1234567890"}
	var err error
	err = v.Var(slice1, "gte=3,dive,required,gte=5,lte=10,number")
	outRes("slice", &err)

	// TODO validate 有问题
	slice2 := [][]string{
		{"12345", "67890", "1234567890"},
		{"12345", "67890", "1234567890"},
		{"12345", "67890", "1234567890"},
	}
	err = v.Var(slice2, "gte=3,dive,gte3,dive,required,gte=5,lte=10,number")
	outRes("slice2", &err)

	slice3 := []*sliceStruct{
		{
			OpCode: 1,
			Op:     "切片操作",
		},
		{
			OpCode: 2,
			Op:     "切片操作",
		},
		{
			OpCode: 3,
			Op:     "切片操作",
		},
	}
	err = v.Var(slice3, "gte=2,dive")
	outRes("slice3", &err)
}

func MapValidate() {
	v := validate
	var err error
	mp1 := map[string]string{
		"A": "12345",
		"B": "67890",
		"C": "12345",
	}
	err = v.Var(mp1, "gte3,dive,keys,len=1,alpha,endkeys,required,gte=5,lte=10,number")
	outRes("map", &err)

	mp2 := map[string]map[string]string{
		"A": {
			"Aa": "12345",
			"Bb": "67890",
			"Cc": "12345",
		},
		"B": {
			"Aa": "12345",
			"Bb": "67890",
			"Cc": "12345",
		},
	}
	err = v.Var(mp2, "gte2,dive,keys,len=1,alpha,endkeys,required,gte3,dive,keys,len=1,alpha,endkeys,required,gte=5,lte=10,number")
	outRes("map2", &err)
}
