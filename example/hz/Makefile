hz:
	hz new -I=idl --idl=idl/psm/psm.proto --mod=a/b/c --protoc-plugins=validator:hz=true,go_mod=a/b/c,func=fix_length=idl/psm/fix_length.txt:. -f

protoc:
	protoc -I=idl -I=idl/psm --validator_out=. --validator_opt=module=a/b/c,func=fix_length=idl/psm/fix_length.txt,recurse=true idl/psm/psm.proto