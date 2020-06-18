package dumpctx

import (
	"fmt"
	"io"
)

type Ctx struct {
	indent string
	isList bool
}

func (ctx *Ctx) SetIsList()   { ctx.isList = true }
func (ctx *Ctx) UnsetIsList() { ctx.isList = false }

func (ctx *Ctx) Indent() string { return ctx.indent }
func (ctx *Ctx) Enter()         { ctx.indent += IndentToken }
func (ctx *Ctx) EnterList()     { ctx.Enter(); ctx.isList = true }

func (ctx *Ctx) Leave() {
	if len(ctx.indent) == 0 {
		return
	}

	ctx.indent = ctx.indent[:len(ctx.indent)-len(IndentToken)]
}
func (ctx *Ctx) LeaveList() { ctx.Leave(); ctx.isList = false }

func (ctx *Ctx) NoList(cb func()) {
	isList := ctx.isList

	// clear is list if ctx is set
	if isList {
		ctx.isList = false
	}

	cb()

	// restore
	if isList {
		ctx.isList = true
	}
}

func (ctx *Ctx) EmitPrefix(w io.Writer) {
	fmt.Fprint(w, ctx.indent)

	if ctx.isList {
		fmt.Fprintf(w, "%c ", ListSeparator)
	} else {
		fmt.Fprint(w, IndentToken)
	}
}

func (ctx *Ctx) WrapList(cb func()) {
	ctx.EnterList()
	cb()
	ctx.LeaveList()
}

func (ctx *Ctx) Wrap(cb func()) {
	ctx.Enter()
	cb()
	ctx.Leave()
}

func (ctx *Ctx) Copy(other *Ctx) {
	ctx.indent = other.indent
	ctx.isList = other.isList
}

func (ctx *Ctx) Init() {
	ctx.indent = ""
	ctx.isList = false
}
