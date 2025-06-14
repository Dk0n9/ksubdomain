package outputter

import (
	"github.com/Dk0n9/ksubdomain/v2/pkg/runner/result"
)

type Output interface {
	WriteDomainResult(domain result.Result) error
	Close() error
}
