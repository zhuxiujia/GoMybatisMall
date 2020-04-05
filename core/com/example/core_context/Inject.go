package core_context

import (
	"github.com/facebookgo/inject"
)

// Typically an application will have exactly one object graph, and
// you will create it and use it within a main function:
var Context = inject.Graph{}
