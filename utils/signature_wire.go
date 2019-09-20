//+build wireinject

package utils

import (
	"github.com/google/wire"
)

func SignatureInstance() Signature {
	wire.Build(ProvideSignature)
	return &signatureImpl{}
}
