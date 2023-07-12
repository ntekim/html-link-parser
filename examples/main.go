package main

import (
	"fmt"
	"github.com/ntekim/html-link-parser/api"
	"strings"
)

var htmlStringExample = `
	<html>
	<body>
		<h1>Hey Man!</h1>
		<p> ejdjdjdjd</>
		<a href="/other-page"> A link to other page</a>
		<a href="/none">New link </>
	</body>
	</html>
`

func main() {
	r := strings.NewReader(htmlStringExample)

	link, err := api.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", link)
}
