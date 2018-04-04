#include <stdio.h>
#include <stdlib.h>

int global;

int main()
{

  int local = 5;
  void *p = malloc(128);

  printf ("Address of main is %p\n", main);
  printf ("Address of global is %p\n", &global);
  printf ("Address of local is %p\n", &local);
  printf ("Address of p is %p\n", p);

}

// this small program demonstrates the arrangement of segments of virtual memory
// we would expect that main will have the lowest address space, followed by global (because it will be in the static segment)
// followed by p because it is allocated by malloc, and then finally local because it will be on the stack