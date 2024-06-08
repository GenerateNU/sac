package auth

import "crypto/rand"

func GenerateOTP(length int) (*string, error) {
	digits := "0123456789"
	otp := make([]byte, length)
	if _, err := rand.Read(otp); err != nil {
		return nil, err
	}

	for i := 0; i < length; i++ {
		otp[i] = digits[int(otp[i])%10]
	}

	outOtp := string(otp)

	return &outOtp, nil
}
