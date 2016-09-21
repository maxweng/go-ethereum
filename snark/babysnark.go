package babysnark

// Compile with the following command
// CC=g++-6 CXX=g++-6 go build github.com/ethereum/go-ethereum/snark

/*
#cgo CPPFLAGS: -Wfatal-errors -Wno-deprecated-declarations -DCURVE_ALT_BN128
#cgo CPPFLAGS: -I./libsnark/src/ -I/usr/local/include/c++/6.2.0/ -I/usr/local/include/c++/6.2.0/x86_64-apple-darwin14.5.0/
#cgo LDFLAGS: -lgmpxx -lgmp -lsnark -lsodium -L./libsnark/

extern bool hackishlibsnarkbindings_verify(unsigned char *vk_bytes, unsigned int vk_size, unsigned char *proof_bytes, unsigned int proof_size, unsigned char *primary_input_bytes, unsigned int primary_input_len);
*/
import "C"
import "unsafe"

func babySnarkVerify(vk, proof, primary []byte) {
	C.hackishlibsnarkbindings_verify((*C.unsignedchar)(unsafe.Pointer(&vk[0])), vk.len(), (*C.unsignedchar)(unsafe.Pointer(&proof[0])), proof.len(), (*C.unsignedchar)(unsafe.Pointer(&vprimary[0])), primary.len())
}
