number=$(find . -maxdepth 5 -name '*.go' | wc -l)
echo "all algorithm: $number"
