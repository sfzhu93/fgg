package fgg

import (
	"fmt"
	"reflect"
	"strings"
)

var _ = fmt.Errorf
var _ = reflect.Append
var _ = strings.Compare

// Return true if *not* nomono
func IsMonomOK(p FGGProgram) (bool, string) {
	ds := p.GetDecls()
	for _, v := range ds {
		if md, ok := v.(MethDecl); ok {
			omega1 := Omega2{make(map[string]Type), make(map[string]MethInstan2)}
			delta := md.Psi_recv.ToDelta()
			for _, v := range md.Psi_meth.tFormals {
				delta[v.name] = v.u_I
			}
			gamma := make(Gamma)
			psi_recv := make(SmallPsi, len(md.Psi_recv.tFormals))
			for i, v := range md.Psi_recv.tFormals {
				psi_recv[i] = v.name
			}
			//psi_recv = md.Psi_recv.Hat()
			u_recv := TNamed{md.t_recv, psi_recv}
			gamma[md.x_recv] = u_recv
			omega1.us[toKey_Wt(u_recv)] = u_recv
			for _, v := range md.pDecls { // TODO: factor out
				gamma[v.name] = v.u
			}
			collectExpr2(ds, delta, gamma, md.e_body, omega1)
			if ok, msg := nomonoOmega(ds, delta, md, omega1); ok {
				return false, msg
			}
		}
	}
	return true, ""
}

// Return true if nomono
func nomonoOmega(ds []Decl, delta Delta, md MethDecl, omega1 Omega2) (bool, string) {
	for auxG2(ds, delta, omega1) {
		for _, v := range omega1.ms {
			if !isStructType(ds, v.u_recv) {
				continue
			}
			u_S := v.u_recv.(TNamed)
			if u_S.t_name == md.t_recv && v.meth == md.name {
				if occurs(md.Psi_recv, u_S.u_args) {
					return true, md.t_recv + md.Psi_recv.String() + " ->* " + md.t_recv +
						"(" + SmallPsi(u_S.u_args).String() + ")"
				}
				if occurs(md.Psi_meth, v.psi) {
					return true, md.t_recv + md.Psi_recv.String() + "." + md.name +
						md.Psi_meth.String() + " ->* " + md.name + "(" + v.psi.String() + ")"
				}
			}
		}
	}
	return false, ""
}

// Pre: len(Psi) == len(psi)
func occurs(Psi BigPsi, psi SmallPsi) bool {
	for i, v := range Psi.tFormals {
		if cast, ok := psi[i].(TNamed); ok { // !!! simplified
			for _, x := range fv(cast) {
				if x.Equals(v.name) {
					return true
				}
			}
		}
	}
	return false
}

func fv(u Type) []TParam {
	if cast, ok := u.(TParam); ok {
		return []TParam{cast}
	}
	res := []TParam{}
	cast := u.(TNamed)
	for _, v := range cast.u_args {
		res = append(res, fv(v)...)
	}
	return res
}

type MethInstan2 struct {
	u_recv Type
	meth   Name
	psi    SmallPsi
}

type Omega2 struct {
	us map[string]Type
	ms map[string]MethInstan2
}

func (w Omega2) clone() Omega2 {
	us := make(map[string]Type)
	ms := make(map[string]MethInstan2)
	for k, v := range w.us {
		us[k] = v
	}
	for k, v := range w.ms {
		ms[k] = v
	}
	return Omega2{us, ms}
}

func (w Omega2) Println() {
	fmt.Println("=== Type instances:")
	for _, v := range w.us {
		fmt.Println(v)
	}
	fmt.Println("--- Method instances:")
	for _, v := range w.ms {
		fmt.Println(v.u_recv, v.meth, v.psi)
	}
	fmt.Println("===")
}

func toKey_Wt2(u Type) string {
	return u.String()
}

func toKey_Wm2(x MethInstan2) string {
	return x.u_recv.String() + "_" + x.meth + "_" + x.psi.String()
}

