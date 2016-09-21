package babysnark

// Compile with the following command
// CC=g++-6 CXX=g++-6 go build github.com/ethereum/go-ethereum/snark

/*
#cgo CPPFLAGS: -Wno-all -DCURVE_ALT_BN128
#cgo CPPFLAGS: -I./libsnark/src/ -I/usr/local/include/c++/6.2.0/ -I/usr/local/include/c++/6.2.0/x86_64-apple-darwin14.5.0/
#cgo CPPFLAGS: -I /usr/local/include/c++/6.2.0/x86_64-apple-darwin15.6.0/
#cgo CPPFLAGS: -I /usr/local/opt/openssl/include/
#cgo LDFLAGS: -lgmpxx -lgmp -lsnark -lsodium -L./libsnark/

#include "hackishlibsnarkbindings.hpp"
extern "C" bool hackishlibsnarkbindings_verify(void*,uint32_t,void*,uint32_t,void*,uint32_t);
*/
import "C"
import "unsafe"

func babySnarkVerify(vk, proof, primary []byte) {
	c_vk := unsafe.Pointer(&vk[0])
	c_proof := unsafe.Pointer(&proof[0])
	c_primary := unsafe.Pointer(&primary[0])
	ret := C.hackishlibsnarkbindings_verify(c_vk, vk.len(), c_proof, proof.len(), c_primary, primary.len())
}
