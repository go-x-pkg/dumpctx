package dumpctx

import (
	"fmt"
	"io"
)

type Ctx struct {
	indent string
	isList bool
}

func (it *Ctx) SetIsList()   { it.isList = true }
func (it *Ctx) UnsetIsList() { it.isList = false }

func (it *Ctx) Indent() string { return it.indent }
func (it *Ctx) Enter()         { it.indent += IndentToken }
func (it *Ctx) EnterList()     { it.Enter(); it.isList = true }

func (it *Ctx) Leave() {
	if len(it.indent) == 0 {
		return
	}

	it.indent = it.indent[:len(it.indent)-len(IndentToken)]
}
func (it *Ctx) LeaveList() { it.Leave(); it.isList = false }

func (it *Ctx) NoList(cb func()) {
	isList := it.isList

	// clear is list if it is set
	if isList {
		it.isList = false
	}

	cb()

	// restore
	if isList {
		it.isList = true
	}
}

func (it *Ctx) EmitPrefix(w io.Writer) {
	fmt.Fprint(w, it.indent)

	if it.isList {
		fmt.Fprintf(w, "%c ", ListSeparator)
	} else {
		fmt.Fprintf(w, "%s", IndentToken)
	}
}

func (it *Ctx) WrapList(cb func()) {
	it.EnterList()
	cb()
	it.LeaveList()
}

func (it *Ctx) Wrap(cb func()) {
	it.Enter()
	cb()
	it.Leave()
}

func (it *Ctx) Copy(other *Ctx) {
	it.indent = other.indent
	it.isList = other.isList
}

func (it *Ctx) Init() {
	it.indent = ""
	it.isList = false
}
