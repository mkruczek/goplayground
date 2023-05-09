package something_like_httpserver

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"temporary/gin"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
)

type UserAuthentication struct {
	UserRepo UserRepository
}

func (ua *DeviceAuthentication) AuthenticateMiddleware() gin.HandlerFunc {
	return func(c context.Context) {

		//TODO
		tokenId := "08b15c6e-ed69-4e0a-9b9c-074b74bc48f5"
		token := "eyJhbGciOiJSU0ExXzUiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0.KDvZMG7eRuIzgiFOYcKNZa3DY3YnnhnVlianZLy7Ae0XaCJCxzgNRSXfppZqGP-prW9FvuYx9Ja2l0jMV5834tUU8H7uLfZtCupNJ_0G1-TmQLDV1YtgYa_aHH2ZWXbDyNRhr-SFABgTL0niSm-f2mmiIH-gE0g6DrcPWFXnLsBrQytypozEDJf9VmfSHBraxi60hFmZ2p_ZUiyktvKQClx-3k6PvMmXSTXqVXtnTQ-8wFevv9Pd0Hi6YGET7ur5NjvlSM5p5_I016leN1btuURdBSauvfhcUE5CSMjLB7K-LNCAgvoaSwhHZLuwYTBWc4H0YJHZcEe4_TvTdrqsxw.JJdmhfX9ZapvH7otD67DMA.TGgLa9vEJqJelh6ofkOV_vFATd0fDR9pC_H-n0g-vFpfRMlzOEjSu_xVkOCdTeVBtFkrOHyUgvOLZ1CYRsnLQNzW50WcJtUUmonpho5cDZlEGHtPiTGE9yqlKsXucD8ph9nXJvikYdE86KVmLh4vjIbzUZe7GU0b9u9Cj3hdMLc2NkE85MymCQXoGKBHg1r18edSBmn1YNrTm2LWZIsa34cSqGPtZ68io7UWC6TRzwglrw-FRLWZJSacpED5uNxhf54RkxtjYZxNOK4perd3uoK3ClnjPybjzZFKLE7d6ZuQRX9Z83GT466ugbQLtwWJLZc9BoB1quxboalRL2nuQymurpKuPDITlsmpn9rFOllGpztOsx8IxxPVUa-7y-6TVw8Z2bVN65CCyl1JopBLCw.KcigqhYJ1nqfOmAiIRgFCA"
		device, err := ua.DeviceRepo.GetDeviceByTokenId(c, tokenId)

		devices, _ := ua.DeviceRepo.GetDevicesByTokenIds(c, []string{tokenId})
		fmt.Sprintf("%v", devices)

		ua.DeviceRepo.InsertDevice(c, Device{
			HouseholdRtvExtId: "dupa",
			RtvSerialNumber:   "",
			RtvTerminalModel:  "",
			ProfileId:         "",
			UserId:            "",
			AuthCreationDate:  time.Now(),
			AuthExpireDate:    time.Now().Add(time.Hour),
			TokenPrivKey:      nil,
			FcmToken:          "",
			AccountCession:    false,
			IsMainRtvMember:   false,
		})

		_, err = ua.Authenticate(c, "", token, device.TokenPrivKey)
		if err != nil {
			panic(err)
		}
	}
}

func (ua *DeviceAuthentication) Authenticate(ctx context.Context, tokenId string, token string, tokenPrvKeyBytes []byte) (ctxWithAuthentication context.Context, err error) {
	authBearer, err := ua.decodeToken(ctx, token, tokenPrvKeyBytes)
	if err != nil {
		//TODO
		return nil, err
	}

	ctxWithAuthentication = context.WithValue(ctx, "authBearer", authBearer)
	return ctxWithAuthentication, nil
}

func (ua *DeviceAuthentication) decodeToken(ctx context.Context, token string, tokenPrvKeyBytes []byte) (authBearer string, err error) {

	decrypted, err := jwe.Decrypt([]byte(token), jwe.WithKey(jwa.RSA1_5, "jakies_tam_dane"))
	if err != nil {
		//TODO
		return "", err
	}

	var payload bytes.Buffer
	_, err = payload.Write(decrypted)
	if err != nil {
		//TODO
		return "", err
	}

	decoder := gob.NewDecoder(&payload)
	var bearer string
	err = decoder.Decode(&bearer)
	if err != nil {
		//TODO
		return "", err
	}

	return bearer, nil
}
