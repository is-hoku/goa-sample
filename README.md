# goa-sample
This is a sample REST API using [Goa v3](https://github.com/goadesign/goa/tree/v3).

## Usage
After the design is done, run this in the `/webapi` directory to create files.  
Note that `make goagen` removes all client files generated by `goa gen`.  
```
make goagen
```
Basically, [air](https://github.com/cosmtrek/air) builds the goa code automatically, but if you want to build it manually, run this.  
```
make build
```
And you can execute the binary.  
```
make exec
```
