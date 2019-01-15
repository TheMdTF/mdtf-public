#include "sample_matching_algorithm.h"
#include <time.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

double fromBinary(const char *s) {
  return (double) strtol(s, NULL, 2);
}

int cpp_create_template(const char *image)
{
    int c = 0;
    for(int i = 0; i < strlen(image); i++)
    {
        c += (int)image[i];
        if (i > pow(2.0, 14.0)) break;
    }

    srand(time(NULL));
    int r = rand();

    return c * r;
}

double cpp_compare_template(const char *template2, const char *template1)
{
    double t1 = abs(fromBinary(template1));
    double t2 = abs(fromBinary(template2));

    if(t1 > t2) return (t2/t1);
    else return (t1/t2);
}
