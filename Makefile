all:
	go install
	cd cmd/ruidgen; go install
	cd cmd/ruid2gen; go install
	cd cmd/ruid3gen; go install
	cd cmd/huidgen; go install
	cd cmd/tuidgen; go install
