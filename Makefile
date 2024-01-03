update-readme:
	cd ./script && \
	ACTION=update-readme go run main.go
create-exam:
	cd ./script && \
	ACTION=create-exam go run main.go