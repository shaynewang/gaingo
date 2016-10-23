all:
	go build train.go
clean:
	 if [ -f train ] ; then rm train ; fi

