# go_package must be like "a/b/c/kitex_gen/d/e"
kitex:
	kitex -service=p.s.m -I=idl -I=idl/psm --module=a/b/c --protobuf-plugin=validator:module=a/b/c,func=fix_length=idl/psm/fix_length.txt,recurse=true:. idl/psm/psm.proto

protoc:
	protoc -I=idl -I=idl/psm --validator_out=. --validator_opt=module=a/b/c,func=fix_length=idl/psm/fix_length.txt,recurse=true idl/psm/psm.proto