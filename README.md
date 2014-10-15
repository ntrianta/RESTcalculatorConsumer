RESTcalculatorConsumer
======================

1. Install Golang
2. Create a new work space:

        mkdir ~/go-workspace
3. Define your $GOPATH

        export GOPATH=~/go-workspace
4. Download the source code

        go get github.com/ntrianta/RESTcalculatorConsumer
5. Get the dependencies

        cd $GOPATH/src/github.com/ntrianta/RESTcalculatorConsumer
        go get
6. Build the binary

        go install
7. Execute the binary

        $GOPATH/bin/RESTcalculatorConsumer
