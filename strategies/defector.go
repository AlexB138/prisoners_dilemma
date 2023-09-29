package strategies

type Defector struct {
	name string
}

func NewDefector() Defector {
	return Defector{name: "Defector"}
}

func (d *Defector) MakeChoice(round int) Action {

}

func (d *Defector) GetName() string {
	return d.name
}
