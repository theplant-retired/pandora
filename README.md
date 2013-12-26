pandora
=======

A Haskell zmq wrapper service for invoking Pandoc. Pandora ≈ (pando)c w(ra)pper



## How to setup pandora

#### 1. Install [Haskell Plaform] (http://www.haskell.org/platform/), make sure the runhaskell version >= 7.6.3 and the [cabal](http://www.haskell.org/cabal/) is installed 

```
➜  ~  runhaskell --version
runghc 7.6.3
➜  ~  cabal --version
cabal-install version 1.16.0.2
using version 1.16.0 of the Cabal library
```

#### 2. Clone the repo 

```
git clone git@github.com:theplant/pandora.git
```

#### 3. Compile the source

```
cd pandora
cabal install 
```
note: be patient, it takes quite a long time

#### 4. Start the daemon

```
pandora
```

#### 5. Check whether the service is running background

```
➜  ~  ps -ef | grep pandora
  501 17943     1   0 10:50AM ??         0:00.01 ./dist/build/pandora/pandora
  501 19664 19575   0 11:50AM ttys014    0:00.00 grep pandora
```



## TODO

- logs

- more tests

- more examples
