package service

import "github.com/Seunghoon-Oh/cloud-ml-foo-manager/data"

func GetFoos() []string {
	return data.GetRedisData()
}

func CreateFoo() string {
	// TODO: Create Kubernetes Object
	return data.CreateRedisData()
}
