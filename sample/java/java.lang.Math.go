package java

import "github.com/juntaki/jnigo"

var jvm *jnigo.JVM

func init() {
	jvm = jnigo.CreateJVM()
}

type Math struct{
	jclass *jnigo.JClass
}

func NewMath(args ...interface{}) (*Math, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil{
		return nil, err
	}
	jclass, err := jvm.NewJClass("java/lang/Math", convertedArgs)
	if err != nil{
		return nil, err
	}
	return &Math{
		jclass: jclass,
	}, nil
}






func MathIEEEremainder(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"DD":"(DD)D"}
	return jvm.CallStaticFunction("java/lang/Math", "IEEEremainder", sigMap[sigArgs], convertedArgs)
}

func Mathabs(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"J":"(J)J", "F":"(F)F", "D":"(D)D", "I":"(I)I"}
	return jvm.CallStaticFunction("java/lang/Math", "abs", sigMap[sigArgs], convertedArgs)
}

func Mathacos(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "acos", sigMap[sigArgs], convertedArgs)
}

func MathaddExact(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"II":"(II)I", "JJ":"(JJ)J"}
	return jvm.CallStaticFunction("java/lang/Math", "addExact", sigMap[sigArgs], convertedArgs)
}

func Mathasin(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "asin", sigMap[sigArgs], convertedArgs)
}

func Mathatan(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "atan", sigMap[sigArgs], convertedArgs)
}

func Mathatan2(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"DD":"(DD)D"}
	return jvm.CallStaticFunction("java/lang/Math", "atan2", sigMap[sigArgs], convertedArgs)
}

func Mathcbrt(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "cbrt", sigMap[sigArgs], convertedArgs)
}

func Mathceil(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "ceil", sigMap[sigArgs], convertedArgs)
}

func MathcopySign(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"DD":"(DD)D", "FF":"(FF)F"}
	return jvm.CallStaticFunction("java/lang/Math", "copySign", sigMap[sigArgs], convertedArgs)
}

func Mathcos(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "cos", sigMap[sigArgs], convertedArgs)
}

func Mathcosh(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "cosh", sigMap[sigArgs], convertedArgs)
}

func MathdecrementExact(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"I":"(I)I", "J":"(J)J"}
	return jvm.CallStaticFunction("java/lang/Math", "decrementExact", sigMap[sigArgs], convertedArgs)
}

func Mathexp(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "exp", sigMap[sigArgs], convertedArgs)
}

func Mathexpm1(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "expm1", sigMap[sigArgs], convertedArgs)
}

func Mathfloor(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "floor", sigMap[sigArgs], convertedArgs)
}

func MathfloorDiv(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"JJ":"(JJ)J", "II":"(II)I"}
	return jvm.CallStaticFunction("java/lang/Math", "floorDiv", sigMap[sigArgs], convertedArgs)
}

func MathfloorMod(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"II":"(II)I", "JJ":"(JJ)J"}
	return jvm.CallStaticFunction("java/lang/Math", "floorMod", sigMap[sigArgs], convertedArgs)
}

func MathgetExponent(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"F":"(F)I", "D":"(D)I"}
	return jvm.CallStaticFunction("java/lang/Math", "getExponent", sigMap[sigArgs], convertedArgs)
}

func Mathhypot(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"DD":"(DD)D"}
	return jvm.CallStaticFunction("java/lang/Math", "hypot", sigMap[sigArgs], convertedArgs)
}

func MathincrementExact(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"J":"(J)J", "I":"(I)I"}
	return jvm.CallStaticFunction("java/lang/Math", "incrementExact", sigMap[sigArgs], convertedArgs)
}

func Mathlog(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "log", sigMap[sigArgs], convertedArgs)
}

func Mathlog10(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "log10", sigMap[sigArgs], convertedArgs)
}

func Mathlog1p(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "log1p", sigMap[sigArgs], convertedArgs)
}

func Mathmax(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"JJ":"(JJ)J", "FF":"(FF)F", "DD":"(DD)D", "II":"(II)I"}
	return jvm.CallStaticFunction("java/lang/Math", "max", sigMap[sigArgs], convertedArgs)
}

func Mathmin(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"FF":"(FF)F", "DD":"(DD)D", "II":"(II)I", "JJ":"(JJ)J"}
	return jvm.CallStaticFunction("java/lang/Math", "min", sigMap[sigArgs], convertedArgs)
}

func MathmultiplyExact(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"II":"(II)I", "JJ":"(JJ)J"}
	return jvm.CallStaticFunction("java/lang/Math", "multiplyExact", sigMap[sigArgs], convertedArgs)
}

func MathnegateExact(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"I":"(I)I", "J":"(J)J"}
	return jvm.CallStaticFunction("java/lang/Math", "negateExact", sigMap[sigArgs], convertedArgs)
}

func MathnextAfter(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"DD":"(DD)D", "FD":"(FD)F"}
	return jvm.CallStaticFunction("java/lang/Math", "nextAfter", sigMap[sigArgs], convertedArgs)
}

func MathnextDown(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D", "F":"(F)F"}
	return jvm.CallStaticFunction("java/lang/Math", "nextDown", sigMap[sigArgs], convertedArgs)
}

func MathnextUp(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D", "F":"(F)F"}
	return jvm.CallStaticFunction("java/lang/Math", "nextUp", sigMap[sigArgs], convertedArgs)
}

func Mathpow(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"DD":"(DD)D"}
	return jvm.CallStaticFunction("java/lang/Math", "pow", sigMap[sigArgs], convertedArgs)
}

func Mathrandom(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"":"()D"}
	return jvm.CallStaticFunction("java/lang/Math", "random", sigMap[sigArgs], convertedArgs)
}

func Mathrint(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "rint", sigMap[sigArgs], convertedArgs)
}

func Mathround(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)J", "F":"(F)I"}
	return jvm.CallStaticFunction("java/lang/Math", "round", sigMap[sigArgs], convertedArgs)
}

func Mathscalb(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"DI":"(DI)D", "FI":"(FI)F"}
	return jvm.CallStaticFunction("java/lang/Math", "scalb", sigMap[sigArgs], convertedArgs)
}

func Mathsignum(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D", "F":"(F)F"}
	return jvm.CallStaticFunction("java/lang/Math", "signum", sigMap[sigArgs], convertedArgs)
}

func Mathsin(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "sin", sigMap[sigArgs], convertedArgs)
}

func Mathsinh(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "sinh", sigMap[sigArgs], convertedArgs)
}

func Mathsqrt(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "sqrt", sigMap[sigArgs], convertedArgs)
}

func MathsubtractExact(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"JJ":"(JJ)J", "II":"(II)I"}
	return jvm.CallStaticFunction("java/lang/Math", "subtractExact", sigMap[sigArgs], convertedArgs)
}

func Mathtan(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "tan", sigMap[sigArgs], convertedArgs)
}

func Mathtanh(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "tanh", sigMap[sigArgs], convertedArgs)
}

func MathtoDegrees(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "toDegrees", sigMap[sigArgs], convertedArgs)
}

func MathtoIntExact(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"J":"(J)I"}
	return jvm.CallStaticFunction("java/lang/Math", "toIntExact", sigMap[sigArgs], convertedArgs)
}

func MathtoRadians(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D"}
	return jvm.CallStaticFunction("java/lang/Math", "toRadians", sigMap[sigArgs], convertedArgs)
}

func Mathulp(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := map[string]string{"D":"(D)D", "F":"(F)F"}
	return jvm.CallStaticFunction("java/lang/Math", "ulp", sigMap[sigArgs], convertedArgs)
}

