package babysnark

/*
#cgo CXXFLAGS: -std=c++11
#cgo CFLAGS: -I./libsnark/src/ -I/usr/include/c++/4.2.1/ -I/usr/include/c++/4.2.1/tr1
#cgo LDFLAGS: -lgmp
#include "./hackishlibsnarkbindings.cpp"
*/
import "C"

func babySnarkVerify(vk, proof, primary []byte) {
	C.hackishlibsnarkbindings_verify(vk, len(vk), proof, len(proof), primary, len(primary))
}
