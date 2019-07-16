package requestParser

import (
	"github.com/stretchr/testify/mock"
	db "github.com/xmidt-org/codex-db"
	"github.com/xmidt-org/voynicrypto"
)

type mockEncrypter struct {
	mock.Mock
}

func (md *mockEncrypter) EncryptMessage(message []byte) ([]byte, []byte, error) {
	args := md.Called(message)
	return message, []byte{}, args.Error(0)
}

func (*mockEncrypter) GetAlgorithm() voynicrypto.AlgorithmType {
	return voynicrypto.None
}

func (*mockEncrypter) GetKID() string {
	return "none"
}

type mockBlacklist struct {
	mock.Mock
}

func (mb *mockBlacklist) InList(ID string) (reason string, ok bool) {
	args := mb.Called(ID)
	reason = args.String(0)
	ok = args.Bool(1)
	return
}

type mockInserter struct {
	mock.Mock
}

func (i *mockInserter) Insert(record db.Record) {
	i.Called(record)
	return
}
