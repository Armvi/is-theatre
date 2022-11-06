package entity

type Performance struct {
	Id                  int64                `json:"id,omitempty"`
	PerformanceName     string               `json:"performance_name,omitempty"`
	Composition         *Composition         `json:"composition,omitempty"`
	PerformanceDirector *PerformanceDirector `json:"performance_director,omitempty"`
	Description         string               `json:"description,omitempty"`
	ActorsRoles         []ActorsRole         `json:"actors_roles"`
}
