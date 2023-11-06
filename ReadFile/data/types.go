
package data

type distance float64 //miles
type distanceKm float64

//Method
func (miles distance) ToKm() distanceKm{
	return distanceKm(1.68934 * miles)
}

func (km distanceKm) ToMiles() distance {
	return distance(km/1.68934)
}

func test(){
	d := distance(4.5)
	km := d.ToKm()
	km.ToMiles()
	print(d)
}