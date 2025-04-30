.PHONY: gen-server-suyiiyii
gen-server-suyiiyii: ## gen service code of {svc}. example: make gen-server svc=product 不加这个--pass会生成一个kitex_gen
	@cd app/${svc} && cwgo server --type RPC --service ${svc} --module github.com/doutokk/doutok/app/${svc}  -I ../../idl  --idl ../../idl/${svc}.proto --template https://github.com/suyiiyii/cwgo-template.git


cwgo server -type RPC --service user --module app/user -I ../idl  --idl ../idl/user.proto --template https://github.com/suyiiyii/cwgo-template.git
