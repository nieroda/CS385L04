#!/bin/sh

replicas=(3 6 9 12 15)

concurrency=(50, 100, 500, 1000, 2000, 5000, 10000, 15000)
requests=(10000, 10000, 10000, 10000, 10000, 20000, 30000, 45000)
ip="35.197.24.73"

echo "1 replica"

for var in {0..4}
do

	for num in {0..8}
	do
	#	sleep 1
		echo "Requests ${requests[num]} ..... Concurrency ${concurrency[num]}"
		ab -q -p payload.json -T application/json -m POST -n ${requests[num]}  -s 250 -c ${concurrency[num]} -r http://${ip}:8080/fibonacci | awk '/Requests per second:/ { print } /99%/ { print } /100%/ { print }'
		sleep 1
	done
	
	kubectl scale deployment finbonacci --replicas ${replicas[var]}
	echo "${replicas[var]} replicas  sleeping....\n\n\n"
	sleep 30
	kubectl get replicaset

done
