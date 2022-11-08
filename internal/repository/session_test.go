package repository

import (
	"fmt"
	"github.com/fastid/fastid/internal/config"
	"testing"
)

func TestSession(t *testing.T) {
	t.Parallel()
	//ctx := context.Background()

	cfg, err := config.NewConfig("../../configs/config.yml")
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	fmt.Println(cfg)

	//repo := NewSessionsRepository(cfg, mockPool)

	//err = repo.Create()
	//if err != nil {
	//	t.Fatalf("error %s", err.Error())
	//}

	//err = repo.Create()
	//if err != nil {
	//	t.Fatalf("error %s", err.Error())
	//}

}
