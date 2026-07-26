type ParkingSystem struct {
    Count [4]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    count := [4]int{0, big, medium, small}
    return ParkingSystem{Count: count}
}


func (this *ParkingSystem) AddCar(carType int) bool {
    this.Count[carType]--
    return this.Count[carType] >= 0 
}
