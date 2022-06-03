init:
	go run cmd/model/init.go

gen:
	go run cmd/generate/generate.go

clean:
	rm *.db