// TODO: refactor with above
func collectExpr2(ds []Decl, delta Delta, gamma Gamma, e FGGExpr, omega Omega2) bool {
	res := false
	switch e1 := e.(type) {
	case Variable:
		return res
	case StructLit:
		for _, elem := range e1.elems {
			res = collectExpr2(ds, delta, gamma, elem, omega) || res
		}
		k := toKey_Wt2(e1.u_S)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = e1.u_S
			res = true
		}
	case Select:
		return collectExpr2(ds, delta, gamma, e1.e_S, omega)
	case Call:
		res = collectExpr2(ds, delta, gamma, e1.e_recv, omega) || res
		for _, e_arg := range e1.args {
			res = collectExpr2(ds, delta, gamma, e_arg, omega) || res
		}
		gamma1 := make(Gamma)
		for k, v := range gamma {
			gamma1[k] = v
		}
		u_recv := e1.e_recv.Typing(ds, delta, gamma1, false)
		k_t := toKey_Wt2(u_recv)
		if _, ok := omega.us[k_t]; !ok {
			omega.us[k_t] = u_recv
			res = true
		}
		m := MethInstan2{u_recv, e1.meth, e1.GetTArgs()} // CHECKME: why add u_recv separately?
		k_m := toKey_Wm2(m)
		if _, ok := omega.ms[k_m]; !ok {
			omega.ms[k_m] = m
			res = true
		}
	case Assert:
		res = collectExpr2(ds, delta, gamma, e1.e_I, omega) || res
		u := e1.u_cast
		k := toKey_Wt2(u)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = u
			res = true
		}
	case String: // CHECKME
		k := toKey_Wt2(STRING_TYPE)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = STRING_TYPE
			res = true // CHECKME
		}
	case Sprintf:
		k := toKey_Wt2(STRING_TYPE)
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = STRING_TYPE
			res = true
		}
		for _, arg := range e1.args {
			res = collectExpr2(ds, delta, gamma, arg, omega) || res
		}
	default:
		panic("Unknown Expr kind: " + reflect.TypeOf(e).String() + "\n\t" +
			e.String())
	}
	return res
}

/* Aux */

// Return true if omega has changed
// N.B. no closure over types occurring in bounds, or *interface decl* method sigs
//func auxG(ds []Decl, omega Omega1) bool {
func auxG2(ds []Decl, delta Delta, omega Omega2) bool {
	res := false
	res = auxF2(ds, omega) || res
	res = auxI2(ds, delta, omega) || res
	res = auxM2(ds, delta, omega) || res
	res = auxS2(ds, delta, omega) || res
	// I/face embeddings
	res = auxE12(ds, omega) || res
	res = auxE22(ds, omega) || res
	//res = auxP(ds, omega) || res
	return res
}

func auxF2(ds []Decl, omega Omega2) bool {
	//func auxF(ds []Decl, delta Delta, omega Omega1) bool {
	res := false
	tmp := make(map[string]Type)
	for _, u := range omega.us {
		if !isStructType(ds, u) || u.Equals(STRING_TYPE) { // CHECKME
			continue
		}
		for _, u_f := range Fields(ds, u.(TNamed)) {
			cast := u_f.u
			tmp[toKey_Wt2(cast)] = cast
		}
	}
	for k, v := range tmp {
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = v
			res = true
		}
	}
	return res
}

func auxI2(ds []Decl, delta Delta, omega Omega2) bool {
	res := false
	tmp := make(map[string]MethInstan2)
	for _, m := range omega.ms {
		if !IsNamedIfaceType(ds, m.u_recv) {
			continue
		}
		for _, m1 := range omega.ms {
			if !IsNamedIfaceType(ds, m1.u_recv) {
				continue
			}
			if m1.u_recv.ImplsDelta(ds, delta, m.u_recv) {
				mm := MethInstan2{m1.u_recv, m.meth, m.psi}
				tmp[toKey_Wm2(mm)] = mm
			}
		}
	}
	for k, v := range tmp {
		if _, ok := omega.ms[k]; !ok {
			omega.ms[k] = v
			res = true
		}
	}
	return res
}

func auxM2(ds []Decl, delta Delta, omega Omega2) bool {
	res := false
	tmp := make(map[string]Type)
	for _, m := range omega.ms {
		gs := methodsDelta(ds, delta, m.u_recv)
		for _, g := range gs { // Should be only g s.t. g.meth == m.meth
			if g.meth != m.meth {
				continue
			}
			eta := MakeEta2(g.Psi, m.psi)
			for _, pd := range g.pDecls {
				u_pd := pd.u.SubsEta2(eta) // HERE: need receiver subs also? cf. map.fgg "type b Eq(b)" -- methods should be ok?
				tmp[toKey_Wt2(u_pd)] = u_pd
			}
			u_ret := g.u_ret.SubsEta2(eta)
			tmp[toKey_Wt2(u_ret)] = u_ret
		}
	}
	for k, v := range tmp {
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = v
			res = true
		}
	}
	return res
}

