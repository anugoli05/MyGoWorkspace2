package subtract

import (
	"MyPackages"
)
//CallSub is a function to call subtract function of the mygopackage to maintain my own package
func CallSub(a, b int) int{
var Subresult int
Subresult=	mathutils.Subtractt(a,b)
return Subresult
}