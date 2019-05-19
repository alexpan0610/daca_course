package tx

type Tx struct {
	Inputs   []Input
	Outputs  []Output
	Programs []interface{}
}
