package mode

type Mode string

const (
	CBC Mode = "CBC"
	CFB Mode = "CFB"
	CTR Mode = "CTR"
	OFB Mode = "OFB"
	GCM Mode = "GCM"
	ECB Mode = "ECB"
)

func (m Mode) String() string {
	return string(m)
}

func (m Mode) Is(ms Mode) bool {
	return m == ms
}

func (m Mode) Not(ms ...Mode) bool {
	for _, m1 := range ms {
		if m == m1 {
			return false
		}
	}
	return true
}

func IsNotSupported(m Mode) bool {
	return m != CBC && m != CFB && m != CTR && m != OFB && m != GCM && m != ECB
}
