/*
 *  Brown University, CS138, Spring 2017
 *
 *  Purpose: Defines the RoutingTable type and provides methods for interacting
 *  with it.
 */

package tapestry
/*
import (
	"testing"
)

func SharedPrefixLength1_Test(t *testing.T) {
	a, _ := PartialID("1")
	b, _ := PartialID("2")

	p := SharedPrefixLength(a, b)
	cc.Assert(p, Equals, 0)
}

func SharedPrefixLength2_Test(t *testing.T) {
	a, _ := PartialID("519BFDD559CB03792ECD7565AEA263D08F3CF528")
	b, _ := PartialID("519BFDD559CB03792ECD7565AEA263D08F3CF528")

	p := SharedPrefixLength(a, b)

	cc.Assert(p, Equals, 40)
}

func BetterChoice2_Test(t *testing.T) {
	a, _ := PartialID("00000007")
	b, _ := PartialID("00000007")
	c, _ := PartialID("00000007")

	better := a.BetterChoice(b, c)

	cc.Assert(better, Equals, false)
}

func BetterChoice5_Test(t *testing.T) {
	a, _ := PartialID("23257")
	b, _ := PartialID("23251")
	c, _ := PartialID("23256")

	better := a.BetterChoice(b, c)

	cc.Assert(better, Equals, true)
}

func BetterChoice6_Test(t *testing.T) {
	a, _ := PartialID("1234567")
	b, _ := PartialID("1234561")
	c, _ := PartialID("1234562")

	better := a.BetterChoice(b, c)

	cc.Assert(better, Equals, true)
}

func Closer3_Test(t *testing.T) {
	a, _ := PartialID("5555555")
	b, _ := PartialID("555555F")
	c, _ := PartialID("5555554")

	closer := a.Closer(c, b)

	cc.Assert(closer, Equals, true)
}

func Closer4_Test(t *testing.T) {
	a, _ := PartialID("5555555")
	b, _ := PartialID("5555559")
	c, _ := PartialID("5555554")

	closer := a.Closer(b, c)

	cc.Assert(closer, Equals, false)
}

func Closer5_Test(t *testing.T) {
	a, _ := PartialID("0FFFFFFF")
	b, _ := PartialID("0FFFFFFD")
	c, _ := PartialID("10000000")

	closer := a.Closer(b, c)

	cc.Assert(closer, Equals, false)
}

func Closer6_Test(t *testing.T) {
	a, _ := PartialID("0FFFFFFF")
	b, _ := PartialID("0FFFFFFD")
	c, _ := PartialID("10000000")

	closer := a.Closer(c, b)

	cc.Assert(closer, Equals, true)
}

func Closer7_Test(t *testing.T) {
	a, _ := PartialID("0FFFFFFF")
	b, _ := PartialID("0FFFFFFE")
	c, _ := PartialID("10000000")

	closer := a.Closer(c, b)

	cc.Assert(closer, Equals, false)
}

func Closer8_Test(t *testing.T) {
	a, _ := PartialID("0FFFFFFF")
	b, _ := PartialID("0FFFFFFE")
	c, _ := PartialID("10000000")

	closer := a.Closer(b, c)

	cc.Assert(closer, Equals, false)
}
*/