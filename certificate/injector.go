// go:build wireinject
//go:build wireinject
// +build wireinject

package certificate

import "github.com/google/wire"

func InitializedCertService(isErrorNotSame bool) (*CertificateService, error) {
	wire.Build(NewCertificateRepository, NewCertificateService)
	return nil, nil
}

/**
pada isBool dibaca berdasarkan tipe data bukan nama paremter
untuk melakuka auto generate :
wire gen namapackage
ex:  wire gen github.com/ugik-dev/go-restful-api.git/certificate
atau bisa masuk dir certificate dlu
ex: cd certificate baru jalankan wire saja di terminal
*/
