#include <stdio.h>
#include "test.h"

void hello() {
    printf("Hello, world!\n");
}

#ifdef __TEST__

int main(int argc, char *argv[]) {
    hello();
    return 0;
}

#endif