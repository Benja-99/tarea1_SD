run_server:
	go run server/server.go
run_client:
	go run client/client.go
run_lider:
	go run lider/lider.go
run_pozo:
	go run pozo_ser/pozo_ser.go
run_namenode:
	go run namenode/namenode.go
run_data1:
	go run data1/data1.go
run_data2:
	go run data2/data2.go
run_data3:
	go run data3/data3.go
run_docker:
	sudo docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management