package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	authDelivery "github.com/ilhampraset/testcasebe/auth/delivery/http"
	userrepo "github.com/ilhampraset/testcasebe/auth/repository/mysql"
	authUc "github.com/ilhampraset/testcasebe/auth/usecase"
	"github.com/ilhampraset/testcasebe/auth/utils"
	merchDelivery "github.com/ilhampraset/testcasebe/merchant/delivery/http"
	merchrepo "github.com/ilhampraset/testcasebe/merchant/repository/mysql"
	merchUc "github.com/ilhampraset/testcasebe/merchant/usecase"

	"github.com/ilhampraset/testcasebe/middleware"
	outletDelivery "github.com/ilhampraset/testcasebe/outlet/delivery/http"
	outletrepo "github.com/ilhampraset/testcasebe/outlet/repository/mysql"
	outlethUc "github.com/ilhampraset/testcasebe/outlet/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/testcase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	var jwt utils.Token = utils.JWTAuth()
	urepo := userrepo.NewAuthRepository(db)
	authUC := authUc.NewAuthUseCase(urepo)

	merrepo := merchrepo.NewMerchantRepository(db)
	merUC := merchUc.NewMerchantUseCase(merrepo)

	outrepo := outletrepo.NewOutletRepository(db)
	outUC := outlethUc.NewOutletUseCase(outrepo)

	outletHandler := outletDelivery.NewOutletHandler(outUC, jwt)
	authHandler := authDelivery.NewAuthHandler(authUC, jwt)
	merchantHandler := merchDelivery.NewMerchantHandler(merUC, jwt)
	g := gin.Default()

	g.POST("/login", authHandler.Login)
	g.Use(middleware.AuthorizeJWT())
	{
		g.GET("/me", authHandler.Me)
		g.GET("/merchant-report-monthly", merchantHandler.MonthlyReport)
		g.GET("/outlet-report-monthly", outletHandler.MonthlyReport)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 3)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
