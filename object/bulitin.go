package object

type BuiltnCondition func(args []Object) Object

var Builtin = map[string]Object{

	"np": &Wrapper{
		Name: "np",
		WrapperFunc: func(args []Object) Object {

			noutStore := &Nout{}

			for _, val := range args {

				noutStore.Statements = append(noutStore.Statements, val)
			}

			return noutStore

		},
	},

	"ns": &Wrapper{
		Name: "ns",
		WrapperFunc: func(args []Object) Object {

			noutStore := &Nout{}
			return noutStore

		},
	},

	"ni": &Wrapper{
		Name: "ni",
		WrapperFunc: func(args []Object) Object {

			noutStore := &Nout{}
			return noutStore

		},
	},
}
