package cmd

import (
	"basket-discount-api/internal/discount"
	"basket-discount-api/internal/discount/controller"
	"basket-discount-api/internal/discount/repositories"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(apiCmd)
	var Port string
	apiCmd.Flags().StringVarP(&Port, "port", "p", "5000", "Service Port")
	context := echo.New()
	context.Debug = false
	context.HideBanner = true
	context.HidePort = !true
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println("Basket-Discount-API is running")

		conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			AppConfig.DiscountDbSettings.Host,
			AppConfig.DiscountDbSettings.User,
			AppConfig.DiscountDbSettings.Pass,
			AppConfig.DiscountDbSettings.DbName,
			AppConfig.DiscountDbSettings.Port,
		)
		db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
		if err != nil {
			return err
		}
		campaignRepo := repositories.NewCampaignRepo(db)
		discountService := discount.NewService(campaignRepo)
		discountController := controller.NewController(discountService)
		controller.MakeDiscountHandler(context, discountController)
		context.GET("/swagger/*", echoSwagger.WrapHandler)
		go func() {
			if err := context.Start(fmt.Sprintf(":%s", "5000")); err != nil {
				//log.Logger.Fatalf("Failed to shutting down the server. Error:%v", err)
			}
		}()
		return nil
	}
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Basket-Discount API",
}