func auxS2(ds []Decl, delta Delta, omega Omega2) bool {
	res := false
	tmp := make(map[string]MethInstan2)
	clone := omega.clone()
	for _, m := range clone.ms {
		for _, u := range clone.us {
			u_recv := bounds(delta, m.u_recv) // !!! cf. plain type param
			if !isStructType(ds, u) || !u.ImplsDelta(ds, delta, u_recv) {
				continue
			}
			u_S := u.(TNamed)
			x0, xs, e := body(ds, u_S, m.meth, m.psi)
			gamma := make(Gamma)
			gamma[x0.name] = x0.u.(TNamed)
			for _, pd := range xs {
				gamma[pd.name] = pd.u
			}
			m1 := MethInstan2{u_S, m.meth, m.psi}
			k := toKey_Wm2(m1)
			//if _, ok := omega.ms[k]; !ok { // No: initial collectExpr already adds to omega.ms
			tmp[k] = m1
			res = collectExpr2(ds, delta, gamma, e, omega) || res
			//}
		}
	}
	for k, v := range tmp {
		if _, ok := omega.ms[k]; !ok {
			omega.ms[k] = v
			res = true
		}
	}
	return res
}

// Add embedded types
func auxE12(ds []Decl, omega Omega2) bool {
	res := false
	tmp := make(map[string]TNamed)
	for _, u := range omega.us {
		if !isNamedIfaceType(ds, u) { // TODO CHECKME: type param
			continue
		}
		u_I := u.(TNamed)
		td_I := getTDecl(ds, u_I.t_name).(ITypeLit)
		eta := MakeEta2(td_I.Psi, u_I.u_args)
		for _, s := range td_I.specs {
			if u_emb, ok := s.(TNamed); ok {
				u_sub := u_emb.SubsEta2(eta).(TNamed)
				tmp[toKey_Wt2(u_sub)] = u_sub
			}
		}
	}
	for k, v := range tmp {
		if _, ok := omega.us[k]; !ok {
			omega.us[k] = v
			res = true
		}
	}
	return res
}

// Propagate method instances up to embedded supertypes
func auxE22(ds []Decl, omega Omega2) bool {
	res := false
	tmp := make(map[string]MethInstan2)
	for _, m := range omega.ms {
		if !isNamedIfaceType(ds, m.u_recv) { // TODO CHECKME: type param
			continue
		}
		u_I := m.u_recv.(TNamed)
		td_I := getTDecl(ds, u_I.t_name).(ITypeLit)
		eta := MakeEta2(td_I.Psi, u_I.u_args)
		for _, s := range td_I.specs {
			if u_emb, ok := s.(TNamed); ok {
				u_sub := u_emb.SubsEta2(eta).(TNamed)
				gs := methods(ds, u_sub)
				for _, g := range gs {
					if m.meth == g.meth {
						m_emb := MethInstan2{u_sub, m.meth, m.psi}
						tmp[toKey_Wm2(m_emb)] = m_emb
					}
				}
			}
		}
	}
	for k, v := range tmp {
		if _, ok := omega.ms[k]; !ok {
			omega.ms[k] = v
			res = true
		}
	}
	return res
}

/*

















































 */

/* Deprecated: Old CFG-based check */

type RecvMethPair struct {
	t_recv Name // Pre: t_S
	m      Name // TODO rename
}

func (x0 RecvMethPair) equals(x RecvMethPair) bool {
	return x0.t_recv == x.t_recv && x0.m == x.m
}

type cTypeArgs struct {
	psi_recv SmallPsi
	psi_meth SmallPsi
}

func (x0 cTypeArgs) equals(x cTypeArgs) bool {
	return x0.psi_recv.Equals(x.psi_recv) && x0.psi_meth.Equals(x.psi_meth)
}

// Static call graph, agnostic of specific type args (cf. MethInstan)
// N.B. nodes are for struct types
type cgraph struct {
	edges map[RecvMethPair]map[RecvMethPair]([]cTypeArgs)
}

func (x0 cgraph) String() string {
	var b strings.Builder
	for k, v := range x0.edges {
		b.WriteString(k.t_recv)
		b.WriteString(".")
		b.WriteString(k.m)
		b.WriteString(": ")
		b.WriteString(fmt.Sprintf("%v", v))
		b.WriteString("\n")
	}
	return b.String()
}

