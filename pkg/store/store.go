package store

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var Data Store
var kkDatabase string

const (
	storeDir  = ".keykube"
	storeFile = "database"
)

type Store []Entity

type Entity struct {
	Name      string `json:"name"`
	Vault     Vault  `json:"vault"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	ViewedAt  int64  `json:"viewedAt"`
	ExpireAt  int64  `json:"expireAt"`
}

type Vault struct {
	Password    string `json:"password"`
	OldPassword string `json:"oldPassword"`
}

func init() {
	homeDir, _ := os.UserHomeDir()
	kkRoot := homeDir + "/" + storeDir
	kkDatabase = kkRoot + "/" + storeFile

	os.Mkdir(kkRoot, 0700)

	file, _ := os.ReadFile(kkDatabase)
	json.Unmarshal(file, &Data)
}

func (s *Store) json() []byte {
	bytes, err := json.Marshal(s)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return bytes
}

func (s *Store) save() {
	file, err := os.OpenFile(kkDatabase, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	file.Truncate(0)
	file.Seek(0, 0)
	file.Write(Data.json())
}

func (e *Entity) Password() string {
	return e.Vault.Password
}

func (e *Entity) OldPassword() string {
	return e.Vault.OldPassword
}

func (e *Entity) Print() {
	fmt.Println()
	fmt.Println("Login:", e.Name)
	fmt.Println("Password:", e.Password())

	if e.OldPassword() != "" {
		fmt.Println("Old password:", e.OldPassword())
	}

	expireAt := "never"

	if e.ExpireAt > 0 {
		expireAt = time.Unix(e.ExpireAt, 0).Format(time.DateOnly)
	}

	fmt.Printf("Expire at: %v\n\n", expireAt)
}

func (e *Entity) Update(p Entity) *Entity {
	if p.Name != "" {
		e.Name = p.Name
	}
	if p.Vault != (Vault{}) {
		e.Vault = p.Vault
	}

	now := time.Now().Unix()

	e.UpdatedAt = now
	e.ViewedAt = now
	e.ExpireAt = p.ExpireAt

	Data.save()

	return e
}

func Add(name string, password string, ExpireAt int64) *Entity {
	now := time.Now().Unix()

	item := Entity{
		Name: name,
		Vault: Vault{
			Password: password,
		},
		CreatedAt: now,
		UpdatedAt: now,
		ExpireAt:  ExpireAt,
	}

	Data = append(Data, item)
	Data.save()

	return &item
}

func Get(name string) (*Entity, bool) {
	for index, entity := range Data {
		if entity.Name == name {
			return &Data[index], true
		}
	}

	return &Entity{}, false
}

func Delete(name string) {
	for index, entity := range Data {
		if entity.Name == name {
			Data = append(Data[:index], Data[index+1:]...)
			Data.save()
			break
		}
	}
}

func Exists(name string) bool {
	_, ok := Get(name)
	return ok
}
