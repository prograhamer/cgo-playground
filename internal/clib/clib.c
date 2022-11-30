#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *Reverse(const char *data, size_t len) {
   char *target;
   target = (char *) malloc(len);

   for (int i = 0; i < len; i++) {
      target[i] = data[len - i - 1];
   }

   return target;
}

void ReverseInPlace(char *data, size_t len) {
   char t;
   size_t mid = len / 2;

   for (int i = 0; i < mid; i++) {
      t = data[i];
      data[i] = data[len-i-1];
      data[len-i-1] = t;
   }
}
