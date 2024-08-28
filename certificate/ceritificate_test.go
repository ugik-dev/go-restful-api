package certificate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallCertificateServiceSuccess(t *testing.T) {
	certificateService, err := InitializedCertService(false)
	// helper.PanicIfError(err)
	// fmt.Print(err)
	// fmt.Print(certificateService.CertificateRepository)
	assert.NotNil(t, certificateService)
	assert.Nil(t, err)
}
func TestCallCertificateServiceError(t *testing.T) {
	certificateService, err := InitializedCertService(true)
	// helper.PanicIfError(err)
	assert.NotNil(t, err)
	assert.Nil(t, certificateService)
}
