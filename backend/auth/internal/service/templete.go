package service

import (
	"fmt"
	"os"
)

func GetTemplate(token string) string {
	host := os.Getenv("SENDER_HOST")
	return fmt.Sprintf(`
    <!DOCTYPE html>
    <html lang="ru">
    <head><meta charset="UTF-8"><title>Подтверждение регистрации</title></head>
    <body style="font-family: Arial, sans-serif; background-color: #f4f4f4; padding: 20px;">
    <div style="background-color: #ffffff; max-width: 600px; margin: 0 auto; border-radius: 8px; padding: 20px;">
        <h2 style="color: #333;">Добро пожаловать в React-Go Shop!</h2>
        <p style="color: #555;">Благодарим за регистрацию. Чтобы завершить процесс, подтвердите адрес электронной почты.</p>
        <p>
            <a href="http://%s/api/verify?token=%s" style="display: inline-block; background-color: #1cc29f; color: #fff; padding: 12px 20px; text-decoration: none; border-radius: 6px;">
                Подтвердить email
            </a>
        </p>
        <p style="color: #888; font-size: 14px;">
            Если вы не регистрировались — просто проигнорируйте это письмо.
        </p>
    </div>
    </body>
    </html>
`, host, token)

}
