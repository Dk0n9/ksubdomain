package output

import (
	"strings"

	"github.com/Dk0n9/ksubdomain/v2/pkg/core/gologger"
	"github.com/Dk0n9/ksubdomain/v2/pkg/runner/result"
)

type ScreenOutputNoWidth struct {
	silent bool
}

func NewScreenOutputNoWidth(silent bool) (*ScreenOutputNoWidth, error) {
	return &ScreenOutputNoWidth{silent: silent}, nil
}
func (s *ScreenOutputNoWidth) WriteDomainResult(domain result.Result) error {
	var msg string
	var domains []string = []string{domain.Subdomain}
	for _, item := range domain.Answers {
		domains = append(domains, item.Value)
	}
	msg = strings.Join(domains, " => ")
	if !s.silent {
		gologger.Silentf("%s\n", msg)
	} else {
		gologger.Silentf("%s\n", domain.Subdomain)
	}
	return nil
}
func (s *ScreenOutputNoWidth) Close() error {
	return nil
}
