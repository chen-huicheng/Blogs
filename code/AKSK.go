package code
import (
    "crypto/hmac"
    "crypto/sha256"
    "fmt"
    "time"
)

func sha256HMAC(key []byte, data []byte) []byte {
    mac := hmac.New(sha256.New, key)
    mac.Write(data)
    return []byte(fmt.Sprintf("%x", mac.Sum(nil)))
}

// Get请求
func SignGet(ak string, sk string, queryString []byte) string {
    expiration := 1800  # 有效时间, 单位是s, 根据自己的业务的实际情况调整
    signKeyInfo := fmt.Sprintf("%s/%s/%d/%d", "x-janus-auth", ak, time.Now().Unix(), expiration)
    signKey := sha256HMAC([]byte(sk), []byte(signKeyInfo))
    signResult := sha256HMAC(signKey, queryString)
    return fmt.Sprintf("%v/%v", signKeyInfo, string(signResult))
}

// Post请求
func SignPost(ak string, sk string, body []byte) string {
    expiration := 1800  # 有效时间, 单位是s, 根据自己的业务的实际情况调整
    signKeyInfo := fmt.Sprintf("%s/%s/%d/%d", "x-janus-auth", ak, time.Now().Unix(), expiration)
    signKey := sha256HMAC([]byte(sk), []byte(signKeyInfo))
    signResult := sha256HMAC(signKey, body)
    return fmt.Sprintf("%v/%v", signKeyInfo, string(signResult))
}