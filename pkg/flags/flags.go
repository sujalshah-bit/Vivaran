package flags

type Spec struct {
	Short   string
	Name    string
	Desc    string
	Example string
}

var Supported = []Spec{
	{
		Short:   "c",
		Name:    "Size",
		Desc:    "Outputs the file size in bytes",
		Example: "go run ./cmd/main.go -c test.txt",
	},
	{
		Short:   "l",
		Name:    "Lines",
		Desc:    "Counts the number of lines",
		Example: "go run ./cmd/main.go -l test.txt",
	},
	{
		Short:   "w",
		Name:    "Words",
		Desc:    "Counts words separated by whitespace",
		Example: "go run ./cmd/main.go -w test.txt",
	},
	{
		Short:   "m",
		Name:    "Characters",
		Desc:    "Counts Unicode characters (multibyte-safe)",
		Example: "go run ./cmd/main.go -m test.txt",
	},
	{
		Short:   "bs",
		Name:    "Buffer size",
		Desc:    "Set buffer size",
		Example: "go run ./cmd/main.go -bs test.txt",
	},
}
