package controller

import "iris-blog/common"

type ExampleController struct {
}

func (self *ExampleController) Get() common.ResponseJson {
	ret := common.ResponseJson{
		"title": "this  is  an example",
	}

	return ret
}