// CFG-based occurs check -- needs method set unification ("open" impls)
// CHECKME: generally, covariant receiver bounds specialisation
func IsMonomOK_CFG(p FGGProgram) bool {
	ds := p.GetDecls()
	graph := cgraph{make(map[RecvMethPair]map[RecvMethPair]([]cTypeArgs))}
	for _, v := range ds {
		if md, ok := v.(MethDecl); ok {
			buildGraph(ds, md, graph)
		}
	}
	//buildGraphExpr(ds, make(Delta), make(Gamma), ...)  // visit main unnecessary -- CHECKME: all type instans seen?
	//fmt.Println("111:\n", graph.String(), "---")
	cycles := make(map[string]cycle)
	findCycles(graph, cycles)
	/*for _, v := range cycles {
		fmt.Println("aaa:", v)
	}*/
	for _, v := range cycles {
		//fmt.Println("bbb:", v)
		if isNomonoCycle(ds, graph, v) {
			return false
		}
		return true
	}
	return true
}

// Occurs check -- N.B. conservative w.r.t. whether type params actually used
func isNomonoCycle(ds []Decl, graph cgraph, c cycle) bool {
	for _, tArgs := range graph.edges[c[0]][c[1]] {
		if isNomonoTypeArgs(tArgs) || isNomonoCycleAux(ds, graph, c, tArgs, 1) {
			return true
		}
	}
	return false
}

func isNomonoTypeArgs(tArgs cTypeArgs) bool {
	for _, v := range tArgs.psi_recv {
		if containsNestedTParam(v) {
			return true
		}
	}
	for _, v := range tArgs.psi_meth {
		if containsNestedTParam(v) {
			return true
		}
	}
	return false
}

func isNomonoCycleAux(ds []Decl, graph cgraph, c cycle, tArgs cTypeArgs, i int) bool {
	if i >= (len(c) - 1) {
		return false
	}
	next := c[i]
	md := getMDecl(ds, next.t_recv, next.m)
	subs := make(Delta)
	for i, v := range tArgs.psi_recv {
		subs[md.Psi_recv.tFormals[i].name] = v
	}
	for i, v := range tArgs.psi_meth {
		subs[md.Psi_meth.tFormals[i].name] = v
	}

	for _, v := range graph.edges[c[i]][c[i+1]] {
		tArgs1 := cTypeArgs{v.psi_recv.TSubs(subs), v.psi_meth.TSubs(subs)}
		if isNomonoTypeArgs(tArgs1) {
			return true
		}
		isNomonoCycleAux(ds, graph, c, tArgs1, i+1)
	}
	return false
}

func getMDecl(ds []Decl, t_recv Name, meth Name) MethDecl {
	for _, v := range ds {
		if md, ok := v.(MethDecl); ok && md.t_recv == t_recv && md.name == meth {
			return md
		}
	}
	panic("MethDecl not found: " + t_recv + "." + meth)
}

func containsNestedTParam(u Type) bool {
	if cast, ok := u.(TNamed); ok {
		for _, v := range cast.u_args {
			if isOrContainsTParam(v) {
				return true
			}
		}
	}
	return false
}

type cycle []RecvMethPair

func (x0 cycle) toHash() string {
	return fmt.Sprintf("%v", x0)
}

func findCycles(graph cgraph, cycles map[string]cycle) {
	for k, _ := range graph.edges {
		stack := []RecvMethPair{k}
		findCyclesAux(graph, stack, cycles)
	}
}

// DFS -- TODO: start from main more efficient? -- CHECKME: maybe more "correct", w.r.t. omega method discarding
func findCyclesAux(graph cgraph, stack []RecvMethPair, cycles map[string]cycle) {
	targets := graph.edges[stack[len(stack)-1]]
	if targets == nil {
		panic("Shouldn't get in here:")
	}
lab:
	for next, _ := range targets {
		stack1 := append(stack, next)
		if stack1[0].equals(next) {
			cycles[cycle(stack1).toHash()] = stack1
			continue
		}
		for _, prev := range stack[1:] {
			if prev.equals(next) {
				continue lab
			}
		}
		findCyclesAux(graph, stack1, cycles)
	}
}

