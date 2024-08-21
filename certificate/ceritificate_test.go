package certificate

import (
	"fmt"
	"testing"
)

func TestCallCertificateService(t *testing.T) {
	certificateService, err := InitializedCertService()
	// helper.PanicIfError(err)
	fmt.Print(err)
	fmt.Print(certificateService)
}
