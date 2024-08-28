package certificate

import "errors"

type CertificateRepository struct {
	Error bool
}

type CertificateService struct {
	*CertificateRepository
}

func NewCertificateRepository(isError bool) *CertificateRepository {
	return &CertificateRepository{
		Error: isError,
	}
}

func NewCertificateService(repository *CertificateRepository) (*CertificateService, error) {
	if repository.Error {
		return nil, errors.New("Gagal membuat service")
	} else {
		return &CertificateService{
			CertificateRepository: repository,
		}, nil
	}
}
