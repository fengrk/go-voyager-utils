package utils

import (
    "github.com/google/logger"
    "os"
)

const logPath = "example.log"

func initLogger() {
    lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    if err != nil {
        logger.Fatalf("Failed to open log file: %v", err)
        os.Exit(1)
    }
    defer lf.Close()
    
    defer logger.Init("LoggerExample", true, true, lf).Close()
    
    logger.Info("init logger!")
    
}
