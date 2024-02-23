package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/actanonvebra/gopixdeneme/config"
)

var apiEndpoint = "https://api.imgur.com/3/image"

func UploadAndDisplayLink(imagePath string) error {
	// Config yapısını config paketindeki fonksiyon kullanarak al
	configValues, err := config.ReadConfigFile()
	if err != nil {
		return fmt.Errorf("Config okuma hatası: %v", err)
	}

	// Resmi yükleme işlemi
	link, err := uploadImageToImgur(apiEndpoint, imagePath, configValues.ClientID, configValues.ClientSecret)
	if err != nil {
		return fmt.Errorf("Resim yükleme hatası: %v", err)
	}

	fmt.Println("Imgur API Yanıtı - Resim Linki:", link)

	return nil
}

func uploadImageToImgur(apiEndpoint, imagePath, clientID, clientSecret string) error {
	// Resim dosyasını aç
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Resim dosyasını okuma
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", imagePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	// Diğer isteğin parametreleri
	writer.WriteField("client_id", clientID)
	writer.WriteField("client_secret", clientSecret)

	// İsteği oluştur
	req, err := http.NewRequest("POST", apiEndpoint, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// İsteği gönder
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// İsteğin sonucunu oku
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("API Yanıtı:", string(responseBody))

	return nil
}
