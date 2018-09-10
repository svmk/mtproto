package mtproto

import (
	"errors"
	"os"
)

type SessionStorage interface {
	IsExists(session ISession) bool
	Save(session ISession) error
	Load(session ISession) error
}

type FileSessionStorage struct {
	authFile string
}

func NewFileSessionStorage(authFile string) *FileSessionStorage {
	return &FileSessionStorage{
		authFile: authFile,
	}
}

func (s *FileSessionStorage) IsExists(session ISession) bool {
	stat, err := os.Stat(s.authFile)
	if os.IsNotExist(err) {
		return false
	}
	if stat.Size() == 0 {
		return false
	}
	return true
}

func (s *FileSessionStorage) Load(session ISession) error {
	file, err := os.OpenFile(s.authFile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	// TODO: Magic number
	buffer := make([]byte, 1024*4)
	n, _ := file.ReadAt(buffer, 0)
	if n <= 0 {
		return errors.New("Need create new session ")
	}

	decoder := NewDecodeBuf(buffer)
	session.SetAuthKey(decoder.StringBytes())
	session.SetAuthKeyHash(decoder.StringBytes())
	session.SetServerSalt(decoder.StringBytes())
	session.SetAddress(decoder.String())
	session.UseIPv6(false)
	if decoder.UInt() == 1 {
		session.UseIPv6(true)
	}

	if decoder.err != nil {
		return decoder.err
	}

	return nil
}

func (s FileSessionStorage) Save(session ISession) error {
	file, err := os.OpenFile(s.authFile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	// TODO: Magic number
	buffer := NewEncodeBuf(1024)
	buffer.StringBytes(session.GetAuthKey())
	buffer.StringBytes(session.GetAuthKeyHash())
	buffer.StringBytes(session.GetServerSalt())
	buffer.String(session.GetAddress())

	var useIPv6UInt uint32
	if session.IsIPv6() {
		useIPv6UInt = 1
	}
	buffer.UInt(useIPv6UInt)

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.WriteAt(buffer.buf, 0)
	if err != nil {
		return err
	}

	return nil
}

type dummySessionStorage struct {
}

func NewDummySessionStorage() SessionStorage {
	return &dummySessionStorage{}
}

func (s *dummySessionStorage) IsExists(session ISession) bool {
	return false
}

func (s *dummySessionStorage) Load(session ISession) error {
	return nil
}

func (s *dummySessionStorage) Save(session ISession) error {
	return nil
}
