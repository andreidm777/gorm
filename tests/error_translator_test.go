package tests_test

import (
	"errors"
	"testing"

	"github.com/andreidm777/gorm"
	"github.com/andreidm777/gorm/utils/tests"
)

func TestDialectorWithErrorTranslatorSupport(t *testing.T) {
	translatedErr := errors.New("translated error")
	db, _ := gorm.Open(tests.DummyDialector{TranslatedErr: translatedErr})

	err := db.AddError(errors.New("some random error"))
	if !errors.Is(err, translatedErr) {
		t.Fatalf("expected err: %v got err: %v", translatedErr, err)
	}
}
