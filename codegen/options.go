package codegen

import "errors"

type Options struct {
	Source string
	Dest   string
	Types  []string
	Pkg    string
}

func (o *Options) Validate() error {
	if o == nil {
		return errors.New("options is nil")
	}
	if o.Source == "" {
		return errors.New("source was empty")
	}
	if len(o.Types) == 0 {
		return errors.New("types was empty")
	}
	return nil
}
