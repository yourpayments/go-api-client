# «Твои платежи»: Интеграция на Go
Готовая библиотека + подробные примеры. 
![](https://repository-images.githubusercontent.com/816212690/e94dc73f-3bcb-4737-9336-793940827cd3)


## Установка за 1 минуту
Убедитесь, что ваш проект использует модули Go (в его корне будет файл go.mod, если он уже есть):
```shell
go mod init
```
Затем импортируйте в программу модуль ypmn:
```go
import (
    "github.com/yourpayments/go-api-client"
)
```
Также, вы можете явно добавить пакет в проект:
```shell
go get -u github.com/yourpayments/go-api-client
```

## Примеры
1. [Cамый простой платёж](examples/simple_get_payment_link.go)
2. [Списание средств](examples/capture.go)
3. [Возврат средств](examples/refund.go)
4. [Запрос отчёта для сверки](examples/general_report.go)
5. [Получение вебхуков](examples/webhook.go)
6. [Запрос статуса транзакции](examples/status.go)

## Ссылки
- [НКО «Твои платежи»](https://YPMN.ru/)
- [Докуметация API](https://ypmn.ru/ru/documentation/)
- [Тестовые банковские карты](https://ypmn.ru/ru/documentation/#tag/testing)
- [Задать вопрос или сообщить о проблеме](https://github.com/yourpayments/go-api-client/issues/new)

-------------
🟢 [«Твои платежи»](https://YPMN.ru/ "Платёжная система для сайтов, платформ и приложений") -- финтех-составляющая для сайтов, платформ и приложений
