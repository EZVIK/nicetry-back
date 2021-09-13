package main

import (
	"context"
	"fmt"
)

/**
 * @Description:
	new context to create a new request, setting timeout
*/
//func main() {
//
//
//	req, err := http.NewRequest("GET", "https://eddycjy.com", nil)
//	if err != nil {
//		fmt.Println("http.NewRequest err: %+v", err)
//		return
//	}
//
//	// get parent context , adding timeout ,
//	ctx, cancel := context.WithTimeout(req.Context(), time.Microsecond * 50 )
//
//	//
//	defer cancel()
//
//	// create a new req with a context
//	req = req.WithContext(ctx)
//
//	// exec request
//	resp, err := http.DefaultClient.Do(req)
//
//	if err != nil {
//		fmt.Println("http.DefaultClient.Do err: %+v", err)
//		return
//	}
//
//	defer resp.Body.Close()
//}
type SHIT struct {
	ID   int
	NAME string
	AGE  int
	DESC string
}

func (this *SHIT) String() string {
	return fmt.Sprint(this.ID, this.NAME, this.AGE, this.DESC)
}

func main() {

	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("🧠")
	s := SHIT{1, "JACK", 29, "DO NOT CLOSE TO ME"}

	ctx := context.WithValue(context.Background(), k, s)

	f(ctx, k)

	f(ctx, "🐟")
}
