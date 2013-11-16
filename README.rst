A Go kernel for IPython.

To install::

    go get github.com/takluyver/igo

To run::

    ipython console --KernelManager.kernel_cmd="['igo', '{connection_file}']"

(You can substitute ``notebook`` or ``qtconsole`` for ``console``)

Go code is run using `go-eval <https://github.com/sbinet/go-eval/>`_.
