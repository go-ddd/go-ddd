package do

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --template ./template ./schema
// //go:generate go run entgo.io/contrib/entproto/cmd/entproto -path ./schema
//go:generate go run -mod=mod github.com/galaxyobe/gen/cmd/getter-gen -i . -o .
//go:generate go run -mod=mod github.com/galaxyobe/gen/cmd/setter-gen -i . -o .
