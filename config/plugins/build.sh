for f in *.go; do
  echo $f
  go build -buildmode=plugin -o "${f%.go}.so"
done
