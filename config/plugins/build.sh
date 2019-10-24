for f in *.go; do
  go build -buildmode=plugin -o "${f%.go}.so"
done
