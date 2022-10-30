package sauce

type PlumTomatoSauce struct {
	name string
}

func CreatePlumTomatoSauce() ISauce {
	return &PlumTomatoSauce{name: "플럼 토마토 소스"}
}

func (s PlumTomatoSauce) ToString() string {
	return s.name
}
