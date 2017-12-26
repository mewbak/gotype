package pmmpost

/*
BSD License
Copyright (c) 2017, Norbert Pillmayer <norbert@pillmayer.com>

All rights reserved.
Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:
1. Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.
3. Neither the name of Norbert Pillmayer nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

---------------------------------------------------------------------------

 * Internal commands for the Poor Man's MetaPost graphics language.

*/

import (
	"bytes"
	"fmt"

	"github.com/npillmayer/gotype/gtcore/arithmetic"
	"github.com/npillmayer/gotype/syntax"
	pmmp "github.com/npillmayer/gotype/syntax/pmmpost/statements"
	dec "github.com/shopspring/decimal"
)

// === Handling Variables and Constants ======================================

/* Allocate a variable in a memory frame. Existing variable references in
 * this memory frame will be overwritten !
 * Clients should probably first call FindVariableReferenceInMemory(vref).
 */
func (intp *PMMPostInterpreter) AllocateVariableInMemory(vref *PMMPVarRef,
	mf *syntax.DynamicMemoryFrame) *PMMPVarRef {
	//
	mf.Symbols().InsertSymbol(vref)
	T.P("var", vref.GetFullName()).Debugf("allocating variable in %s", mf.GetName())
	return vref
}

/* Given a variable reference, locate an incarnation in a memory frame. The
 * frame is determined by the variable's declaring scope: search for the top
 * frame linked to the scope.
 *
 * Variable references live in memory frames. Memory frames correspond to
 * scopes. To find a variable reference -- i.e. a living variable with a possible
 * value -- we have to proceed as follows:
 * (1) find the variable declaration in a scope, beginning at the top
 * (2) find the most recent memory frame pointing to this scope
 * (3) find a variable reference with the correct name in the memory frame
 * (4) if no reference/incarnation exists, create one
 *
 * Parameter doAlloc: should step (4) be performed ?
 */
func (intp *PMMPostInterpreter) FindVariableReferenceInMemory(vref *PMMPVarRef, doAlloc bool) (
	*PMMPVarRef, *syntax.DynamicMemoryFrame) {
	//
	if vref.decl == nil {
		T.P("var", vref.GetFullName()).Error("attempt to store variable without decl. in memory")
		return vref, nil
	}
	var sym *PMMPVarRef
	var memframe *syntax.DynamicMemoryFrame
	tagname := vref.decl.BaseTag.GetName()
	tag, scope := intp.scopeTree.Current().ResolveSymbol(tagname)
	if tag != nil { // found tag declaration in scope
		memframe = intp.memFrameStack.FindMemoryFrameWithScope(scope)
		varname := vref.GetName()
		T.P("var", varname).Debugf("var in ? %s", memframe)
		s := memframe.Symbols().ResolveSymbol(varname)
		if s == nil { // no variable ref incarnation => create one
			T.P("var", varname).Debug("not found in memory")
			if doAlloc {
				sym = intp.AllocateVariableInMemory(vref, memframe)
			}
		} else { // already present, return this on
			T.P("var", varname).Debug("variable already present in memory")
			sym = s.(*PMMPVarRef)
		}
	} else {
		// this should never happen: we could neither find nor construct a var decl
		panic(fmt.Sprintf("declaration for %s mysteriously vanished...", tagname))
	}
	return sym, memframe
}

/* The expression stack knows nothing about the interpreter's symbols, except
 * the few properties of interface Symbol. The expression stack deals with
 * polynomials and serial IDs of variables.
 *
 * Push a variable (numeric or pair type) onto the expression stack.
 */
func (intp *PMMPostInterpreter) PushVariable(vref *PMMPVarRef, asLValue bool) {
	if vref.IsPair() {
		if vref.IsKnown() && !asLValue {
			intp.PushConstant(vref) // put constant on expression stack
		} else {
			xpart, ypart := vref.XPart(), vref.YPart()
			intp.exprStack.PushVariable(xpart, ypart)
		}
	} else {
		if vref.IsKnown() && !asLValue {
			intp.PushConstant(vref) // put constant on expression stack
		} else {
			intp.exprStack.PushVariable(vref, nil)
		}
	}
}

