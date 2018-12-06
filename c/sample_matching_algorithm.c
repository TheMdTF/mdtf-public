#include "sample_matching_algorithm.h"
#include <time.h>
#include <stdlib.h>
#include <string.h>

int cpp_create_template(const char *image)
{
   int c = 0;
   for(int i = 0; i < strlen(image); i++) c += (int)image[i];

   srand(time(NULL));
   int r = rand();

   return c * r;
}