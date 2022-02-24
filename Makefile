run:
	go run .

brun:
	go build .
	./hashcode

clean:
	rm -rf hashcode

ultraclean: clean
	rm -rf result/*.out
