A Go kernel for IPython
=======================

To install::

    pip install ipython[notebook]

    go get github.com/takluyver/igo

    mkdir -p ~/.ipython/kernels/igo
    cp -r $GOPATH/src/github.com/takluyver/igo/kernel/* ~/.ipython/kernels/igo

Edit ~/.ipython/kernels/igo/kernel.json and replace $GOPATH with your actual GOPATH

To run::

    ipython notebook

Go code is run using `go-eval <https://github.com/sbinet/go-eval/>`_.
