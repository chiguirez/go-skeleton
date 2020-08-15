#### Architectural notes
* Hexagonal architecture:
    * The root of the internal package is the domain layer. Each aggregate (value objects and entities) could be a different package here
    * Application layer is formed by the different packages refering to some action or service in each aggregate, like `creator` or `reader` for example.
        * Commands and query services shoud live here
    * Any other package refers to the infrastructure implementation of the domain, like `storage` or `messaging`
    * IO is outside the internal package
* Dependency injection with [google wire](https://github.com/google/wire)
* [Command bus](https://github.com/chiguirez/cromberbus). 
* [Snout](https://github.com/chiguirez/snout)
    
#### Prerequisites
* [GOLang](https://golang.org/) v1.14
* [Make](https://ftp.gnu.org/old-gnu/Manuals/make-3.79.1/html_chapter/make_2.html)

the following shows the commands

    > make help

    usage: make <command>
    
    commands:
        install              - get the modules
        test                 - run all tests
        run                  - execute
        
to install the modules

    > make install

#### Running

let

    > make run

#### Testing

to execute the tests

    > make test

#### Authors

* [Chiguirez](https://github.com/chiguirez):
    * [Edu Golding](https://github.com/K4L1Ma)
    * [Jose Ortiz](https://github.com/hosseio)
