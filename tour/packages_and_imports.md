## A go program is made of packages.

packages can be imported like below, multiple imports or factored imports

import "fmt"
import "math/rand"

or

import (
  "fmt"
  "math/rand"
)


## Convention:

"math/rand" - import path

- last element ("rand") is package name
- "math" is directory or package group, math itself is also a package