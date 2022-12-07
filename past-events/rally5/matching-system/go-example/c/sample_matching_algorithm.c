#include "sample_matching_algorithm.h"
#include <time.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <stdio.h>

double fromBase(const char *s)
{
   for(int i = 1; i < 32; i++){
     double val = (double)strtol(s, NULL, i);
     if(val != 0) return val;
   }

   return 1;
}

int cpp_create_template(const char *image)
{
    int c = 0;
    for(int i = 0; i < strlen(image); i++){
        c += (int)image[i];
        if (i > pow(2.0, 14.0)) break;
    }

    srand(time(NULL));
    int r = rand();

    return c * r;
}

double cpp_compare_template(const char *template2, const char *template1)
{
    double t1 = abs(fromBase(template1));
    double t2 = abs(fromBase(template2));
    double val = -1.0;

    if (t1 > t2) val = (t2 / t1);
    else val = (t1 / t2);
    while(val <= 1) val *= 10;

    return val / 10;
}
