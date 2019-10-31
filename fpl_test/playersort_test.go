package fpl_test

import (
	"log"
	"os"

	"testing"

	"github.com/joho/godotenv"
	fpl "github.com/jonnoking/vidukavindaloo-fpl"
	"github.com/jonnoking/vidukavindaloo-fpl/config"
)

func main() {
	FPL := initFPL()
	log.Println(FPL.Bootstrap.Events.GetCurrentEvent().Name)

	FPL.Bootstrap.Players.SortByTotalPoints()

	for i := 0; i <= 10; i++ {
		log.Printf("%s - %d \n", FPL.Bootstrap.Players.Players[i].GetFullName(), FPL.Bootstrap.Players.Players[i].TotalPoints)
		//log.Printf("%s - %d \n", np[i].GetFullName(), np[i].TotalPoints)
	}

}

func initFPL() *fpl.FPL {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	fplConfig := config.New(getEnv("FPL_USER", ""), getEnv("FPL_PASSWORD", ""), 8, "", "", "")
	FPL := fpl.New(fplConfig)
	FPL.LoadBoostrapLive()

	return FPL
}

func TestPlayerSort(t *testing.T) {

	FPL := initFPL()
	log.Println(FPL.Bootstrap.Events.GetCurrentEvent().Name)

	FPL.Bootstrap.Players.SortByTotalPoints()

	for i := 0; i <= 10; i++ {
		log.Printf("[%d] %s - %d \n", FPL.Bootstrap.Players.Players[i].ID, FPL.Bootstrap.Players.Players[i].GetFullName(), FPL.Bootstrap.Players.Players[i].TotalPoints)
		//log.Printf("%s - %d \n", np[i].GetFullName(), np[i].TotalPoints)
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
