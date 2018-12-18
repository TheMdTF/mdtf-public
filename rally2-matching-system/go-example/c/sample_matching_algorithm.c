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

double cpp_compare_template(const char *template2, const char *template1)
{
    int shortLength = -1;
    if(strlen(template1) > strlen(template2)) shortLength = strlen(template2);
    else shortLength = strlen(template1);

    double t1 = 0;
    double t2 = 0;
    for(int i = 0; i  < shortLength; i++)
    {
        t1 += (double)(int)template1[i];
        t2 += (double)(int)template2[i];
    }
    if(t1 > t2) return (t2 / t1);
    else return (t1 / t2);
}