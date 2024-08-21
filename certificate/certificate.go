package certificate

type CertificateRepository struct {
}

type CertificateService struct {
	*CertificateRepository
}

func NewCertificateRepository() *CertificateRepository {
	return &CertificateRepository{}
}

func NewCertificateService(repository *CertificateRepository) *CertificateService {
	return &CertificateService{
		CertificateRepository: repository,
	}
}
