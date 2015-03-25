outputFile="coverage.html"
outfile="c.out"

if [ -f "$outputFile" ]
then
	rm -rf "$outputFile"
fi

goapp test -cover -covermode count -test.v=true -test.coverprofile=$outfile
goapp tool cover -html c.out -o $outputFile

if [ -f "$outfile" ]
then
	rm -rf "$outfile"
fi
