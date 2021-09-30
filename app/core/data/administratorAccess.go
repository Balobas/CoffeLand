package data

type AdministratorAccess string

const (
	AdministratorAccessLow = AdministratorAccess("loh") // Начальный уровень доступа для ознакомления с системой
	AdministratorAccessMiddle = AdministratorAccess("middle") // Уровень для работы с продуктами и скидками
	AdministratorAccessHigh = AdministratorAccess("high") // Уровень управления администраторами
)
