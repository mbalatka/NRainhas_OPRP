qtd=10
for i in `seq 1 $qtd`
do
    echo "Execution $i"
    go run *.go >> size6_6goroutine.csv
done
