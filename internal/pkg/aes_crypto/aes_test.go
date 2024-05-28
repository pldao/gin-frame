package aes_crypto

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAesEncryptCFB(t *testing.T) {
	realData := "username=super_admin&password=123456"
	aecData := "2a4cc91f9032c4de77584dd8c74997b22192491ab0374116b9f0b8370a3af3ed74b9fbb1319af212dbc61bf8ae60918925209270"
	aecDataRes, err := EnPwdCode(realData)
	if err != nil {
		t.Error(err)
	}
	data, err := DePwdCode(aecData)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, data, realData)
	aecDataRes2, err := DePwdCode(aecDataRes)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, aecDataRes2, realData)
}
