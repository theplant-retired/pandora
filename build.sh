cabal build || { echo "build failed"; exit 1; }
killall -9 ghc
killall -9 pandora
./dist/build/pandora/pandora
echo "done"
