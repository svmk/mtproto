package mtproto

import (
	"math/rand"
	"time"
)

// Session storage interface
type ISession interface {
	IsIPv6() bool
	// IsEncrypted returns true if AuthKey, ServerSalt and SessionID fields aren't empty
	IsEncrypted() bool

	GetAddress() string
	GetAuthKey() []byte
	GetAuthKeyHash() []byte
	GetServerSalt() []byte
	GetSessionID() int64

	SetAddress(string)
	SetAuthKey([]byte)
	SetAuthKeyHash([]byte)
	SetServerSalt([]byte)
	SetSessionID(int64)

	UseIPv6(bool)
	Encrypted(bool)
}

type Session struct {
	address     string
	authKey     []byte
	authKeyHash []byte
	serverSalt  []byte
	sessionId   int64
	useIPv6     bool
	encrypted   bool
}

func NewSession() ISession {
	session := &Session{}

	rand.Seed(time.Now().UnixNano())
	session.SetSessionID(rand.Int63())

	return session
}

func (s Session) IsIPv6() bool {
	return s.useIPv6
}

func (s Session) IsEncrypted() bool {
	return s.encrypted
}

func (s Session) GetAddress() string {
	return s.address
}

func (s Session) GetAuthKey() []byte {
	return s.authKey
}

func (s Session) GetAuthKeyHash() []byte {
	return s.authKeyHash
}

func (s Session) GetServerSalt() []byte {
	return s.serverSalt
}

func (s Session) GetSessionID() int64 {
	return s.sessionId
}

func (s *Session) SetAddress(address string) {
	s.address = address
}

func (s *Session) SetAuthKey(authKey []byte) {
	s.authKey = make([]byte, len(authKey))
	copy(s.authKey, authKey)
}

func (s *Session) SetAuthKeyHash(authKeyHash []byte) {
	s.authKeyHash = make([]byte, len(authKeyHash))
	copy(s.authKeyHash, authKeyHash)
}

func (s *Session) SetServerSalt(salt []byte) {
	s.serverSalt = make([]byte, len(salt))
	copy(s.serverSalt, salt)
}

func (s *Session) SetSessionID(ID int64) {
	s.sessionId = ID
}

func (s *Session) UseIPv6(useIPv6 bool) {
	s.useIPv6 = useIPv6
}

func (s *Session) Encrypted(encrypted bool) {
	s.encrypted = encrypted
}
