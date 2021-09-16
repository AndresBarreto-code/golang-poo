package main

import (
	factory "CursoGolangPlatzi/ObjectOriented/Classes/factory"
	"fmt"
)

func main() {
	smsFactory, _ := factory.GetNotificationFactory("SMS")
	emailFactory, _ := factory.GetNotificationFactory("Email")
	fmt.Println(smsFactory)
	fmt.Println(emailFactory)
	factory.SendNotification(smsFactory)
	factory.SendNotification(emailFactory)
	factory.GetMethod(smsFactory)
	factory.GetMethod(emailFactory)

}
