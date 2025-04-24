package subpackage

import (
	"fmt"

	json "github.com/goccy/go-json"
	sandbox "github.com/yookoala/go-subpackage-sandbox"
)

type T struct {
	X int
	U *U
}

type U struct {
	T *T
}

func jsonTest() string {
	b, err := json.Marshal(&T{
		X: 1,
		U: &U{
			T: &T{
				X: 2,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return string(b)
}

func Hello() string {
	fmt.Println(jsonTest())
	return sandbox.Hello() + ", " + jsonTest()
}
