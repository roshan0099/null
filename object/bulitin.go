package object

type BuiltnCondition func(args []Object) Object

var Builtin = map[string]Object{

	"nout": &Wrapper{

		WrapperFunc: func(args []Object) Object {

			noutStore := &Nout{}

			for _, val := range args {

				noutStore.Statements = append(noutStore.Statements, val)
			}

			return noutStore

		},
	},

	"nin": &Wrapper{

		WrapperFunc: func(args []Object) Object {

			noutStore := &Nout{}
			return noutStore

		},
	},
}
