package interfaces



type MemberRepo interface {
	SaveMember() (bool, error)
	
}
