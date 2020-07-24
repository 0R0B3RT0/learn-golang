package b3

type RequestStruct struct {
	Name         string //ITUB3
	FriendlyName string //ITUB3
	Columns      []Column
	Values       []Values
}

type Column struct {
	Name            string //TickerSymbol
	FriendlyName    string //Papel
	FriendlyNamePt  string //Papel
	FriendlyNameEn  string //Symbol
	Type            string //5
	ColumnAlignment string //1
	ValueAlignment  string //2
}

type Values []interface{}

// type Values struct {
// 	name string //ITUB3
// 	quantity int //100
// 	price value //24.96
// 	bulk int //8910
// 	date string //"2020-07-24"
// 	time string //"13:32:39"
// }
