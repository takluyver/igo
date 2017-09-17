A toy Go kernel for Jupyter
===========================

Warning
-------

This project was a 2013 proof-of-concept Go kernel for Jupyter in which Go code
was run using `go-eval <https://github.com/sbinet/go-eval/>`_.
The state of the Go eval tools was too limited for this kernel to be useful.
This project is no longer supported.

See also
--------

This project served as inspiration for a full-fledged Go kernel for Jupyter born in 2016
and called `gophernotes <https://github.com/gopherdata/gophernotes>`_.

See that project if you are looking for a usable, actively developed Go kernel for Jupyter.

Installation
------------

To install::

    pip install jupyter

    go get github.com/takluyver/igo

    mkdir -p ~/.jupyter/kernels/igo
    cp -r $GOPATH/src/github.com/takluyver/igo/kernel/* ~/.jupyter/kernels/igo

Edit ``~/.jupyter/kernels/igo/kernel.json`` and replace ``$GOPATH`` with your actual ``GOPATH``.

Running
-------

To run::

    jupyter notebook
