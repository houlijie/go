package model

type student struct {
	Name  string
	score float64
}

// student首字母小写， 只能在包内使用， 通过工厂模式解决
func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s, // 引用的时候会报错， 私有方法不能被使用
	}
}

func (s *student) GetScore() float64 {
	return s.score // 类似php的set/get
}
