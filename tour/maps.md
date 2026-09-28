Maps, keys to valeus
Just like any other programming language?
You cannot add keys to a nil map, wow.

var m map[string]string

m + make(map[string]string)

or

m := make(map[string]string)


type Vertex struct {
  Lat, Long float64
}

coords := make(map[string]Vertex)

Map literals, like struct literals
but keys are required.

var m = map[string]Vertex{
  "Cafe": Vertex{37.2323, -454.7865}
}

or 

var m = map[string]Vertex{
  "Cafe": {37.2323, -454.7865}
}


Mutating maps

insertion: m[key] = elem

get: elem = m[key]

del: delete(m, key)

checking for a key: elem, ok := m[key]
ok will be true if key exists