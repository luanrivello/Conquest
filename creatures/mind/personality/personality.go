package personality

type Personality struct {
	E int8
	I int8

	S int8
	N int8

	T int8
	F int8

	J int8
	P int8
}

var ISFP Personality = Personality{I:1,S:1,F:1,P:1}

func GetISFP () *Personality {
	return &ISFP
}
