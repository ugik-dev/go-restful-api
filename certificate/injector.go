// go:build wireinject
//go:build wireinject
// +build wireinject

package certificate

import "github.com/google/wire"

func InitializedCertService() *CertificateService {
	wire.Build(NewCertificateRepository, NewCertificateService)
	return nil
}

/**
untuk melakuka auto generate :
wire gen namapackage
ex:  wire gen github.com/ugik-dev/go-restful-api.git/certificate
atau bisa masuk dir certificate dlu
ex: cd certificate baru jalankan wire saja di terminal
*/