// "Flat" graph building -- calls not visited (i.e., `body` not used)
// Output: mutates cgraph
func buildGraph(ds []Decl, md MethDecl, graph cgraph) {
	n := RecvMethPair{md.t_recv, md.name}
	graph.edges[n] = make(map[RecvMethPair]([]cTypeArgs))
	delta := md.Psi_meth.ToDelta() // recv params added below
	gamma := make(Gamma)
	psi := make(SmallPsi, len(md.Psi_recv.tFormals))
	for i, v := range md.Psi_recv.tFormals {
		delta[v.name] = v.u_I
		psi[i] = v.name
	}
	gamma[md.x_recv] = TNamed{md.t_recv, psi}
	for _, v := range md.pDecls { // TODO: factor out
		gamma[v.name] = v.u
	}
	buildGraphExpr(ds, delta, gamma, n, md.e_body, graph)
}

// "Flat" graph building -- calls not visited (i.e., `body` not used)
func buildGraphExpr(ds []Decl, delta Delta, gamma Gamma, curr RecvMethPair, e1 FGGExpr, graph cgraph) {
	switch e := e1.(type) {
	case Variable:
	case StructLit:
		for _, elem := range e.elems {
			buildGraphExpr(ds, delta, gamma, curr, elem, graph)
		}
	case Select:
		buildGraphExpr(ds, delta, gamma, curr, e.e_S, graph)
	case Call:
		buildGraphExpr(ds, delta, gamma, curr, e.e_recv, graph)
		for _, arg := range e.args {
			buildGraphExpr(ds, delta, gamma, curr, arg, graph)
		}
		u_recv := e.e_recv.Typing(ds, delta, gamma, true)

		if isStructType(ds, u_recv) { // u_recv is a TNamed struct
			u_S := u_recv.(TNamed)
			putTArgs(graph, curr, u_S, e.meth, e.t_args)

		} else { // TNamed interface or TParam
			u_I := u_recv // Or type param
			if _, ok := u_I.(TParam); ok {
				u_I = u_I.TSubs(delta) // CHECKME
			}
			for _, v := range ds {
				if d, ok := v.(STypeLit); ok {

					// CHECKME: method set unification instead of basic impls? -- or is using bounds (hat) sufficient?
					u_S := TNamed{d.t_name, d.Psi.Hat()} // !!!
					if u_S.ImplsDelta(ds, delta, u_I) {
						putTArgs(graph, curr, u_S, e.meth, e.t_args)
					}

				}
			}
		}
	case Assert:
		buildGraphExpr(ds, delta, gamma, curr, e.e_I, graph)
	default:
		panic("Unknown Expr kind: " + reflect.TypeOf(e1).String() + "\n\t" +
			e1.String())
	}
}

// u_recv is target u_S
func putTArgs(graph cgraph, curr RecvMethPair, u_recv TNamed, meth Name, psi_meth SmallPsi) {
	edges := graph.edges[curr]
	/*if edges == nil {
		edges = make(map[node]([]cTypeArgs))
		graph.edges[curr] = edges
	}*/
	target := RecvMethPair{u_recv.t_name, meth}
	tArgs := edges[target]
	if tArgs == nil {
		tArgs = []cTypeArgs{}
	}
	tArgs = append(tArgs, cTypeArgs{u_recv.u_args, psi_meth})
	edges[target] = tArgs
}

/*



















































 */

type RecvMethPair1 struct {
	u_recv string
	m      Name
}

var meths []RecvMethPair1 = make([]RecvMethPair1, 0) // TODO refactor

