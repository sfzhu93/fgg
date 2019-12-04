package main;
type Bool interface { Not() Bool; Equal(that Bool) Bool };
type TT struct {};
type FF struct {};
func (this TT) Not() Bool { return FF{} };
func (this FF) Not() Bool { return TT{} };
func (this TT) Equal(that Bool) Bool { return that };
func (this FF) Equal(that Bool) Bool { return that.Not() };
type Nat interface { Add(n Nat) Nat; Equal(n Nat) Bool; equalZero() Bool; equalSucc(m Nat) Bool };
type Zero struct {};
type Succ struct { pred Nat };
func (m Zero) Add(n Nat) Nat { return n };
func (m Succ) Add(n Nat) Nat { return Succ{m.pred.Add(n)} };
func (m Zero) Equal(n Nat) Bool { return n.equalZero() };
func (m Succ) Equal(n Nat) Bool { return n.equalSucc(m.pred) };
func (n Zero) equalZero() Bool { return TT{} };
func (n Succ) equalZero() Bool { return FF{} };
func (n Zero) equalSucc(m Nat) Bool { return FF{} };
func (n Succ) equalSucc(m Nat) Bool { return m.Equal(n.pred) };
type FuncNatNat interface { Apply(x Nat) Nat };
type incr struct { n Nat };
func (this incr) Apply(x Nat) Nat { return x.Add(this.n) };
type composeNatNatNat struct { f FuncNatNat; g FuncNatNat };
func (this composeNatNatNat) Apply(x Nat) Nat { return this.g.Apply(this.f.Apply(x)) };
type D struct {};
func (d D) _1() Nat { return Succ{Zero{}} };
func (d D) _2() Nat { return D{}._1().Add(D{}._1()) };
func (d D) _3() Nat { return D{}._2().Add(D{}._1()) };
func main() { _ = composeNatNatNat{incr{D{}._1()}, incr{D{}._2()}}.Apply(D{}._3()).Add(Zero{}) }