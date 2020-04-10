package test

import (
	"fmt"
	"testing"
	"../common"
)
func TestGCD(t *testing.T){
	if(common.GCD(4,6) != 2){
		fmt.Println(common.GCD(4,6))
		t.Errorf("Failed GCD test")
	}
}
func TestVectorGCD(t *testing.T)  {
	tv := []int{4,10,16,14}
	if(common.VectorGCD(tv) != 2){
		t.Errorf("Failed VGCD test")
	}
}