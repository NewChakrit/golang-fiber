package logs

import "go.uber.org/zap"

var Log *zap.Logger

func init () {
	// Log, _=zap.NewDevelopment() // console
	Log, _= zap.NewProduction() 
}