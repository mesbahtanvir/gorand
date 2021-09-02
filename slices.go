package gorand

import (
	"math/rand"
)

func SliceInt63(len int) []int64 {
	res := make([]int64, len)
	for i := 0; i < len; i++ {
		res[i] = rand.Int63()
	}
	return res
}

func SliceInt63n(len int, n int64) []int64 {
	res := make([]int64, len)
	for i := 0; i < len; i++ {
		res[i] = rand.Int63n(n)
	}
	return res
}

func SliceInt(len int) []int {
	res := make([]int, len)
	for i := 0; i < len; i++ {
		res[i] = rand.Int()
	}
	return res
}

func SliceIntn(len, n int) []int {
	res := make([]int, len)
	for i := 0; i < len; i++ {
		res[i] = rand.Intn(n)
	}
	return res
}

func SliceInt31(len int) []int32 {
	res := make([]int32, len)
	for i := 0; i < len; i++ {
		res[i] = rand.Int31()
	}
	return res
}

func SliceInt31n(len int, n int32) []int32 {
	res := make([]int32, len)
	for i := 0; i < len; i++ {
		res[i] = rand.Int31n(n)
	}
	return res
}

func SliceFloat64(len int) []float64 {
	res := make([]float64, len)
	for i := 0; i < len; i++ {
		res[i] = rand.Float64()
	}
	return res
}

func SliceFloat32(len int) []float32 {
	res := make([]float32, len)
	for i := 0; i < len; i++ {
		res[i] = rand.Float32()
	}
	return res
}
