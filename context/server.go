package main

import (
	"net/http"
	"time"
	"log"
	"context"
	"fmt"
	"errors"
)



func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

		go func() {
			var result string

			for _, c := range s.response {
				select {
				case <-ctx.Done():
					log.Println("spy store got cancelled")
					return
				
				default:
					time.Sleep(10 * time.Millisecond)
					result += string(c)
				}

			}

			data <- result

			
		}()

		select {

		case <- ctx.Done():
			return "", ctx.Err()
		
		case res := <-data:
			return res, nil
		}

}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}


type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		fmt.Fprint(w, data)
	
		if err != nil {
			return errors.New("Got an error but didn't need one")
		}

		fmt.Fprint(w, data)
	}
}