// Check for any nested type params in a cycle -- very conservative
func IsMonomOK_naive(ds []Decl) {
	graph := make(map[RecvMethPair1]map[RecvMethPair1]([][]Type))    // caller->callee->[list of meth type args]
	bools := make(map[RecvMethPair1]map[RecvMethPair1]bool)          // caller->callee->true/false (cycle detection convenience) -- this more the actual call-graph
	recvargs := make(map[RecvMethPair1]map[RecvMethPair1]([][]Type)) // caller->callee->[list of receiver type args]
	for _, v := range ds {
		switch d := v.(type) {
		case STypeLit:
		case ITypeLit:
		case MethDecl:
			delta := d.GetMDeclPsi().ToDelta()
			tfs := d.GetRecvPsi().GetTFormals()
			u_args := make([]Type, len(tfs))
			for i := 0; i < len(tfs); i++ {
				//u_args[i] = tfs[i].GetUpperBound()
				u_args[i] = tfs[i].GetTParam()
				delta[tfs[i].GetTParam()] = tfs[i].GetUpperBound()
			}
			u_recv := TNamed{d.t_recv, u_args}
			gamma := make(Gamma)
			gamma[d.x_recv] = u_recv
			for _, v := range d.GetParamDecls() {
				gamma[v.GetName()] = v.GetType()
			}
			//ctxt := RecvMethPair1{u_recv.TSubs(delta).String(), d.name}
			ctxt := RecvMethPair1{u_recv.t_name, d.name}
			meths = append(meths, ctxt)
			bar(ds, delta, gamma, ctxt, d.e_body, graph, bools, recvargs)
		default:
			panic("Unknown Decl kind: " + reflect.TypeOf(v).String() + "\n\t" +
				v.String())
		}
	}

	////war(bools)
	//fmt.Println("1111a: ", graph, "\n")
	//fmt.Println("1111b: ", recvargs)
	//fmt.Println("2222: ", bools)

	findCycles1(bools)
	fmt.Println("3333: ", cycles)

	for _, v := range cycles {
		for i := 0; i < len(v); i++ {
			var next RecvMethPair1
			if i == len(v)-1 {
				next = v[0]
			} else {
				next = v[i+1]
			}
			tmp := graph[v[i]]
			if tmp != nil {
				tmp2 := tmp[next]
				if tmp2 != nil {
					for _, t_args := range tmp2 {
						for _, u := range t_args {
							if u1, ok := u.(TNamed); ok {
								for _, x := range u1.u_args {
									if isOrContainsTParam(x) { // CHECKME: basically the naive syntactic restriction, OK?
										panic("Not monomorphisable, potential polymorphic recursion: " +
											fmt.Sprintf("%v", v))
									}
								}
							}
						}
					}
				}
			}

			rtmp := recvargs[v[i]]
			if rtmp != nil {
				rtmp2 := rtmp[next]
				if rtmp2 != nil {
					for _, t_args := range rtmp2 {
						for _, u := range t_args {
							if u1, ok := u.(TNamed); ok {
								for _, x := range u1.u_args {
									if isOrContainsTParam(x) { // CHECKME: basically the naive syntactic restriction, OK?
										panic("Not monomorphisable, potential polymorphic recursion: " +
											fmt.Sprintf("%v", v))
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

// Populate call graph by visiting Expr
// N.B. mutates graph
func bar(ds []Decl, delta Delta, gamma Gamma, ctxt RecvMethPair1, e FGGExpr,
	graph map[RecvMethPair1]map[RecvMethPair1]([][]Type),
	bools map[RecvMethPair1]map[RecvMethPair1]bool,
	recvargs map[RecvMethPair1]map[RecvMethPair1]([][]Type)) {
	switch e1 := e.(type) {
	case Variable:
	case StructLit:
		for _, elem := range e1.elems {
			bar(ds, delta, gamma, ctxt, elem, graph, bools, recvargs)
		}
	case Select:
		bar(ds, delta, gamma, ctxt, e1.e_S, graph, bools, recvargs)
	case Call:
		bar(ds, delta, gamma, ctxt, e1.e_recv, graph, bools, recvargs)
		for _, arg := range e1.args {
			bar(ds, delta, gamma, ctxt, arg, graph, bools, recvargs)
		}
		//g := methods(u_recv)[e1.meth]  // Want u_recv from Typing...
		/*var psi Psi
		for _, v := range ds {
			if v1, ok := v.(MDecl); ok && v1.name == e1.meth {
				psi = v1.GetMDeclPsi()
				break
			}
		}
		delta1 := psi.ToDelta()
		for k, v := range delta {
			delta1[k] = v
		}*/
		delta1 := delta // TODO refactor
		u_recv := e1.e_recv.Typing(ds, delta1, gamma, true)

		/*if _, ok := u_recv.(TParam); ok { // E.g., compose, x.Equal()(xs.head), x is `a`
			u_recv = delta[u_recv.(TParam)]
		}*/

		tmp := graph[ctxt]
		btmp := bools[ctxt]
		rtmp := recvargs[ctxt]
		if tmp == nil {
			tmp = make(map[RecvMethPair1]([][]Type))
			graph[ctxt] = tmp
			btmp = make(map[RecvMethPair1]bool)
			bools[ctxt] = btmp
			rtmp = make(map[RecvMethPair1]([][]Type))
			recvargs[ctxt] = rtmp
		}
		if isStructType(ds, u_recv) {
			//key := RecvMethPair1{u_recv.TSubs(delta1).String(), e1.meth}
			key := RecvMethPair1{u_recv.TSubs(delta).(TNamed).t_name, e1.meth}
			tmp2 := tmp[key]
			if tmp2 == nil {
				tmp2 = make([][]Type, 0)
			}
			tmp2 = append(tmp2, e1.t_args)
			tmp[key] = tmp2
			btmp[key] = true
			if y, ok := u_recv.(TNamed); ok { // CHECKME: how about TParam?
				rtmp2 := rtmp[key]
				if rtmp2 == nil {
					rtmp2 = make([][]Type, 0)
				}
				rtmp2 = append(rtmp2, y.u_args)
				rtmp[key] = rtmp2
			}
		} else {
			u_I := u_recv // Or type param
			for _, v := range ds {
				switch d := v.(type) {
				case STypeLit:
					tfs := d.GetBigPsi().GetTFormals()
					u_args := make([]Type, len(tfs))
					for i := 0; i < len(tfs); i++ {
						u_args[i] = tfs[i].GetUpperBound()
					}
					u_S := TNamed{d.t_name, u_args}
					if p, ok := u_I.(TParam); (ok && u_S.ImplsDelta(ds, delta1, delta1[p])) || // CHECKME: delta1[p] ?
						(!ok && u_S.ImplsDelta(ds, delta1, u_I)) {
						//key := RecvMethPair1{u_S.TSubs(delta1).String(), e1.meth}
						key := RecvMethPair1{u_S.t_name, e1.meth}
						// TODO factor out below with above
						tmp2 := tmp[key]
						if tmp2 == nil {
							tmp2 = make([][]Type, 0)
						}
						tmp2 = append(tmp2, e1.t_args)
						tmp[key] = tmp2
						btmp[key] = true
						if y, ok := u_recv.(TNamed); ok { // CHECKME: how about TParam?
							rtmp2 := rtmp[key]
							if rtmp2 == nil {
								rtmp2 = make([][]Type, 0)
							}
							rtmp2 = append(rtmp2, y.u_args)
							rtmp[key] = rtmp2
						}
					}
				case ITypeLit:
				case MethDecl:
				default:
					panic("Unknown Decl kind: " + reflect.TypeOf(e).String() + "\n\t" +
						e.String())
				}
			}
		}
	case Assert:
		bar(ds, delta, gamma, ctxt, e1.e_I, graph, bools, recvargs)
	default:
		panic("Unknown Expr kind: " + reflect.TypeOf(e).String() + "\n\t" +
			e.String())
	}
}

/* Aux */

var cycles [][]RecvMethPair1

func findCycles1(bools map[RecvMethPair1]map[RecvMethPair1]bool) {
	for _, v := range meths {
		stack := []RecvMethPair1{v}
		aux(bools, stack)
	}
}

// DFS
func aux(bools map[RecvMethPair1]map[RecvMethPair1]bool, stack []RecvMethPair1) {
	tmp := bools[stack[len(stack)-1]]
	if tmp == nil {
		return
	}
	for i := 0; i < len(meths); i++ {
		m := meths[i]
		if tmp[m] {
			if stack[0] == m {
				//stack1 := append(stack, m)
				cycles = append(cycles, stack)
				return
			}
			for j := 1; j < len(stack); j++ {
				if stack[j] == m {
					return
				}
			}
			stack1 := append(stack, m)
			aux(bools, stack1)
		}
	}
}

/*// Currently redundant
// Mutates graph
func war(graph map[RecvMethPair1]map[RecvMethPair1]bool) {
	for k := 0; k < len(meths); k++ {
		for i := 0; i < len(meths); i++ {
			for j := 0; j < len(meths); j++ {
				tmp := graph[meths[i]]
				if tmp == nil {
					/*tmp = make(map[RecvMethPair1]bool)
					graph[meths[i]] = tmp* /
					return
				}
				if !tmp[meths[j]] {
					tmp2 := graph[meths[i]]
					tmp3 := graph[meths[k]]
					if tmp2 != nil && tmp3 != nil {
						tmp[meths[j]] = tmp2[meths[k]] && tmp3[meths[j]]
					}
				}
			}
		}
	}
}
*/
