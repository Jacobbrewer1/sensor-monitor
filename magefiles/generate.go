//go:build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func GenerateAll() {
	mg.Deps(Init)
	mg.Deps(mg.F(Generate.mock, false))
	//mg.Deps(mg.F(Generate.db, false))
	//mg.Deps(mg.F(Generate.openAPI, false))
}

type Generate mg.Namespace

func (Generate) Mock() error {
	mg.Deps(mg.F(Generate.mock, false))

	// Cover any new dependencies that may have been added by the code generation.
	return VendorDeps()
}

func (Generate) mock(shouldVendor bool) error {
	mg.Deps(Init)

	args := []string{
		"run",
		"//:gen_mock",
	}

	if err := sh.Run("bazel", args...); err != nil {
		return fmt.Errorf("error writing generated mock files: %w", err)
	}

	if shouldVendor {
		if err := VendorDeps(); err != nil {
			return err
		}
	}

	return nil
}

func (Generate) DB() error {
	mg.Deps(mg.F(Generate.db, true))

	// Cover any new dependencies that may have been added by the code generation.
	return VendorDeps()
}

func (Generate) db(shouldVendor bool) error {
	if err := sh.Run("sqlc", "generate"); err != nil {
		return fmt.Errorf("error generating sqlc files: %w", err)
	}

	if err := sh.Run("sqlc", "vet"); err != nil {
		return fmt.Errorf("error vetting sqlc files: %w", err)
	}

	if err := sh.Run("atlas", "migrate", "diff", "retailers",
		"--dir", "file://database/migrations",
		"--to", "file://database/schema.sql",
		"--dev-url", "docker://mariadb/latest/storepro",
	); err != nil {
		return fmt.Errorf("error generating atlas migration: %w", err)
	}

	if shouldVendor {
		if err := VendorDeps(); err != nil {
			return fmt.Errorf("error vendoring dependencies: %w", err)
		}
	}

	return nil
}

func (Generate) OpenAPI() error {
	mg.Deps(mg.F(Generate.openAPI, false))

	// Cover any new dependencies that may have been added by the code generation.
	return VendorDeps()
}

func (Generate) openAPI(shouldVendor bool) error {
	args := []string{
		"run",
		"//:gen_openapi",
	}

	if err := sh.Run("bazel", args...); err != nil {
		return fmt.Errorf("error writing generated open-api files: %w", err)
	}

	if shouldVendor {
		return VendorDeps()
	}

	return nil
}
