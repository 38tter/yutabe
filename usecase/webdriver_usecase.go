package usecase

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sclevine/agouti"
)

const (
	TargetUrl      = "https://sm.rakuten.co.jp/"
	TargetFormID   = "topSearchKeyword"
	TargetButtonID = "topSearchButton"
)

func UseWebDriver(ctx *gin.Context) {
	items, _ := ctx.GetQueryArray("item")
	if len(items) == 0 {
		fmt.Errorf("validation error: no query params")
		return
	}
	fmt.Printf("item is %v\n", items[0])

	agoutiDriver := agouti.ChromeDriver()
	err := agoutiDriver.Start()
	if err != nil {
		fmt.Errorf("failed to start driver: %s\n", err)
		return
	}
	defer func() {
		err = agoutiDriver.Stop()
		if err != nil {
			fmt.Errorf("failed to start driver: %s\n", err)
			return
		}
	}()
	page, _ := agoutiDriver.NewPage()

	if err := page.CancelPopup(); err != nil {
		fmt.Errorf("failed to cancel popup %s\n", err)
	}
	if err = page.SetImplicitWait(20); err != nil {
		fmt.Errorf("failed to set implicit wait %s\n", err)
	}

	if err = page.Navigate(TargetUrl); err != nil {
		fmt.Errorf("failed to navigate site: %s\n", err)
	}
	searchForm := page.FindByID(TargetFormID)
	err = searchForm.Fill(items[0])
	if err != nil {
		fmt.Errorf("failed to fill search form: %s\n", err)
	}

	btn := page.FindByID(TargetButtonID)
	err = btn.Submit()
	if err != nil {
		log.Fatalf("failed to submit button: %s\n", err)
	}

	page.Screenshot("screenshot.png")
}
