# Concurrency pattern 6 - Fan-In

`Fan-In` can be considered as the opposite of `Fan-Out`.

`Fan-In` can be considered as a function that might read from multiple inputs and continue until all are closed by multiplexing the input cannels onto a single channel which is closed once all the inputs are closed.

Essentially `Fan-In` is a process of combining multiple results from `goroutines` into one channel.

- output:

```bash
⇒  go run .

	Fan-In Concurrency Pattern
	Handling of multiple processes in the pipeline
	stage to handle cpu intensive or computationally
	intensive tasks.


m,95,w,109,111,o,w,114,174,o,o,78,123,i,b,i,g,32,u,28,s,171,q,n,27,127,o,153,t,31,s,145,k,64,z,i,156,x,78,94,t,y,58,40,n,21,a,x,171,200,a,33,v,c,94,r,80,o,50,j,s,41,174,f,n,24,g,87,b,43,k,103,y,8,l,i,146,g,194,b,165,r,m,x,84,79,v,f,113,38,s,o,134,91,f,z,59,65,g,37,b,12,y,136,b,x,p,y,83,p,9,u,y,122,102,p,56,
```