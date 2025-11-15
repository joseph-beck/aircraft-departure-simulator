package main

import (
	"fmt"
	"net/rpc"

	"github.com/gin-gonic/gin"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype/aircraftconstants"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/pkg/carrier/fleet/aircraft/subtype/maxweight"
)

func main() {
	g := gin.Default()

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	g.GET("/v1/health/ping", func(ctx *gin.Context) {
		var healthResp string
		err = client.Call("ManageService.Ping", "pong", &healthResp)
		if err != nil {
			ctx.Status(500)
			return
		}

		ctx.JSON(200, gin.H{
			"message": healthResp,
		})
	})

	g.GET("/v1/manage/ac", func(ctx *gin.Context) {
		var acResp []aircraftconstants.AircraftConstant
		err = client.Call("ManageService.GetAircraftConstants", "args", &acResp)
		if err != nil {
			fmt.Println(err)
			ctx.Status(500)
			return
		}

		ctx.JSON(200, acResp)
	})

	g.GET("/v1/manage/mw", func(ctx *gin.Context) {
		var mwResp []maxweight.MaxWeight
		err = client.Call("ManageService.GetMaximumWeights", "args", &mwResp)
		if err != nil {
			fmt.Println(err)
			ctx.Status(500)
			return
		}

		ctx.JSON(200, mwResp)
	})

	g.Run(":8080")
}
