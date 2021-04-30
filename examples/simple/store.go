package main

import "github.com/bytedisciple/state"

var myGlobalInt state.Int
var myOtherGlobalInt state.Int

func init(){
	myGlobalInt = state.NewInt("myGlobalInt")
	myOtherGlobalInt = state.NewInt("myOtherGlobalInt")
}
