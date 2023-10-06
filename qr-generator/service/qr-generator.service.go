package service

import (
    "github.com/skip2/go-qrcode"
)

func GenerateQRPNG(data string) ([]byte, error) {
    qr, err := qrcode.New(data, qrcode.Medium)
    if err != nil {
        return nil, err
    }
    
    qrPNG, err := qr.PNG(256)
    if err != nil {
        return nil, err
    }
    
    return qrPNG, nil
}