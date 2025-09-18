package interfaces

import "github.com/eswaribala/claimapp/healthinsurance/store"

type MemberRepo interface {
	SaveMember() (bool, error)
	GetAllMembers() ([]*store.Member, error)
	GetMemberByID(id uint) (*store.Member, error)
	GetMemberByEmail(email string) (*store.Member, error)
	GetMemberByPhone(phone string) (*[]store.Member, error)
	UpdateMember(email string, contactNo string) (bool, error)
	DeleteMember(id uint) (bool, error)
}
