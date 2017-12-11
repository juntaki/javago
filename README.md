# javago

Java wrapper code generator for Go.

## How to use


Install jnigo javago

~~~
go get github.com/juntaki/jnigo
go get github.com/juntaki/javago
~~~

Run command

~~~
$ javago -h
NAME:
   javago - Java binding generator for Go

USAGE:
   javago [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --classpath value  Specifies the path the javap command uses to look up classes. (default: "*:.")
   --classfile value  classfile
   --help, -h         show help
   --version, -v      print the version
~~~

e.g. generate java.lang.Math wrapper code.
Generated file is in sample directory in this repository.

~~~
$ javago --classfile java.lang.Math
$ ls java
java.lang.Math.go
~~~

## Run your app with javago wrapper

Don't forget to add libjvm.so's directory to LD_LIBRARY_PATH.

### for debian/ubuntu

~~~
export LD_LIBRARY_PATH=/usr/lib/jvm/default-java/jre/lib/amd64/server
export CLASSPATH="." # add your classpath
go build
./yourapp
~~~

### for mac

~~~
export JAVA_HOME=`/usr/libexec/java_home` 
export LD_LIBRARY_PATH=$JAVA_HOME/jre/lib/server/
export CLASSPATH="."
go build
./yourapp
~~~
