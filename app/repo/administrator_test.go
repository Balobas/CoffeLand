package repo

import (
	"CoffeLand/app/core/data"
	"CoffeLand/app/interfaces/database"
	"testing"
)

func getAdministratorModel(t *testing.T) data.AdministratorRepository {
	db, err := database.MySQLDB()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return NewAdministratorRepo(db)
}

func TestAdministratorModel_Store(t *testing.T) {
	if err := getAdministratorModel(t).Store(data.Administrator{
		ID:         "aaa",
		FirstName:  "Виктор",
		LastName:   "Викториев",
		Patronymic: "Викторович",
	}); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestAdministratorModel_GetByID(t *testing.T) {
	if admin, err := getAdministratorModel(t).GetByID("aaa"); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log(admin)
	}

	if _, err := getAdministratorModel(t).GetByID("dsafsdlkad"); err == nil {
		t.Errorf("Expected no records found error")
		t.FailNow()
	}
}

func TestAdministratorModel_GetByFIO(t *testing.T) {
	adminModel := getAdministratorModel(t)

	firstName := "Виктор"
	lastName := "Викториев"
	patronymic := "Викторович"

	if admins, err := adminModel.GetByFIO(firstName, lastName, patronymic); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log(admins)
	}

	if admins, err := adminModel.GetByFIO(firstName, "", ""); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log(admins)
	}

	if admins, err := adminModel.GetByFIO(firstName, "", patronymic); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log(admins)
	}

	if admins, err := adminModel.GetByFIO("", lastName, patronymic); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log(admins)
	}

	if admins, err := adminModel.GetByFIO(firstName, "", patronymic); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log(admins)
	}

	if admins, err := adminModel.GetByFIO("", "", ""); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		t.Log(admins)
	}

	if _, err := adminModel.GetByFIO("Димидрол", lastName, patronymic); err == nil {
		t.Errorf("expected no records found error")
		t.FailNow()
	}
}
