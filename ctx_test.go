package dumpctx

import (
	"bytes"
	"testing"
)

func TestCtx(t *testing.T) {
	ctx := &Ctx{}
	ctx.Init()

	ctx.SetIsList()
	if !ctx.isList {
		t.Errorf("list attr is not set")
	}

	ctx.UnsetIsList()
	if ctx.isList {
		t.Errorf("list attr was not unset")
	}

	ctx.NoList(func() {})

	b := bytes.Buffer{}
	ctx.EmitPrefix(&b)
	if !ctx.isList && b.String() != IndentToken {
		t.Errorf("prefix wasn't emited")
	}
}
