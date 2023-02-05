/*
 *  Brown University, CS138, Spring 2018
 *
 *  Purpose: Defines the RoutingTable type and provides methods for interacting
 *  with it.
 */

package tapestry

// import (
// 	"testing"
// )

// func doPublish(data []PublishSpec, t *Testing) {
// 	for _, d := range data {
// 		err := d.store.Store(d.key, []byte{})
// 		cc.Assert(err, IsNil)
// 		time.Sleep(10 * time.Millisecond)
// 	}
// }
//
// func doLookup(data []PublishSpec, t *Testing) {
// 	for _, d := range data {
// 		nodes, err := d.lookup.Lookup(d.key)
// 		cc.Assert(err, IsNil)
// 		cc.Assert(hasnode(nodes, d.store.node), Equals, true)
// 		time.Sleep(10 * time.Millisecond)
// 	}
// }
//
// func doPublishAddLookup(count, numtests int, t *Testing) {
// 	seed := int64(1)
//
//   r := rand.New(rand.NewSource(seed))
//
//   // start tapestry nodes
//   ts := make([]*Node, 0, count)
//
//   for i := 0; i < count; i++ {
//     connectTo := ""
//     if i > 0 {
//       connectTo = ts[0].node.Address
//     }
//     t, err := start(IntToID(r.Int()), 0, connectTo)
//     if err != nil {
//       t.Errorf(err)
//       return
//     }
//     // registerCachedTapestry(t)
//     ts = append(ts, t)
//     time.Sleep(10 * time.Millisecond)
//   }
//
// 	time.Sleep(1 * time.Second)
//
//   // publish and lookup data
// 	// data := GenerateData(aseed, numtests, ts)
//   for i := 0; i < count; i++ {
//     store := ts[r.Intn(len(ts))]
//     lookup := ts[r.Intn(len(ts))]
//     key := fmt.Sprintf("%v-%v-%v", i, storeI, lookupI)
//
//     // specs = append(specs, PublishSpec{store, lookup, key})
//     err := d.store.Store(d.key, []byte{})
//     if err != nil {
//       t.Errorf(err)
//       return
//     }
//     time.Sleep(10 * time.Millisecond)
//
//     nodes, err := d.lookup.Lookup(d.key)
//     if err != nil {
//       t.Errorf(err)
//       return
//     }
//     set := make(map[RemoteNode]struct{}, len(slice))
//     for _, s := range slice {
//       set[s] = struct{}{}
//     }
//
//     _, ok := set[item]
//     if !ok {
//       t.Errorf("Couldn't find node")
//     }
//     // cc.Assert(hasnode(nodes, d.store.node), Equals, true)
//     time.Sleep(10 * time.Millisecond)
//   }
//
// 	// doPublish(data, t)
//
//
// 	time.Sleep(1 * time.Second)
//
// 	ts, err = MakeMoreRandomTapestries(aseed, count, ts)
//
// 	time.Sleep(1 * time.Second)
//
// 	// doLookup(data, t)
//
// }
