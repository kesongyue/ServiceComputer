package entity

type Meeting struct {
	Sponsor      string
	Participator []string
	Start, End   Date
	Title        string
}

func GetSponsor(m Meeting) string {
	return m.Sponsor
}
func GetParticipator(m Meeting) []string {
	return m.Participator
}
func GetStart(m Meeting) Date {
	return m.Start
}
func GetEnd(m Meeting) Date {
	return m.End
}
func GetTitle(m Meeting) string {
	return m.Title
}
func (meet Meeting) IsParticipator(username string) bool {
	for _, partic := range meet.Participator {
		if partic == username {
			return true
		}
	}
	return false
}

func (meet *Meeting) DeleteParticipator(name string) {
	var remainPartic []string
	for _, partic := range meet.Participator {
		if partic != name {
			remainPartic = append(remainPartic, partic)
		}
	}
	meet.Participator = remainPartic
}
