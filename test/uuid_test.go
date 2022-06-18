package test

import (
	"cloud_disk/core/helper"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	uuid := helper.UUID()
	println(uuid)
}
