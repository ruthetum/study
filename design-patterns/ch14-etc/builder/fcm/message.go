package fcm

type Message struct {
	Platform string
	Token    string
	AOS      AOSData
	IOS      IOSData
}

type AOSData struct {
	Data string
	Path string
}

type IOSData struct {
	APN      string
	Content  string
	Redirect string
}
