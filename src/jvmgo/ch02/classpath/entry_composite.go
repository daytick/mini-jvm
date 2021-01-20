package classpath

import (
	"strings"
	"errors"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	
}

func (self CompositeEntry) readClass(className string) ([]byte, CompositeEntry, error) {}

func (self CompositeEntry) String() string {}
