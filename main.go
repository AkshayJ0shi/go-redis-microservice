package main
import (
	"fmt"
	"context"
	"github.com/AkshayJ0shi/go-redis-microservice/application"
)

func main(){
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start", err)
	}
}
