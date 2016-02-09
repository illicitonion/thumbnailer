## Example project showing go benchmarking

Example output:

$ GOPATH=$(pwd):$(pwd)/vendor go test -bench=. thumbmailer
testing: warning: no tests to run
PASS
BenchmarkVips-8	      30	  51015198 ns/op
BenchmarkNFNT-8	       5	 273560975 ns/op
ok  	thumbmailer	4.076s
