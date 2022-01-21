package mypackage

// CarPlublic struc con acceso publico
// struct con los atributos con:
// primera letra en Mayucula es publido
// primera letra en minuscula es privado
type CarPublic struct {
	Brand string
	Year  int
}

// carPrivate struct ~ clase privada
type carPrivate struct {
	brand string
	year  int
}
