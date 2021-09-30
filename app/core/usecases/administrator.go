package usecases

import (
	"CoffeLand/app/core/data"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AdministratorUsecases struct {
	Admin data.Administrator
	Core data.Core
}

func NewAdministratorUsecases(core data.Core, adminId string) (*AdministratorUsecases, error) {
	admin, err := core.AdminRepo.GetByID(adminId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &AdministratorUsecases{admin, core}, nil
}

func(a *AdministratorUsecases) PutAdministrator(admin data.Administrator) (string, error) {
	if err := admin.Validate(); err != nil {
		return "", errors.WithStack(err)
	}

	if err := a.Admin.HasAddAdministratorAccess(); err != nil {
		return "", err
	}

	if len(admin.ID) == 0 {
		admin.ID = uuid.New().String()
	}
	if err := a.Core.AdminRepo.Store(admin); err != nil {
		return admin.ID, errors.WithStack(err)
	}
	return admin.ID, nil
}

func(a *AdministratorUsecases) PutDiscount(discount data.Discount) (string, error) {
	if err := a.Admin.HasDiscountAccess(); err != nil {
		return "", errors.WithStack(err)
	}
	return NewDiscountUsecases(a.Core).putDiscount(discount)
}

func(a *AdministratorUsecases) PutProduct(product data.Product) (string, error) {
	if err := a.Admin.HasProductAccess(); err != nil {
		return "", errors.WithStack(err)
	}
	return NewProductUsecases(a.Core).putProduct(product)
}

func(a *AdministratorUsecases) PutBlogPost(post data.BlogPost) (string, error) {
	if err := a.Admin.HasAddBlogPostsAccess(); err != nil {
		return "", errors.WithStack(err)
	}
	return NewBlogPostUsecases(a.Core).put(post)
}
