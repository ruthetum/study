package sauce

type MarinaraSauce struct {
	name string
}

func CreateMarinaraSauce() ISauce {
	return &MarinaraSauce{name: "마리나라 소스"}
}

func (s MarinaraSauce) ToString() string {
	return s.name
}
