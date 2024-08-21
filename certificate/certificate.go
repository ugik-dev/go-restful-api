package certificate

import "errors"

type CertificateRepository struct {
	Error bool
}

type CertificateService struct {
	*CertificateRepository
}

func NewCertificateRepository() *CertificateRepository {
	return &CertificateRepository{
		Error: true,
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
