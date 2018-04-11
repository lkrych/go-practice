package main

var opts struct {
	Name string `short:"n" long:"name default: "World" description:"A name to say hello to."`
	Spanish bool `short:"s" long:"spanish description:"Use Spanish Language"`
}

func main {
	flags.parse(&opts)

	if opts.Spanish == true {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
