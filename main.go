package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// LikeRequest struct untuk permintaan like
type LikeRequest struct {
	AccountID string `json:"account_id"`
}

// Fungsi untuk mengirimkan permintaan like ke akun Warpcast
func likeAccount(accountID string, token string) error {
	url := "https://api.warpcast.com/like" // Gantilah dengan URL endpoint yang benar

	// Membuat payload permintaan
	payload := LikeRequest{AccountID: accountID}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("gagal mengubah payload ke JSON: %v", err)
	}

	// Membuat HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("gagal membuat request baru: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Mengirimkan request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("gagal mengirimkan request: %v", err)
	}
	defer resp.Body.Close()

	// Memeriksa respons dari server
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server mengembalikan status tidak OK: %v", resp.Status)
	}

	return nil
}

// Fungsi untuk mengekstrak AccountID dari URL Warpcast
func extractAccountID(url string) (string, error) {
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return "", fmt.Errorf("URL tidak valid")
	}
	return parts[len(parts)-1], nil
}

func main() {
	// Token Bearer untuk autentikasi
	token := "MK-wSj12+UJt2K08rMSoIyx/oDnETFs1pEQ+OxJt3NSeAwCdiFPZ46vFzb6eEtr7zSSaKVbgGX9X5Cxm78ivCwryw==" // Gantilah dengan token akses yang benar

	// Daftar URL Warpcast yang akan di-like
	urlList := []string{
		"
"https://warpcast.com/zeroyogi/0xa4d4dd29

"https://warpcast.com/samyangmercon/0xb52b63ab

"https://warpcast.com/rizkiiaprmn/0x25c61213

"https://warpcast.com/zeeejkt48/0xf25a870b

"https://warpcast.com/arastra/0x1d1f6a80

"https://warpcast.com/tjakrahardja/0xa57e6e8e

"https://warpcast.com/masdim/0x62314d67

"https://warpcast.com/tempek123/0x5863e48b

"https://warpcast.com/mprutz/0xed265b42

"https://warpcast.com/mrwanjrwow/0x0a46b3cf",
		"https://warpcast.com/account2",
		"https://warpcast.com/account3",
	}

	// Melakukan like untuk setiap URL dalam daftar
	for _, url := range urlList {
		accountID, err := extractAccountID(url)
		if err != nil {
			log.Printf("Gagal mengekstrak AccountID dari URL %s: %v\n", url, err)
			continue
		}
		err = likeAccount(accountID, token)
		if err != nil {
			log.Printf("Gagal memberi like pada akun %s: %v\n", accountID, err)
		} else {
			log.Printf("Berhasil memberi like pada akun %s\n", accountID)
		}
		time.Sleep(2 * time.Second) // Jeda waktu 2 detik antara permintaan like
	}
}