/* Push a constant (numeric or pair type) onto the expression stack.
 */
func (intp *PMMPostInterpreter) PushConstant(vref *PMMPVarRef) {
	if vref.IsPair() {
		x := vref.XPart().GetValue()
		y := vref.YPart().GetValue()
		pair := arithmetic.MakePair(x.(dec.Decimal), y.(dec.Decimal))
		intp.exprStack.PushPairConstant(pair)
	} else {
		intp.exprStack.PushConstant(vref.value.(dec.Decimal))
	}
}

/* The expression stack knows nothing about the interpreter's symbols, except
 * the few properties of interface Symbol. The expression stack deals with
 * polynomials and serial IDs of variables. To get back from IDs to
 * variable references, we ask the expression stack for a Symbol (from an
 * ID). If the variable is of type pair, the Symbol may be a pair part (x-part
 * or y-part). Parts point to their parent symbol, thus giving us the
 * variable reference.
 */
func (intp *PMMPostInterpreter) getVariableFromExpression(e syntax.Expression) *PMMPVarRef {
	var v *PMMPVarRef
	if sym := intp.exprStack.GetVariable(e); sym != nil {
		var part *PairPartRef
		var ok bool
		if part, ok = sym.(*PairPartRef); ok {
			sym = part.pairvar
		}
		v = sym.(*PMMPVarRef)
		T.P("var", v.GetName()).Debugf("variable of type %s", typeString(v.GetType()))
	}
	return v
}

/* A variable which goes out of scope becomes a "capsule". We send a message
 * to the expression stack to forget the Symbol(s) for the ID(s) of a
 * variable. Variables are of type numeric or pair.
 */
func (intp *PMMPostInterpreter) encapsulateVariable(v *PMMPVarRef) {
	intp.exprStack.EncapsuleVariable(v.GetID())
	if v.IsPair() {
		var ypart *PairPartRef = v.GetFirstChild().(*PairPartRef)
		intp.exprStack.EncapsuleVariable(ypart.GetID())
	}
}

/* Make all variables in a memory frame "capsules".
 * When a memory frame is popped from the stack, the local variables living
 * in the frame have to be made "capsules". This is necessary, because they
 * may still be relevant to the LEQ-solver. The LEQ will finally decide
 * when to abondon the "zombie" variable.
 */
func (intp *PMMPostInterpreter) encapsulateVarsInMemory(mf *syntax.DynamicMemoryFrame) {
	mf.Symbols().Each(func(name string, sym syntax.Symbol) {
		vref := sym.(*PMMPVarRef)
		T.P("var", vref.GetFullName()).Debug("encapsule")
		intp.exprStack.EncapsuleVariable(vref.GetID()) // vref is now capsule
	})
}

// === Commands ==============================================================

/* Variable assignment.
 *
 * assignment : lvalue ASSIGN numtertiary
 *
 * (1) Retract lvalue from the resolver's table (make a capsule)
 * (3) Unset the value of lvalue
 * (3) Re-incarnate lvalue (get a new ID for it)
 * (4) Create equation on expression stack
 */
func (intp *PMMPostInterpreter) assign(lvalue *PMMPVarRef, e syntax.Expression) {
	varname := lvalue.GetName()
	oldserial := lvalue.GetID()
	T.P("var", varname).Debugf("assignment of lvalue #%d", oldserial)
	intp.encapsulateVariable(lvalue)
	vref, mf := intp.FindVariableReferenceInMemory(lvalue, false)
	vref.SetValue(nil) // now lvalue is unset / unsolved
	T.P("var", varname).Debugf("unset in %v", mf)
	vref.reincarnate()
	T.P("var", vref.GetName()).Debugf("new lvalue incarnation #%d", vref.GetID())
	intp.PushVariable(vref, false) // push LHS on stack
	intp.exprStack.Push(e)         // push RHS on stack
	intp.exprStack.EquateTOS2OS()  // construct equation
}

/* Save a tag within a group. The tag will be restored at the end of the
 * group. Save-commands within global scope will be ignored.
 * This method simply creates a var decl for the tag in the current scope.
 */
