package geo

// Point ...
type Point struct {
	lat float64
	lon float64
}

// Bound ...
type Bound struct {
	near Point
	far  Point
}

func (b *Bound) init(lat0 float64, long0 float64, lat1 float64, long1 float64) {
	b.near = Point{lat0, long0}
	b.far = Point{lat1, long1}
}

func (b *Bound) center() Point {
	return Point{(b.near.lat + b.far.lat) / 2, (b.near.lon + b.far.lon) / 2}
}

func (b *Bound) northWest() Bound {
	c := b.center()
	return Bound{near: Point{c.lat, b.near.lon}, far: Point{b.far.lat, c.lon}}
}

func (b *Bound) northEast() Bound {
	c := b.center()
	return Bound{near: Point{c.lat, c.lon}, far: Point{b.far.lat, b.far.lon}}
}

func (b *Bound) southWest() Bound {
	c := b.center()
	return Bound{near: Point{b.near.lat, b.near.lon}, far: Point{c.lat, c.lon}}
}

func (b *Bound) southEast() Bound {
	c := b.center()
	return Bound{near: Point{b.near.lat, c.lon}, far: Point{c.lat, b.far.lon}}
}

func (b *Bound) contains(p Point) bool {
	return p.lat >= b.near.lat && p.lat <= b.far.lat && p.lon >= b.near.lon && p.lon <= b.far.lon
}

// QuadTree ... Devide space into quadtree
type QuadTree struct {
	id    int64
	level int
	bound Bound
	grid0 *QuadTree
	grid1 *QuadTree
	grid2 *QuadTree
	grid3 *QuadTree
}

func (q *QuadTree) init(id int64, level int, bound Bound) {
	q.id = id
	q.level = level
	q.bound = bound
}

func (q *QuadTree) subID(pid int64, sub int) int64 {
	return pid*4 + int64(sub)
}

func (q *QuadTree) divide() {
	if q.level == 20 {
		return
	}
	// if aleady divided , return
	if q.grid0 != nil {
		return
	}

	q.grid0 = new(QuadTree)
	q.grid0.init(q.subID(q.id, 0), q.level+1, q.bound.northWest())

	q.grid1 = new(QuadTree)
	q.grid1.init(q.subID(q.id, 1), q.level+1, q.bound.northEast())

	q.grid2 = new(QuadTree)
	q.grid2.init(q.subID(q.id, 2), q.level+1, q.bound.southWest())

	q.grid3 = new(QuadTree)
	q.grid3.init(q.subID(q.id, 3), q.level+1, q.bound.southEast())
}

//merge ... reverse of divide
func (q *QuadTree) merge() {
	q.grid0 = nil
	q.grid1 = nil
	q.grid2 = nil
	q.grid3 = nil
}

// EARTH ...
var EARTH = QuadTree{
	id:    0,
	level: 0,
	bound: Bound{
		near: Point{lat: -90, lon: -180},
		far:  Point{lat: 90, lon: 180},
	},
}

// GenerateHashID generate hash id of location
func GenerateHashID(lat float64, lon float64, level int) int64 {
	p := Point{lat, lon}
	return EARTH.HashID(p, level)
}

// HashID ...
func (q *QuadTree) HashID(p Point, level int) int64 {

	if !q.bound.contains(p) {
		return -1
	}

	if q.level == level {
		return q.id
	}

	q.divide()
	if q.grid0 == nil {
		return -1
	}

	id := q.grid0.HashID(p, level)
	if id != -1 {
		return id
	}

	id = q.grid1.HashID(p, level)
	if id != -1 {
		return id
	}

	id = q.grid2.HashID(p, level)
	if id != -1 {
		return id
	}

	return q.grid3.HashID(p, level)

}
