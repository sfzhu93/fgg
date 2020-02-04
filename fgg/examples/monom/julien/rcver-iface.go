package main;

/* SA{} is an IA(Int()) and SB an IA(Bool()). In a "method-driven
 * mono, the call to CallFunctionBool observes that MyFunction(Int) is called
 * on an IA(Bool) and so all possible MyFunction(Int) would be mono'd. This
 * would potentially mono MyFunction(Int) on SA{} which is not an IA(Bool). */

type Any(type ) interface {};

type Int(type ) struct {};
type Bool(type ) struct {};

type IA(type a Any()) interface{
	MyFunction(type b Any())(x b) a // Instance found: MyFunction<Int>(x Int) Bool
};

type SA(type ) struct {}; // SA <: IA(Int())

// Can't "monomorphise" this method to match "MyFunction<Int>(x Int) Bool"
func (x SA(type )) MyFunction(type b Any())(y b) Int() {return Int(){}};


type SB(type ) struct {}; // SB <: IA(Bool())
func (x SB(type )) MyFunction(type b Any())(y b) Bool() {return Bool(){}};


type Dummy(type ) struct{};

func (x Dummy(type )) CallFunctionBool(type )(y IA(Bool())) Bool() {
	return y.MyFunction(Int())(Int(){})
};

// func (x Dummy(type )) CallFunctionInt(type )(y IA(Int())) Int() {
// 	return y.MyFunction(Int())(Int(){})
// };

// type Pair(type a Any(), b Any() ) struct { 
// 	fst a;
// 	snd b
// };

func main() { _ =
	Dummy(){}.CallFunctionBool()(SB(){})

	// Pair(Int(),Bool()){Dummy(){}.CallFunctionInt()(SA(){}), 
	// 	   Dummy(){}.CallFunctionBool()(SB(){})
	// 	}

	// Dummy(){}.CallFunctionInt()(SA(){})

	// Pair(Bool(),Int()){
	// 	Dummy(){}.CallFunctionBool()(SB(){}),
	// 	SA(){}.MyFunction(Int())(Int(){})}


	// Pair(Bool(),IA(Int())){
	// 	Dummy(){}.CallFunctionBool()(SB(){}),
	// 	SA(){}
	// }.snd.MyFunction(Bool())(Bool(){})

}