func (intp *PMMPostInterpreter) save(tag string) {
	sym, scope := intp.scopeTree.Current().ResolveSymbol(tag)
	if sym != nil { // already found in scope stack
		T.P("tag", tag).Debugf("save: found tag in scope %s",
			scope.GetName())
	}
	T.Debugf("declaring %s in current scope", tag)
	sym, _ = intp.scopeTree.Current().DefineSymbol(tag)
}

/* Declare a tag to be of type tp.
 *
 * If var decl is new, insert a new symbol in global scope. If var decl
 * already exists, erase all variables and re-enter var decl (MetaFont
 * semantics). If var decl has been "saved" in the current or in an outer scope,
 * make this tag a new undefined symbol.
 */
func (intp *PMMPostInterpreter) declare(tag string, tp int) *PMMPVarDecl {
	sym, scope := intp.scopeTree.Current().ResolveSymbol(tag)
	if sym != nil { // already found in scope stack
		T.P("tag", tag).Debug("declare: found tag in scope %s", scope.GetName())
		T.P("decl", tag).Debug("variable already declared - re-declaring")
		// Erase all existing variables and re-define symbol
		sym, _ = scope.DefineSymbol(tag)
		sym.(*PMMPVarDecl).SetType(tp)
		T.P("decl", tag).Errorf("TODO: retract variables (LEQ)")
	} else { // enter new symbol in global scope
		scope = intp.scopeTree.Globals()
		sym, _ = scope.DefineSymbol(tag)
		sym.(*PMMPVarDecl).SetType(tp)
	}
	T.P("decl", sym.GetName()).Debugf("declared symbol in %s", scope.GetName())
	return sym.(*PMMPVarDecl)
}

/* Create a variable reference. Parameters are the declaration for the variable,
 * a value and a flag, indicating if this variable should go to global memory.
 * The subscripts parameter is a slice of array-subscripts, if the variable
 * declaration is of array (complex) type.
 */
func (intp *PMMPostInterpreter) variable(decl *PMMPVarDecl, value interface{},
	subscripts []dec.Decimal, global bool) *PMMPVarRef {
	//
	var v *PMMPVarRef
	if decl.GetType() == pmmp.NumericType {
		v = CreatePMMPVarRef(decl, value, subscripts)
	} else {
		v = CreatePMMPPairTypeVarRef(decl, value, subscripts)
	}
	if global {
		intp.memFrameStack.Globals().Symbols().InsertSymbol(v)
	} else {
		intp.memFrameStack.Current().Symbols().InsertSymbol(v)
	}
	return v
}

/* Apply a numeric math function, given by name, to a numeric argument.
 */
func (intp *PMMPostInterpreter) mathfunc(n dec.Decimal, fun string) dec.Decimal {
	switch fun {
	case "floor":
		n = n.Floor()
	case "ceil":
		n = n.Ceil()
	case "sqrt":
		T.P("mathf", fun).Error("function not yet implemented")
	}
	return n
}

// === Show Commands =========================================================

/* Show all declarations and references for a tag.
 */
func (intp *PMMPostInterpreter) Showvariable(tag string) string {
	sym, scope := intp.scopeTree.Current().ResolveSymbol(tag)
	if sym == nil {
		//T.P("symbol", tag).Debug("no declaration found for symbol")
		return fmt.Sprintf("%s : tag\n", tag)
	} else {
		v := sym.(*PMMPVarDecl)
		var b *bytes.Buffer
		b = v.showDeclarations(b)
		//vname := v.BaseTag.GetName()
		//fmt.Print(b.String())
		// now find all incarnations in top memory-frame(scope)
		if mf := intp.memFrameStack.FindMemoryFrameWithScope(scope); mf != nil {
			for _, v := range mf.Symbols().Table {
				vref := v.(*PMMPVarRef)
				if vref.decl.BaseTag == sym {
					s := fmt.Sprintf("%s = %s\n", vref.GetFullName(), vref.ValueString())
					b.WriteString(s)
				}
			}
		}
		return b.String()
	}
}