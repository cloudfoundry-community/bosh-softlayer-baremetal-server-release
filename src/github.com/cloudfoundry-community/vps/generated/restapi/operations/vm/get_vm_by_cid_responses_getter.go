package vm

import (
	"github.com/cloudfoundry-community/vps/generated/models"
)

func (o *GetVMByCidOK) GetPayload() *models.VMResponse {
	return o.Payload
}

func (o *GetVMByCidNotFound) GetStatusCode() int{
	return 404
}

func (o *GetVMByCidDefault) GetStatusCode() int {
	return o._statusCode
}

func (o *GetVMByCidDefault) GetPayload() *models.Error{
	return o.Payload
}
