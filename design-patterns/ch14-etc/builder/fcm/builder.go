package fcm

type IBuilder interface {
	SetPlatform(platform string) IBuilder
	SetToken(token string) IBuilder
	SetAOS(aos AOSData) IBuilder
	SetIOS(ios IOSData) IBuilder
	Build() Message
}

type Builder struct {
	platform string
	token    string
	aos      AOSData
	ios      IOSData
}

func (b *Builder) SetPlatform(platform string) IBuilder {
	b.platform = platform
	return b
}

func (b *Builder) SetToken(token string) IBuilder {
	b.token = token
	return b
}

func (b *Builder) SetAOS(aos AOSData) IBuilder {
	b.aos = aos
	return b
}

func (b *Builder) SetIOS(ios IOSData) IBuilder {
	b.ios = ios
	return b
}

func (b *Builder) Build() Message {
	return Message{
		Platform: b.platform,
		Token:    b.token,
		AOS:      b.aos,
		IOS:      b.ios,
	}
}

func NewBuilder() IBuilder {
	return &Builder{}
}

func NewAOSBuilder() IBuilder {
	return &Builder{
		platform: "aos",
	}
}

func NewIOSBuilder() IBuilder {
	return &Builder{
		platform: "ios",
	}
}
