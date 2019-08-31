#define X Y
#define foo(dir1) c->dir1 = d##dir1;
foo(X)
