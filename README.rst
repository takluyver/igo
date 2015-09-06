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

Support
-------

The project is not being supported currently.
The current state of the Go eval tools are too limited for it to be useful.
If you want to push it forwards, please get in touch.
