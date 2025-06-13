package instrucciones

import "github.com/antlr4-go/antlr/v4"

const (
	ExternNameParam = iota
	ExternEqualInnerParam
	PositionalParam
)

type Parametros struct {
	ExternName      string
	InnerName       string
	Type            string
	PassByReference bool
	Token           antlr.Token
}

// 3 types of Parametrosantlr
// 1. extern and inner name declared 			[args are specified with the extern name]
// 2. extern name = "_", inner name declared 	[args are specified with the order]
// 3. extern name = "", inner name declared	[args are specified with the inner name]

func (p *Parametros) ParamType() int {
	if p.ExternName != "" {
		if p.ExternName == "_" {
			return PositionalParam
		} else {
			return ExternNameParam
		}
	} else {
		return ExternEqualInnerParam
	}
}
