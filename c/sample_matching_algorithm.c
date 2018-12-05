#include "sample_matching_algorithm.h"
#include <time.h>
#include <stdlib.h>

int cpp_extract_template()
{
   srand(time(NULL));
   int r = rand();
   return r;
}