package vm

import (
	"github.com/cloudfoundry-community/vps/generated/models"
)

func (o *AddVMOK) GetPayload() string {
	return o.Payload
}

func (o *AddVMDefault) GetStatusCode() int {
	return o._statusCode
}

func (o *AddVMDefault) GetPayload() *models.Error{
	return o.Payload
}
