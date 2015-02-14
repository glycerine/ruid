all:
	go install
	cd cmd/ruidgen; go install
	cd cmd/ruid2gen; go install
