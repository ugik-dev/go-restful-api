package certificate

import (
	"fmt"
	"testing"
)

func TestCallCertificateService(t *testing.T) {
	certificateService := InitializedCertService()
	fmt.Print(certificateService.CertificateRepository)
}
