#include<stdio.h>
#include<string.h>

void func(void) {
    int *ptr;
    ptr = (int *)0x100;
    *ptr = 41;
}

void array_test() {
    int arr[2];
    0[arr] = 1;
    1[arr] = 2;
    int i;
    for(i=0;i<2;i++) {
        printf("%d\n", arr[i]);
    }
}

char *evil(char *a) {
    char s[128];
    strcpy(s, a);
    return s;
}

int main() {
    //func();
    array_test();
    char *str = "test";
    char *res = evil(str);
    return 0;
}
